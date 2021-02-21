package user

import "gorm.io/gorm"

type Repository interface {
	createUser(user User) (User, error)
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
