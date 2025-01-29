package bucket

import (
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type AwsConfig struct {
	Config         *aws.Config
	BucketDownload string
	BucketUpload   string
}

func newAwsSession(cfg AwsConfig) *awsSession {
	c := session.New(cfg.Config)

	return &awsSession{
		sess:           c,
		BucketDownload: cfg.BucketDownload,
		BucketUpload:   cfg.BucketUpload,
	}
}

type awsSession struct {
	sess           *session.Session
	BucketDownload string

	BucketUpload string
}

func (as *awsSession) Download(src, dst string) (file *os.File, err error) {
	return
}

func (as *awsSession) Upload(file io.Reader, key string) error {
	return nil
}

func (as *awsSession) Delete(key string) error {
	return nil
}
