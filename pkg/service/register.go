package service

import (
	"bwastartup/pkg/model"
	"bwastartup/pkg/request"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) RegisterUserService(input request.RegisterUserInput) (*model.User, error) {
	// password hash
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	users := &model.User{
		UUID:        uuid.New(),
		Fullname:    input.Fullname,
		Occupation:  input.Occupation,
		Email:       input.Email,
		Password:    string(password),
		Role:        "admin",
		CreatedDate: time.Now(),
		CreatedBy:   "admin",
	}

	result, err := s.repository.CreateUser(*users)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
