package handler

import (
	"bwastartup/pkg/request"
	"bwastartup/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) CheckEmailAvailabilityHandler(c *gin.Context) {
	req := new(request.CheckEmailAvailable)
	err := c.ShouldBind(req)
	if err != nil {
		errorMessage := gin.H{"error": response.ErrorValidationResponse(err)}
		resp := response.ResponseAPI("email checking failed", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	isEmailAvailable, err := h.service.CheckEmailAvailabilityService(*req)
	if err != nil {
		errorMessage := gin.H{"error": "server error"}
		resp := response.ResponseAPI("email checking failed", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	metaMessage := "Email has been registered"
	if isEmailAvailable {
		metaMessage = "Email is available"
	}
	resp := response.ResponseAPI(metaMessage, "success", http.StatusAccepted, isEmailAvailable)
	c.JSON(http.StatusAccepted, resp)
}