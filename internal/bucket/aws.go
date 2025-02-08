package bucket

import (
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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
	file, err = os.Create(dst) //TODO - maybe delete file here by defer instead of in the main func
	if err != nil {
		return nil, err
	}

	defer file.Close()

	downloader := s3manager.NewDownloader(as.sess)

	_, err = downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(as.BucketDownload),
		Key:    aws.String(src),
	})

	if err != nil {
		return nil, err
	}

	return
}

func (as *awsSession) Upload(file io.Reader, key string) error {
	uploader := s3manager.NewUploader(as.sess)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(as.BucketDownload),
		Key:    aws.String(key),
		Body:   file,
	})

	return err
}

func (as *awsSession) Delete(key string) error {
	svc := s3.New(as.sess)

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(as.BucketDownload),
		Key:    aws.String(key),
	})

	if err != nil {
		return err
	}

	return svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(as.BucketDownload),
		Key:    aws.String(key),
	})
}
