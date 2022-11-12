package GET

import (
	"fmt"
	"net/http"
	helpers "piepay/helpers/awsFile"
	"piepay/structs/response"

	"github.com/gin-gonic/gin"
)

func GetImages(c *gin.Context) {
	resp := response.ResponseFile{}
	ctx := c.Request.Context()
	response, err := helpers.FetchAWSBuckets(ctx) //the DAO level
	if err != nil {
		fmt.Println("Unable to list buckets, ", err)
		c.JSON(http.StatusInternalServerError, err)
	}
	resp.Status = "Success"
	resp.BucketNames = response
	c.JSON(http.StatusOK, resp)
}
