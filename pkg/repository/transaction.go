package repository

import "bwastartup/pkg/model"

func (r *repository) GetTransactionByCampaignID(campaignID string) (transaction []model.Transaction, err error) {
	err = r.db.Where("campaign_id = ?", campaignID).Preload("Users").Order("Id desc").Find(&transaction).Error
	if err != nil {
		return
	}
	return
}

func (r *repository) GetTransactionByUserID(userID string) (transaction []model.Transaction, err error) {
	err = r.db.Where("user_id = ?", userID).Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Order("Id desc").Find(&transaction).Error
	if err != nil {
		return
	}
	return
}
