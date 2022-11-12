package POST

import (
	"fmt"
	"net/http"
	helpers "piepay/helpers/es"
	"piepay/structs/response"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("image")

	if err != nil {
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
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	ctx := c.Request.Context()
	response, err := helpers.AcessAWSHelper(ctx, image) //the DAO level
	if err != nil {
		resp.Status = "Failed"
		resp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Status = "Success"
	resp.Message = response
	c.JSON(http.StatusOK, resp)

}
