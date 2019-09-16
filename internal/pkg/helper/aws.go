package helper

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// GetFileFromS3 fetchs a file from S3 storage.
// The AWS credential is included in the function.
func GetFileFromS3(imagePath string) (interface{}, error) {
	// 1. Define your bucket and item names
	bucket := "static.kiren.tanawitp.me"
	item := imagePath

	// 2. Create an AWS session
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
	})

	// 3. Create a new AWS S3 downloader
	downloader := s3manager.NewDownloader(sess)

	// 4. Download the item from the bucket. If an error occurs, log it and exit.
	// Otherwise, notify the user that the download succeeded.
	file, err := os.Create(item)
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		},
	)
	if err != nil {
		log.Fatalf("Unable to download item %q, %v", item, err)
	}
	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
	return file, nil
}
