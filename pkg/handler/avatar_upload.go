package handler

import (
	"bwastartup/pkg/model"
	"bwastartup/pkg/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) UploadAvatarHandler(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		resp := response.ResponseAPI("error get file", "error", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	currentUser := c.MustGet("current_user").(model.User)
	path := fmt.Sprintf("assets/images/avatar/%s_%s", currentUser.Fullname, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false, "errors": err.Error()}
		resp := response.ResponseAPI("save avatar failed", "error", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	_, err = h.service.SaveAvatarService(currentUser.UUID.String(), path)
	if err != nil {
		data := gin.H{"is_uploaded": false, "errors": err.Error()}
		resp := response.ResponseAPI("avatar upload failed", "error", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	data := gin.H{"is_uploaded": true}
	resp := response.ResponseAPI("avatar upload success", "success", http.StatusOK, data)
	c.JSON(http.StatusOK, resp)
}
