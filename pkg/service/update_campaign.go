package service

import (
	"bwastartup/pkg/model"
	"bwastartup/pkg/request"
	"errors"
	"time"
)

func (s *service) UpdateCampaign(reqID request.DetailCampaignRequest, reqData request.CreateCampaignRequest) (model.Campaign, error) {
	campaign, err := s.repository.GetCampaignByID(reqID.UUID)
	if err != nil {
		return campaign, err
	}

	if reqData.User.UUID.String() != campaign.UserID {
		return campaign, errors.New("Not an owner of the campaign")
	}

	campaign = model.Campaign{
		UUID:              campaign.UUID,
		UserID:            campaign.UserID,
		Name:              reqData.Name,
		ShortDescriptions: reqData.ShortDescription,
		Descriptions:      reqData.Description,
		Perks:             reqData.Perks,
		BackerCount:       campaign.BackerCount,
		GoalAmount:        reqData.GoalAmount,
		CurrentAmount:     campaign.CurrentAmount,
		Slug:              campaign.Slug,
		CreatedBy:         campaign.CreatedBy,
		CreatedDate:       campaign.CreatedDate,
		UpdatedBy:         reqData.User.Email,
		UpdatedDate:       time.Now(),
		CampaignImages:    campaign.CampaignImages,
		Users:             campaign.Users,
	}

	updateCampaign, err := s.repository.UpdateCampaign(campaign)
	if err != nil {
		return updateCampaign, err
	}
	return updateCampaign, nil
}
