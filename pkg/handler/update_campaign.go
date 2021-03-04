package handler

import (
	"bwastartup/pkg/model"
	"bwastartup/pkg/request"
	"bwastartup/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) UpdateCampaignHandler(c *gin.Context) {
	reqID := new(request.DetailCampaignRequest)
	reqData := new(request.CreateCampaignRequest)

	err := c.ShouldBindUri(reqID)
	if err != nil {
		response := response.ResponseAPI("Failed to get update of campaign", "failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = c.ShouldBindJSON(reqData)
	if err != nil {
		errorMessage := gin.H{"error": response.ErrorValidationResponse(err)}
		resp := response.ResponseAPI("failed to update campaign", "failed", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	currentUser := c.MustGet("current_user").(model.User)
	reqData.User = currentUser

	updateCampaign, err := h.service.UpdateCampaign(*reqID, *reqData)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		resp := response.ResponseAPI("failed to update campaign", "failed", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := response.ResponseAPI("success to update campaign", "success", http.StatusOK, response.FormaterCampaignResponse(updateCampaign))
	c.JSON(http.StatusOK, resp)
}
