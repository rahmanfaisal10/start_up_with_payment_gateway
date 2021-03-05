package request

import "bwastartup/pkg/model"

type ListTransactionRequest struct {
	CampaignID string `uri:"campaign_id" binding:"required"`
	User       model.User
}
