package handler

import (
	"bwastartup/pkg/model"
	"bwastartup/pkg/request"
	"bwastartup/pkg/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateCampaignImageHandler(c *gin.Context) {
	request := new(request.CreateCampaignImageRequest)

	err := c.ShouldBind(request)
	if err != nil {
		errorMessage := gin.H{"error": response.ErrorValidationResponse(err)}
		resp := response.ResponseAPI("failed to create upload campaign image", "failed", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		resp := response.ResponseAPI("failed to upload campaign image", "error", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	currentUser := c.MustGet("current_user").(model.User)
	path := fmt.Sprintf("assets/images/campaign_image/%s_%s", currentUser.Fullname, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false, "errors": err.Error()}
		resp := response.ResponseAPI("save campaign image failed", "error", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	request.User = currentUser
	_, err = h.service.CreateCampaignImage(request, path)
	if err != nil {
		data := gin.H{"is_uploaded": false, "errors": err.Error()}
		resp := response.ResponseAPI("campaign image upload failed", "error", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	data := gin.H{"is_uploaded": true}
	resp := response.ResponseAPI("campaign image upload success", "success", http.StatusOK, data)
	c.JSON(http.StatusOK, resp)
}
