package service

import (
	"bwastartup/pkg/model"
	"bwastartup/pkg/request"
)

func (s *service) DetailCampaign(req request.DetailCampaignRequest) (model.Campaign, error) {
	campaign, err := s.repository.GetCampaignByID(req.UUID)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}
