package s3

import (
	"piepay/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var awsSession s3.S3
var awsSes *session.Session

func Init() {
	sess := session.Must(session.NewSession(
		&aws.Config{
			Region:      aws.String("ap-south-1"),
			Credentials: credentials.NewStaticCredentials(config.Get().AWSAccessKey, config.Get().AWSSecretKey, ""),
		},
	))
	svc := s3.New(sess, &aws.Config{
		DisableRestProtocolURICleaning: aws.Bool(true),
	})
	awsSes = sess
	awsSession = *svc
}

func AWSService() *s3.S3 {
	return &awsSession
}
func AwsSession() *session.Session {
	return awsSes
}

// func Init() *elastic.Client {
// 	esOnce.Do(func() {
// 		esURL := config.Get().EsURL
// 		Key := config.Get().Key
// 		var err error
// 		client, err = elastic.NewClient(elastic.SetURL(esURL), elastic.SetHttpClient(&http.Client{
// 			Transport: apmelasticsearch.WrapRoundTripper(http.DefaultTransport),
// 		}), elastic.SetSniff(false), elastic.SetBasicAuth("elastic", Key), elastic.SetScheme("https"))
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		_, _, err = client.Ping(config.Get().EsURL).Do(context.Background()) //ping elasticsearch to see if connection is stable
// 		if err != nil {
// 			panic(err)
// 		}

// 	})
// 	return client
// }

// func Client() *elastic.Client {
// 	return client
// }
