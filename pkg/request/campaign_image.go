package request

import "bwastartup/pkg/model"

type CreateCampaignImageRequest struct {
	UUID      string `form:"uuid" binding:"required"`
	IsPrimary bool   `form:"is_primary"`
	User      model.User
}
