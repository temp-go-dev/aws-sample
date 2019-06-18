package s3mightiness

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func S3Update(newSession *session.Session, bucket string, key string) error {
	s3Client := s3.New(newSession)
	_, err := s3Client.PutObject(&s3.PutObjectInput{
		Body:   strings.NewReader("Hello from MinIO!!"),
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		//エラーから取れる情報(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		return err
	} else {
		return nil
	}
}
