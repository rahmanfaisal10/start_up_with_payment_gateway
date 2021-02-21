package user

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (*User, error)
}

type service struct {
	repository Repository
}

func InitService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (*User, error) {
	// password hash
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	users := &User{
		Fullname:   input.Fullname,
		Occupation: input.Occupation,
		Email:      input.Email,
		Password:   string(password),
		Role:       "user",
		CreatedAt:  time.Now(),
	}

	result, err := s.repository.createUser(*users)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
