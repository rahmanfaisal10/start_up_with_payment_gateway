package repository

import "bwastartup/pkg/model"

func (r *repository) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) GetUserByEmail(email string) (user model.User, err error) {
	err = r.db.Where("email=?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) GetUserByID(ID int) (user model.User, err error) {
	err = r.db.Where("id=?", ID).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) UpdateUser(user model.User) (model.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
