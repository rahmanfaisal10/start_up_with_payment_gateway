package service

import (
	"bwastartup/pkg/model"
	"bwastartup/pkg/request"
	"errors"
	"time"

	"github.com/google/uuid"
)

func (s *service) CreateCampaignImage(request *request.CreateCampaignImageRequest, fileLocation string) (model.CampaignImage, error) {
	//check user create campaign image
	campaign, err := s.repository.GetCampaignByID(request.UUID)
	if err != nil {
		return model.CampaignImage{}, err
	}
	if campaign.UserID != request.User.UUID.String() {
		return model.CampaignImage{}, errors.New("This user does not have permission to create campaign images")
	}

	isPrimary := 0
	if request.IsPrimary {
		isPrimary = 1
		_, err := s.repository.MarkAllImagesAsNonPrimary(request.UUID)
		if err != nil {
			return model.CampaignImage{}, err
		}
	}

	campaignImage := &model.CampaignImage{
		UUID:        uuid.New(),
		CampaignID:  request.UUID,
		FileName:    fileLocation,
		IsPrimary:   isPrimary,
		CreatedBy:   request.User.UUID.String(),
		CreatedDate: time.Now(),
		UpdatedBy:   "",
		UpdatedDate: time.Time{},
	}

	newCampaignImage, err := s.repository.CreateCampaignImage(*campaignImage)
	if err != nil {
		return newCampaignImage, err
	}

	return newCampaignImage, nil
}
