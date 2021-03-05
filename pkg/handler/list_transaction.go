package handler

import (
	"bwastartup/pkg/model"
	"bwastartup/pkg/request"
	"bwastartup/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) ListTransactionByCampaignIDHandler(c *gin.Context) {
	req := new(request.ListTransactionRequest)

	err := c.ShouldBindUri(req)
	if err != nil {
		response := response.ResponseAPI("Failed to get campaign transaction", "failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	req.User = c.MustGet("current_user").(model.User)
	
	transaction, err := h.service.GetTransactionByCampaignID(*req)
	if err != nil {
		response := response.ResponseAPI("Failed to get campaign's transaction", "failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	

	response := response.ResponseAPI("Success to list campaign's transaction", "success", http.StatusOK, response.FormaterListCampaignTransactionResponse(transaction))
	c.JSON(http.StatusOK, response)
}
