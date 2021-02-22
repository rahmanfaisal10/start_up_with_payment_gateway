package repository

import (
	"bwastartup/pkg/model"

	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user model.User) (model.User, error)
	GetUserByEmail(email string) (user model.User, err error)
	GetUserByID(ID int) (user model.User, err error)
	UpdateUser(user model.User) (model.User, error)
}

type repository struct {
	db *gorm.DB
}

func InitRepository(db *gorm.DB) *repository {
	return &repository{db}
}
