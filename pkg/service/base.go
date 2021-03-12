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
	SaveAvatarService(ID string, fileLocation string) (model.User, error)
	GetUserByIDService(ID string) (model.User, error)

	ListCampaign(userID string) (campaign []model.Campaign, err error)
	DetailCampaign(uuid request.DetailCampaignRequest) (model.Campaign, error)
	Createcampaign(request request.CreateCampaignRequest) (model.Campaign, error)
	UpdateCampaign(reqID request.DetailCampaignRequest, reqData request.CreateCampaignRequest) (model.Campaign, error)

	CreateCampaignImage(request *request.CreateCampaignImageRequest, fileLocation string) (model.CampaignImage, error)

	GetTransactionByCampaignID(campaignID request.ListTransactionRequest) ([]model.Transaction, error)
	GetTransactionByUserID(request request.ListUserTransactionRequest) ([]model.Transaction, error)
}

type service struct {
	repository repository.Repository
}

func InitService(repository repository.Repository) *service {
	return &service{repository}
}
