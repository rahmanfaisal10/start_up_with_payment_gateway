package service

import (
	"bwastartup/pkg/model"
	"bwastartup/pkg/repository"
	"bwastartup/pkg/request"
)

type Service interface {
	RegisterUserService(input request.RegisterUserInput) (*model.User, error)
	LoginUserService(req request.LoginUserInput) (*model.User, error)
	CheckEmailAvailabilityService(req request.CheckEmailAvailable) (bool, error)
	SaveAvatarService(ID int, fileLocation string) (model.User, error)
	GetUserByIDService(ID int) (model.User, error)
}

type service struct {
	repository repository.Repository
}

func InitService(repository repository.Repository) *service {
	return &service{repository}
}
