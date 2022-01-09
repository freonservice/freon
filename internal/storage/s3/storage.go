package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/freonservice/freon/internal/storage"
	"github.com/pkg/errors"
	"github.com/powerman/structlog"
)

type client struct {
	client        *s3.S3
	appleBucket   string
	androidBucket string
	webBucket     string
	logger        *structlog.Logger
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
