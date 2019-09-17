package helper

import (
	"bytes"
	"io"
	"kiren-backend-go/internal/app"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// GetFileFromS3 fetchs a file from S3 storage.
// The function fetches the data from `mocked/products.json` on local environment.
// Otherwise, it fetches the data from S3 storage.
// The AWS credential is imported from environment variables.
func GetFileFromS3(imagePath string) ([]byte, error) {
	// Define bucket and item names
	bucket := "static.kiren.tanawitp.me"
	item := imagePath
	// fileByte := make([]byte, 0)
	var fileByte []byte

	// Create an AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(
			app.CFG.AWS.AccessKeyID,
			app.CFG.AWS.SecretAccessKey,
			"",
		),
	})
	if err != nil {
		return fileByte, err
	}

	// Create a new AWS S3 downloader
	downloader := s3manager.NewDownloader(sess)

	// Create an emtry file.
	file, err := os.Create(item)
	if err != nil {
		return fileByte, err
	}

	// Download the item from the bucket.
	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		},
	)
	if err != nil {
		return fileByte, err
	}

	// Convert File to []byte
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, file)
	fileByte = buf.Bytes()

	return fileByte, nil
}
