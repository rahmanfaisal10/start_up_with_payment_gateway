package repository

import "bwastartup/pkg/model"

func (r *repository) GetAllCampaign() (campaign []model.Campaign, err error) {
	err = r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaign).Error
	if err != nil {
		return
	}
	return
}

func (r *repository) GetCampaignByUserID(userID string) (campaign []model.Campaign, err error) {
	err = r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaign).Error
	if err != nil {
		return
	}
	return
}
