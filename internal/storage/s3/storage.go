package s3

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/freonservice/freon/internal/storage"
	api "github.com/freonservice/freon/pkg/freonApi"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/private/protocol/rest"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
	"github.com/powerman/structlog"
)

const (
	uploadTimeout = 15 * time.Second
)

type client struct {
	client        *s3.S3
	appleBucket   string
	androidBucket string
	webBucket     string
	logger        *structlog.Logger
}

func (c *client) Create(parameter storage.FileParameter) (*storage.File, error) {
	var (
		err                error
		fileName           string
		webFullPath        string
		localizationFolder string
		fileID             string

		bucket       = c.getBucketByPlatform(parameter.Platform)
		platformType = storage.GetPlatformByString(parameter.Platform)
	)
	switch api.PlatformType(platformType) { //nolint:exhaustive
	case api.PlatformType_PLATFORM_TYPE_IOS:
		localizationFolder = "/" + parameter.LocalizationLocale + ".lproj"
		fileName = storage.DefaultAppleFile
	case api.PlatformType_PLATFORM_TYPE_ANDROID:
		localizationFolder = "/values-" + parameter.LocalizationLocale
		fileName = storage.DefaultAndroidFile
	case api.PlatformType_PLATFORM_TYPE_WEB:
		fileName = fmt.Sprintf("%s.json", parameter.LocalizationLocale)
	}

	if api.PlatformType(platformType) != api.PlatformType_PLATFORM_TYPE_WEB {
		err = c.checkOrCreateFolder(bucket, localizationFolder)
		if err != nil {
			return nil, err
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), uploadTimeout)
	defer cancel()
	if api.PlatformType(platformType) == api.PlatformType_PLATFORM_TYPE_IOS {
		var text = []string{parameter.TranslatedText.TextFirst, parameter.TranslatedText.TextSecond}
		var webPath string
		var webPaths []string
		var fileIDs []string
		for i := range text {
			webPath, fileID, err = c.upload(ctx, []byte(text[i]), bucket, localizationFolder, fileName+storage.IosFormat[i])
			if err != nil {
				return nil, err
			}
			webPaths = append(webPaths, webPath)
			fileIDs = append(fileIDs, fileID)
		}
		webFullPath = strings.Join(webPaths, ",")
		fileID = strings.Join(fileIDs, ",")
	} else {
		webFullPath, fileID, err = c.upload(ctx, []byte(parameter.TranslatedText.TextFirst), bucket, localizationFolder, fileName)
		if err != nil {
			return nil, err
		}
	}

	return &storage.File{
		Name:     fileName,
		WebPath:  webFullPath,
		S3FileID: fileID,
		S3Bucket: bucket,
	}, nil
}

func NewStorage(config *StorageConfiguration, logger *structlog.Logger) (storage.Storage, error) {
	err := config.checkValidity()
	if err != nil {
		return nil, err
	}

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(config.AccessKeyID, config.SecretAccessKey, ""),
		Endpoint:         aws.String(config.URL),
		Region:           aws.String(config.Region),
		DisableSSL:       aws.Bool(config.DisableSSL),
		S3ForcePathStyle: aws.Bool(config.ForcePathStyle),
	}
	newSession, err := session.NewSession(s3Config)
	if err != nil {
		return nil, err
	}

	c := &client{
		client:        s3.New(newSession),
		appleBucket:   config.AppleBucket,
		androidBucket: config.AndroidBucket,
		webBucket:     config.WebBucket,
		logger:        logger,
	}

	err = c.initBucketsIfNotExist()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *client) initBucketsIfNotExist() error {
	bucketList := []string{c.appleBucket, c.androidBucket, c.webBucket}
	for i := range bucketList {
		_, err := c.client.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(bucketList[i]),
			ACL:    aws.String("public-read"),
		})
		if err != nil && !isAlreadyOwnedByMeError(err) {
			return errors.Wrapf(err, "Unable to create bucket %s", bucketList[i])
		}
		// Wait until bucket is created before finishing
		c.logger.Printf("Waiting for bucket %q to be created...\n", bucketList[i])
		err = c.client.WaitUntilBucketExists(&s3.HeadBucketInput{
			Bucket: aws.String(bucketList[i]),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func isAlreadyOwnedByMeError(err error) bool {
	if awsErr, ok := err.(awserr.Error); ok {
		return awsErr.Code() == "BucketAlreadyOwnedByYou"
	}
	return false
}

func (c *client) delete(fileID, bucket string) error { // nolint:unused
	_, err := c.client.DeleteObject(&s3.DeleteObjectInput{
		Key:    aws.String(fileID),
		Bucket: aws.String(bucket),
	})
	return err
}

func (c *client) checkOrCreateFolder(bucketName, folder string) error {
	_, err := c.client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(folder + "/"),
	})
	return err
}

func (c *client) getBucketByPlatform(platform string) string {
	switch platform {
	case "ios":
		return c.appleBucket
	case "android":
		return c.androidBucket
	default:
		return c.webBucket
	}
}

func (c *client) upload( //nolint:gocritic
	ctx context.Context, body []byte,
	bucket, localizationFolder, fileName string,
) (string, string, error) {
	fileID := localizationFolder + "/" + fileName
	req, _ := c.client.PutObjectRequest(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(fileID),
		Body:        strings.NewReader(string(body)),
		ContentType: aws.String("text/plain"),
	})
	req.SetContext(ctx)

	err := req.Send()
	if err != nil {
		return "", "", err
	}

	req, _ = c.client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileID),
	})
	rest.Build(req)
	return req.HTTPRequest.URL.Path, fileID, nil
}
