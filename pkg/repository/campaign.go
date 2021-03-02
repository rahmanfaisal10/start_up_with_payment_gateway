package repository

import "bwastartup/pkg/model"

func (r *repository) GetAll() (campaign []model.Campaign, err error) {
	err = r.db.Find(&campaign).Error
	if err != nil {
		return
	}
	return
}

func (r *repository) GetByUserID(userID string) (campaign []model.Campaign, err error) {
	err = r.db.Where("user_id = ?", userID).Find(&campaign).Error
	if err != nil {
		return
	}
	return
}
