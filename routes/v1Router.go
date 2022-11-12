package routes

import (
	POST "piepay/controllers/POST"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

func v1Routes(route *gin.RouterGroup) {

	router := gin.New()
	router.Use(apmgin.Middleware(router))

	v1Routes := route.Group("/v1")
	{
		v1Routes.POST("/upload/image", POST.UploadImage)
	}
}
