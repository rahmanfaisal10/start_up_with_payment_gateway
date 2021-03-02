package handler

import (
	"bwastartup/pkg/request"
	"bwastartup/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) DetailCampaignHandler(c *gin.Context) {
	req := new(request.DetailCampaignRequest)

	err := c.ShouldBindUri(req)
	if err != nil {
		response := response.ResponseAPI("Failed to get detail of campaign", "failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaign, err := h.service.DetailCampaign(*req)
	if err != nil {
		response := response.ResponseAPI("Failed to get detail of campaign", "failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := response.ResponseAPI("Success to get detail of campaign", "success", http.StatusOK, response.FormatterDetailCampaignResponse(campaign))
	c.JSON(http.StatusOK, response)
}
