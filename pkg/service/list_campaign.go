package service

import "bwastartup/pkg/model"

func (s *service) ListCampaign(userID string) (campaign []model.Campaign, err error) {
	if userID != "" {
		campaign, err = s.repository.GetCampaignByUserID(userID)
		if err != nil {
			return
		}
		return
	}

	campaign, err = s.repository.GetAllCampaign()
	if err != nil {
		return
	}
	return
}
