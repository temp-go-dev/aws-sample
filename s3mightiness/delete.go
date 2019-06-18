package s3mightiness

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func S3Delete(newSession *session.Session, bucket string, key string) error {

	s3Client := s3.New(newSession)

	_, err := s3Client.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(bucket), Key: aws.String(key)})

	err = s3Client.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}) // errorが帰ってこなければ成功

	if err != nil {
		//エラーから取れる情報(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		return err
	} else {
		return nil
	}
}
