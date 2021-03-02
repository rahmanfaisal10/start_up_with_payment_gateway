package handler

import (
	"bwastartup/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) ListCampaignHandler(c *gin.Context) {
	userID := c.Query("user_id")

	campaign, err := h.service.ListCampaign(userID)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		resp := response.ResponseAPI("Error to get campaigns", "failed", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp := response.ResponseAPI("succes to list campaigns", "success", http.StatusOK, campaign)
	c.JSON(http.StatusOK, resp)
}
