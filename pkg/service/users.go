package service

import (
	"bwastartup/pkg/model"
	"bwastartup/pkg/request"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) RegisterUserService(input request.RegisterUserInput) (*model.User, error) {
	// password hash
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	users := &model.User{
		Fullname:   input.Fullname,
		Occupation: input.Occupation,
		Email:      input.Email,
		Password:   string(password),
		Role:       "user",
		CreatedAt:  time.Now(),
	}

	result, err := s.repository.CreateUser(*users)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

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

	//create jwt token
	//TODO

	return &user, nil
}

func (s *service) CheckEmailAvailabilityService(req request.CheckEmailAvailable) (bool, error) {
	user, err := s.repository.GetUserByEmail(req.Email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) SaveAvatarService(ID int, fileLocation string) (model.User, error) {
	/* dapatkan user berdasarkan ID
	update attribute avatar file name
	simpan perubahan avatar file name*/
	user, err := s.repository.GetUserByID(ID)
	if err != nil {
		return user, err
	}

	user.AvatarFilename = fileLocation

	updatedUser, err := s.repository.UpdateUser(user)
	if err != nil {
		return user, err
	}

	return updatedUser, nil
}
