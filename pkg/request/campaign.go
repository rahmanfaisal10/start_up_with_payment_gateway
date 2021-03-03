package request

import "bwastartup/pkg/model"

type DetailCampaignRequest struct {
	UUID string `uri:"uuid" binding:"required"`
}

type CreateCampaignRequest struct {
	Name             string  `json:"name" binding:"required"`
	ShortDescription string  `json:"short_description" binding:"required"`
	Description      string  `json:"description" binding:"required"`
	GoalAmount       float64 `json:"goal_amount" binding:"required"`
	Perks            string  `json:"perks" binding:"required"`
	User             model.User
}
