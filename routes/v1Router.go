package routes

import (
	"piepay/controllers/GET"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

func v1Routes(route *gin.RouterGroup) {

	router := gin.New()
	router.Use(apmgin.Middleware(router))

	v1Routes := route.Group("/v1")
	{
		v1Routes.GET("/get", GET.GetVideo)       //the route to get the paginated video
		v1Routes.GET("/search", GET.SearchVideo) //the route to search the video on basis of title and description
	}
}
