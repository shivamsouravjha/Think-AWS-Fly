package server

import (
	"piepay/config"
	"piepay/routes"
	s3 "piepay/services/s3Bucket"
)

func Init() {
	s3.Init()
	config := config.Get()         //getting all the env configs
	r := routes.NewRouter()        //initialsing routes
	r.Run(":" + config.ServerPort) //running the server at port
}
