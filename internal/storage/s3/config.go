package s3

import "errors"

type StorageConfiguration struct {
	SecretAccessKey string
	AccessKeyID     string
	Region          string
	AppleBucket     string
	AndroidBucket   string
	WebBucket       string
	URL             string
	DisableSSL      bool
	ForcePathStyle  bool
}

func (s *StorageConfiguration) checkValidity() error {
	if len(s.SecretAccessKey) == 0 { //nolint:gocritic
		return errors.New("s3 config error. Empty SecretAccessKey")
	}
	if len(s.AccessKeyID) == 0 { //nolint:gocritic
		return errors.New("s3 config error. Empty AccessKeyID")
	}
	if len(s.Region) == 0 { //nolint:gocritic
		return errors.New("s3 config error. Empty Region")
	}
	if len(s.AppleBucket) == 0 { //nolint:gocritic
		return errors.New("s3 config error. Empty AppleBucket")
	}
	if len(s.AndroidBucket) == 0 { //nolint:gocritic
		return errors.New("s3 config error. Empty AndroidBucket")
	}
	if len(s.WebBucket) == 0 { //nolint:gocritic
		return errors.New("s3 config error. Empty WebBucket")
	}
	if len(s.URL) == 0 { //nolint:gocritic
		return errors.New("s3 config error. Empty URL")
	}
	return nil
}
