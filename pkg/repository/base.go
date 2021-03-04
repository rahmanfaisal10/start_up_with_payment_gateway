package repository

import (
	"bwastartup/pkg/model"

	"gorm.io/gorm"
)

type Repository interface {
	//User
	CreateUser(user model.User) (model.User, error)
	GetUserByEmail(email string) (user model.User, err error)
	GetUserByID(ID string) (user model.User, err error)
	UpdateUser(user model.User) (model.User, error)

	//Campaign
	GetAllCampaign() (campaign []model.Campaign, err error)
	GetCampaignByUserID(userID string) (campaign []model.Campaign, err error)
	GetCampaignByID(campaignID string) (campaign model.Campaign, err error)
	CreateCampaign(campaign model.Campaign) (model.Campaign, error)
	UpdateCampaign(campaign model.Campaign) (model.Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func InitRepository(db *gorm.DB) *repository {
	return &repository{db}
}
