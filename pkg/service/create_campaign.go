package service

import (
	"bwastartup/pkg/model"
	"bwastartup/pkg/request"
	"fmt"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

func (s *service) Createcampaign(request request.CreateCampaignRequest) (model.Campaign, error) {

	slugCandidate := fmt.Sprintf("%s-%s", request.Name, request.User.UUID)
	// campaign := new(model.Campaign)
	campaign := &model.Campaign{
		UUID:              uuid.New(),
		UserID:            request.User.UUID.String(),
		Name:              request.Name,
		ShortDescriptions: request.ShortDescription,
		Descriptions:      request.Description,
		Perks:             request.Perks,
		GoalAmount:        request.GoalAmount,
		Slug:              slug.Make(slugCandidate),
	}

	//insert data to database
	newCampaign, err := s.repository.CreateCampaign(*campaign)
	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil
}
