package s3mightiness

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func S3Cpoy(bucket string, item string, other string) {
	newSession := session.New(S3Access())
	s3Client := s3.New(newSession)

	/*Bucket:コピー先bucket
	  CopySource:コピー元bucket/コピーもとファイルkey（スラッシュで区切り）
	  Key:コピー先ディレクトリ、ファイル名
	*/
	_, err := s3Client.CopyObject(&s3.CopyObjectInput{Bucket: aws.String(bucket), CopySource: aws.String(item), Key: aws.String(other)})

	if err != nil {
		fmt.Printf("Failed to copy", err.Error())
		return
	}
}
