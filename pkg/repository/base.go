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

	//campaign image
	CreateCampaignImage(campaignImage model.CampaignImage) (model.CampaignImage, error)
	MarkAllImagesAsNonPrimary(campaignID string) (bool, error)

	//Transaction
	GetTransactionByCampaignID(campaignID string) (transaction []model.Transaction, err error)
	GetTransactionByUserID(userID string) (transaction []model.Transaction, err error)
}

type repository struct {
	db *gorm.DB
}

func InitRepository(db *gorm.DB) *repository {
	return &repository{db}
}
