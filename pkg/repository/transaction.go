package repository

import "bwastartup/pkg/model"

func (r *repository) GetTransactionByCampaignID(campaignID string) (transaction []model.Transaction, err error) {
	err = r.db.Where("campaign_id = ?", campaignID).Preload("Users").Order("Id desc").Find(&transaction).Error
	if err != nil {
		return
	}
	return
}
