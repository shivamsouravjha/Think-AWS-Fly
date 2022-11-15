package helpers

import (
	"context"
	"fmt"
	"os"
	"piepay/config"
	s3 "piepay/services/s3Bucket"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/getsentry/sentry-go"
)

func AcessAWSHelper(ctx context.Context, filename string, sentryCtx context.Context) (string, error) {
	// The session the S3 Uploader will use
	defer sentry.Recover()
	span := sentry.StartSpan(sentryCtx, "[DAO] UploadDataHandler") //sentry to log db calls
	defer span.Finish()
	sess := s3.AwsSession()

	dbSpan1 := sentry.StartSpan(span.Context(), "[DB] Upload data on S3 Bucket")

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	dbSpan1.Finish()
	f, err := os.Open(filename)
	if err != nil {
		sentry.CaptureException(err)
		fmt.Println("failed to open file ", err)
		return "", err
	}

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.Get().Bucket),
		Key:    &filename,
		Body:   f,
	})
	dbSpan1.Finish()

	if err != nil {
		span.Status = sentry.SpanStatusFailedPrecondition
		sentry.CaptureException(err)
		fmt.Println("failed to upload file, ", err)
		return "", err

	}
	go os.Remove(filename)
	return result.Location, nil

}
