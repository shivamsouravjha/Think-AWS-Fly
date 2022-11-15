package server

import (
	"log"
	"piepay/config"
	"piepay/routes"
	s3 "piepay/services/s3Bucket"
	"time"

	"github.com/getsentry/sentry-go"
)

func Init() {
	s3.Init()
	config := config.Get() //getting all the env configs

	err := sentry.Init(sentry.ClientOptions{ //initialsing sentry for error tracing and monitoring performance
		Dsn:              config.SentryDSN,
		Debug:            true,
		Environment:      config.AppEnv,
		TracesSampleRate: float64(config.SentrySamplingRate),
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)
	r := routes.NewRouter()        //initialsing routes
	r.Run(":" + config.ServerPort) //running the server at port
}
