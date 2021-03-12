package request

import "bwastartup/pkg/model"

type ListTransactionRequest struct {
	CampaignID string `uri:"campaign_id" binding:"required"`
	User       model.User
}

type ListUserTransactionRequest struct {
	UserID string `uri:"user_id" binding:"required"`
}
