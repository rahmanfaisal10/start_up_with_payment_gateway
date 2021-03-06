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

func (r *repository) GetCampaignByID(campaignID string) (campaign model.Campaign, err error) {
	err = r.db.Where("uuid = ?", campaignID).Preload("Users").Preload("CampaignImages").First(&campaign).Error
	if err != nil {
		return
	}
	return
}

func (r *repository) CreateCampaign(campaign model.Campaign) (model.Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) UpdateCampaign(campaign model.Campaign) (model.Campaign, error) {
	err := r.db.Where("UUID=?", campaign.UUID).Save(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, err
}
