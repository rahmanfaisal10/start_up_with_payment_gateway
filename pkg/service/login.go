package service

import (
	"bwastartup/pkg/model"
	"bwastartup/pkg/request"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) LoginUserService(req request.LoginUserInput) (*model.User, error) {
	//get data user from database table user
	user, err := s.repository.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	//compare hash password database with request
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *service) GetUserByIDService(ID string) (model.User, error) {
	user, err := s.repository.GetUserByID(ID)
	if err != nil {
		return user, err
	}

	if user.Email == "" {
		return user, errors.New("No User found on that email")
	}

	return user, nil
}
