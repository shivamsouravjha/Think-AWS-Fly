package POST

import (
	"context"
	"fmt"
	"net/http"
	helpers "piepay/helpers/awsFile"
	"piepay/structs/response"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadImage(c *gin.Context) {
	defer sentry.Recover()
	span := sentry.StartSpan(context.TODO(), "[GIN] UploadDataHandler", sentry.TransactionName("Upload a file")) //setting transaction sentry
	defer span.Finish()

	file, err := c.FormFile("image")

	if err != nil {
		span.Status = sentry.SpanStatusFailedPrecondition
		sentry.CaptureException(err)
		c.JSON(http.StatusBadRequest, "failed")
	}

	uniqueId := uuid.New()

	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]
	image := fmt.Sprintf("%s.%s", filename, fileExt)
	err = c.SaveUploadedFile(file, fmt.Sprintf("./%s", image))
	resp := response.Response{}

	if err != nil {
		resp.Status = "Failed"
		resp.Message = err.Error()
		span.Status = sentry.SpanStatusFailedPrecondition
		sentry.CaptureException(err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	ctx := c.Request.Context()
	response, err := helpers.AcessAWSHelper(ctx, image, span.Context()) //the DAO level
	if err != nil {
		resp.Status = "Failed"
		resp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Status = "Success"
	resp.Message = response
	span.Status = sentry.SpanStatusOK

	c.JSON(http.StatusOK, resp)

}
