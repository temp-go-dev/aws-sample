package s3mightiness

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func S3Download(newSession *session.Session, bucket string, key string, fileRoot string) error {

	//ローカルにファイル生成
	file, err := os.Create(fileRoot)
	if err != nil {
		return err
	}
	defer file.Close()

	//ファイルをダウンロード
	downloader := s3manager.NewDownloader(newSession)
	_, err = downloader.Download(file,
		&s3.GetObjectInput{
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
