package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/go-chi/chi/v5"
	"github.com/gvcastellain/go-driver/internal/auth"
	"github.com/gvcastellain/go-driver/internal/bucket"
	"github.com/gvcastellain/go-driver/internal/files"
	"github.com/gvcastellain/go-driver/internal/folders"
	"github.com/gvcastellain/go-driver/internal/queue"
	"github.com/gvcastellain/go-driver/internal/users"
	"github.com/gvcastellain/go-driver/pkg/database"
)

func main() {
	db, b, qc := getSessions()

	r := chi.NewRouter()
	r.Post("/auth", auth.HandleAuth(func(login, pass string) (auth.Authenticated, error) {
		return users.Authenticate(login, pass)
	}))

	files.SetRoutes(r, db, b, qc)
	folders.SetRoutes(r, db)
	users.SetRoutes(r, db)

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")), r)
}

func getSessions() (*sql.DB, *bucket.Bucket, *queue.Queue) {
	db, err := database.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	qcfg := queue.RabbitMQConfig{
		URL:       os.Getenv("RABBIT_URL"),
		TopicName: os.Getenv("RABBIT_TOPIC_NAME"),
		Timeout:   time.Second * 30,
	}

	qc, err := queue.New(queue.RabbitMQ, qcfg)

	if err != nil {
		panic(err) //todo
	}

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

	return db, b, qc
}
