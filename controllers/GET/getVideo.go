package GET

import (
	"context"
	"net/http"
	helpers "piepay/helpers/es"
	"piepay/structs/requests"
	"piepay/structs/response"
	"piepay/utils"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func GetVideo(c *gin.Context) {
	defer sentry.Recover()
	span := sentry.StartSpan(context.TODO(), "[GIN] GetVideoHandler", sentry.TransactionName("Get latest video")) //setting transaction sentry
	defer span.Finish()

	formRequest := requests.GetVideo{}

	if err := c.ShouldBind(&formRequest); err != nil {
		span.Status = sentry.SpanStatusFailedPrecondition
		sentry.CaptureException(err)
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	ctx := c.Request.Context()
	resp := response.VideoResponse{}

	response, err := helpers.GetLatestVideo(ctx, &formRequest, span.Context()) //the DAO level
	if err != nil {
		resp.Status = "Failed"
		resp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Status = "Success"
	resp.Message = "Video fetched successfully"
	resp.Data = response
	span.Status = sentry.SpanStatusOK

	c.JSON(http.StatusOK, resp)

}
