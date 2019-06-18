package s3mightiness

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func S3Delete(bucket string, key string) {

	newSession := session.New(S3Access())
	s3Client := s3.New(newSession)

	_, err := s3Client.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(bucket), Key: aws.String(key)})

	err = s3Client.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}) // errorが帰ってこなければ成功

	if err != nil {
		fmt.Printf("Failed to delete", err.Error())
		return
	}
}
