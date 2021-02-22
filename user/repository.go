package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	createUser(user User) (User, error)
	GetUserByEmail(email string) (user User, err error)
}

type repository struct {
	db *gorm.DB
}

func InitRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) createUser(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) GetUserByEmail(email string) (user User, err error) {
	err = r.db.Where("email=?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
