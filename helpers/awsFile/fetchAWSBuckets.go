package helpers

import (
	"context"
	s3 "piepay/services/s3Bucket"
)

func FetchAWSBuckets(ctx context.Context) ([]string, error) {
	response := make([]string, 0)
	result, err := s3.AWSService().ListBuckets(nil)
	for _, a := range result.Buckets {
		response = append(response, *a.Name)
	}
	if err != nil {
		return response, err
	}
	return response, nil

}
