package helpers

import (
	"context"
	"fmt"
	"os"
	"piepay/config"
	s3 "piepay/services/s3Bucket"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func AcessAWSHelper(ctx context.Context, filename string) (string, error) {
	// The session the S3 Uploader will use
	sess := s3.AwsSession()

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("failed to open file ", err)
		return "", err
	}

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.Get().Bucket),
		Key:    aws.String(config.Get().AWSSecretKey),
		Body:   f,
	})
	if err != nil {
		fmt.Println("failed to upload file, ", err)
		return "", err

	}
	go os.Remove(filename)
	return result.Location, nil

}
