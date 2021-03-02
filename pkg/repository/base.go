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
	GetAll() (campaign []model.Campaign, err error)
	GetByUserID(userID string) (campaign []model.Campaign, err error)
}

type repository struct {
	db *gorm.DB
}

func InitRepository(db *gorm.DB) *repository {
	return &repository{db}
}
