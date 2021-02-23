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
		resp := response.ResponseAPI("email checking failed", "error", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	path := fmt.Sprintf("assets/images/%s", file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false, "errors": err.Error()}
		resp := response.ResponseAPI("email checking failed", "error", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	cuurentUser := c.MustGet("current_user").(model.User)
	userID := int(cuurentUser.ID)
	fmt.Println("berapa nilai userID", userID)

	_, err = h.service.SaveAvatarService(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false, "errors": err.Error()}
		resp := response.ResponseAPI("email checking failed", "error", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	data := gin.H{"is_uploaded": true}
	resp := response.ResponseAPI("email checking failed", "success", http.StatusOK, data)
	c.JSON(http.StatusOK, resp)
}
