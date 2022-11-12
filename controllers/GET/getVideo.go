package GET

import (
	"fmt"
	"net/http"
	helpers "piepay/helpers/es"
	"piepay/structs/requests"
	"piepay/structs/response"
	"piepay/utils"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func AcessAWS(c *gin.Context) {

	formRequest := requests.GetVideo{}

	if err := c.ShouldBind(&formRequest); err != nil {
		sentry.CaptureException(err)
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	ctx := c.Request.Context()
	resp := response.VideoResponse{}

	response, err := helpers.AcessAWSHelper(ctx) //the DAO level
	if err != nil {
		resp.Status = "Failed"
		resp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Status = "Success"
	resp.Message = "Video fetched successfully"
	fmt.Println(response)
	c.JSON(http.StatusOK, resp)

}
