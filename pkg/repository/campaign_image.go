package repository

import "bwastartup/pkg/model"

func (r *repository) CreateCampaignImage(campaignImage model.CampaignImage) (model.CampaignImage, error) {
	err := r.db.Create(&campaignImage).Error
	if err != nil {
		return campaignImage, err
	}
	return campaignImage, nil
}

func (r *repository) MarkAllImagesAsNonPrimary(campaignID string) (bool, error) {
	// Update with conditions
	err := r.db.Model(&model.CampaignImage{}).Where("campaign_id = ?", campaignID).Update("is_primary", 0).Error
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE active=true;
	if err != nil {
		return false, err
	}
	return true, nil
}
