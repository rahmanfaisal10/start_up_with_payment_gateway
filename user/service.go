package user

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUserService(input RegisterUserInput) (*User, error)
	LoginUserService(req LoginUserInput) (*User, error)
	CheckEmailAvailabilityService(req CheckEmailAvailable) (bool, error)
}

type service struct {
	repository Repository
}

func InitService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUserService(input RegisterUserInput) (*User, error) {
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

func (s *service) LoginUserService(req LoginUserInput) (*User, error) {
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

	//create jwt token
	//TODO

	return &user, nil
}

func (s *service) CheckEmailAvailabilityService(req CheckEmailAvailable) (bool, error) {
	user, err := s.repository.GetUserByEmail(req.Email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}
