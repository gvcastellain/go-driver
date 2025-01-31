package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/gvcastellain/go-driver/internal/bucket"
	"github.com/gvcastellain/go-driver/internal/queue"
)

func main() {
	qcfg := queue.RabbitMQConfig{
		URL:       os.Getenv("RABBIT_URL"),
		TopicName: os.Getenv("RABBIT_TOPIC_NAME"),
		Timeout:   time.Second * 30,
	}

	qc, err := queue.New(queue.RabbitMQ, qcfg)

	if err != nil {
		panic(err) //todo
	}

	c := make(chan queue.QueueDto)
	qc.Consume(c)

	bcfg := bucket.AwsConfig{
		Config: &aws.Config{
			Region:      aws.String(os.Getenv("AWS_REGION")),
			Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_KEY"), os.Getenv("AWS_SECRET"), ""),
		},
		BucketDownload: "go-drive-raw",
		BucketUpload:   "go-drive-gzip",
	}

	b, err := bucket.New(bucket.AwsProvider, bcfg)

	if err != nil {
		panic(err)
	}

	for msg := range c {
		src := fmt.Sprintf("%d_%s", msg.Path, msg.Filename)
		dst := fmt.Sprintf("%d_%s", msg.ID, msg.Filename)

		file, err := b.Download(src, dst)
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}

		body, err := io.ReadAll(file)
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}

		var buf bytes.Buffer
		zw := gzip.NewWriter(&buf)
		_, err = zw.Write(body)
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}

		if err := zw.Close(); err != nil {
			log.Printf("error: %v", err)
			continue
		}

		zr, err := gzip.NewReader(&buf)
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}

		if b.Upload(zr, src); err != nil {
			log.Printf("error: %v", err)
			continue
		}

	}
}
