package cleaner

import (
	"context"
	"fmt"
	"log"
	"time"

	"piepay/config"
	s3bucket "piepay/services/s3Bucket"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var timer *time.Ticker //initalise ticker to tick at 1000 s

func Cleaner() {
	if timer == nil {
		timer = time.NewTicker(24 * time.Hour) //initialise ticker
		fmt.Println("Started the ticker wait for 24 hours to delete files")
		go func() {
			for range timer.C {
				cleaner() //after every 1000s upload data to db
			}
		}()
	}
}
func cleaner() {
	sess := s3bucket.AWSService()
	resp, err := sess.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(config.Get().Bucket)})
	if err != nil {
		fmt.Println("Unable to list items in bucket %", config.Get().Bucket, err)

	}
	currentTime := time.Now().Add(-24 * time.Hour)

	for _, item := range resp.Contents {
		fmt.Println("Name:         ", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:         ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
		if item.LastModified.Before(currentTime) {
			iter := s3manager.NewDeleteListIterator(sess, &s3.ListObjectsInput{
				Bucket: aws.String(config.Get().Bucket),
				Prefix: aws.String(*item.Key),
			})
			if err := s3manager.NewBatchDeleteWithClient(sess).Delete(context.TODO(), iter); err != nil {
				log.Fatalf("failed to delete files under given directory: %v", err)
			}
		}

	}
}
