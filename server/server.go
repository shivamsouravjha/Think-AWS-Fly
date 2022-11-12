package server

import (
	"log"
	"piepay/config"
	"piepay/routes"
	"piepay/services/es"
	"piepay/utils"
	"time"

	"github.com/getsentry/sentry-go"
)

func Init() {

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
	es.Init()                      //initalising connection to elastic search
	go utils.Uploader()            // starting the backgroup goroutine to upload the videos
	r := routes.NewRouter()        //initialsing routes
	r.Run(":" + config.ServerPort) //running the server at port
}
