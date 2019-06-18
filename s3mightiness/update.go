package s3mightiness

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func S3Update(bucket string, key string) {

	newSession := session.New(S3Access())
	s3Client := s3.New(newSession)

	_, err := s3Client.PutObject(&s3.PutObjectInput{
		Body:   strings.NewReader("Hello from MinIO!!"),
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		fmt.Printf("Failed to upload data to %s/%s, %s\n", bucket, key, err.Error())
		return
	}
	fmt.Printf("Successfully created bucket %s and uploaded data with key %s\n", bucket, key)
}
