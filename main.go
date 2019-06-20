package main

import (
	"github.com/temp-go-dev/aws-sample"

	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	s3Config := s3mightiness.Prop{
		AccessKey:       "hogehoge",
		SecretKey:       "hogehoge",
		ServiceEndpoint: "http://127.0.0.1:9001",
		Region:          "ut1",
	}
	s3Client := session.New(s3Config.S3Access())

	s3mightiness.S3Update(s3Client, "test", "testteeetttt")
	s3mightiness.S3Cpoy(s3Client, "test2", "test/testobject", "testboject")
	s3mightiness.S3Download(s3Client, "test", "testttt", "tttttt4")
	s3mightiness.S3Delete(s3Client, "test2", "testboject")
}
