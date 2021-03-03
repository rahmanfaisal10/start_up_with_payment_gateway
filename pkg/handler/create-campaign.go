package handler

import (
	"bwastartup/pkg/model"
	"bwastartup/pkg/request"
	"bwastartup/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateCampaignHandler(c *gin.Context) {
	req := new(request.CreateCampaignRequest)

	err := c.ShouldBindJSON(req)
	if err != nil {
		errorMessage := gin.H{"error": response.ErrorValidationResponse(err)}
		resp := response.ResponseAPI("failed to create campaign", "failed", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	currentUser := c.MustGet("current_user").(model.User)
	req.User = currentUser

	campaign, err := h.service.Createcampaign(*req)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		resp := response.ResponseAPI("failed to create campaign", "failed", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := response.ResponseAPI("success to create campaign", "success", http.StatusCreated, response.FormaterCampaignResponse(campaign))
	c.JSON(http.StatusCreated, resp)
}
