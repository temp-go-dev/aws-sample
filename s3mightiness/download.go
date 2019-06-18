package s3mightiness

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func S3Download(bucket string, key string, fileRoot string) {
	newSession := session.New(S3Access())

	file, err := os.Create(fileRoot)
	if err != nil {
		fmt.Println("Failed to create file", err)
		return
	}
	defer file.Close()

	downloader := s3manager.NewDownloader(newSession)
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})
	if err != nil {
		fmt.Println("Failed to download file", err)
		return
	}
	fmt.Println("Downloaded file", file.Name(), numBytes, "bytes")
}
