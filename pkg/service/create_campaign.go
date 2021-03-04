package service

import (
	"bwastartup/pkg/model"
	"bwastartup/pkg/request"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

func (s *service) Createcampaign(request request.CreateCampaignRequest) (model.Campaign, error) {

	slugCandidate := fmt.Sprintf("%s-%s", request.Name, request.User.Email)
	campaign := &model.Campaign{
		UUID:              uuid.New(),
		UserID:            request.User.UUID.String(),
		Name:              request.Name,
		ShortDescriptions: request.ShortDescription,
		Descriptions:      request.Description,
		Perks:             request.Perks,
		BackerCount:       0,
		GoalAmount:        request.GoalAmount,
		CurrentAmount:     0,
		Slug:              slug.Make(slugCandidate),
		CreatedBy:         request.User.Fullname,
		CreatedDate:       time.Now(),
		UpdatedBy:         "",
		UpdatedDate:       time.Time{},
		CampaignImages:    []model.CampaignImage{},
		Users:             model.User{},
	}

	//insert data to database
	newCampaign, err := s.repository.CreateCampaign(*campaign)
	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil
}
