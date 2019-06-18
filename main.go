package main

import "sample/s3mightiness"

func main() {
	s3mightiness.S3Update("test", "testttt")
	s3mightiness.S3Cpoy("test2", "test/testobject", "testboject")
	s3mightiness.S3Download("test", "testttt", "tttttt4")
	s3mightiness.S3Delete("test2", "testboject")
}
