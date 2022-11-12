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

func SearchVideo(c *gin.Context) {
	defer sentry.Recover()
	span := sentry.StartSpan(context.TODO(), "[GIN] SearchHandler", sentry.TransactionName("Search any video handler"))
	defer span.Finish()

	formRequest := requests.SearchVideo{}

	if err := c.ShouldBind(&formRequest); err != nil {
		span.Status = sentry.SpanStatusFailedPrecondition
		sentry.CaptureException(err)
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	ctx := c.Request.Context()
	resp := response.VideoResponse{}
	var response response.GetVideo

	var err error
	if len(formRequest.Description) == 0 && len(formRequest.Title) == 0 { //if both title and description is empty its searching anything
		getRequest := requests.GetVideo{
			Page: formRequest.Page,
			Size: formRequest.Size,
		}
		response, err = helpers.GetLatestVideo(ctx, &getRequest, span.Context())

	} else {
		response, err = helpers.GetSearchVideo(ctx, &formRequest, span.Context())
	}

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
