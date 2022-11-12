package server

import (
	"piepay/config"
	"piepay/routes"
)

func Init() {

	config := config.Get()         //getting all the env configs
	r := routes.NewRouter()        //initialsing routes
	r.Run(":" + config.ServerPort) //running the server at port
}
