package service

import (
	"bwastartup/pkg/model"
	"bwastartup/pkg/request"
	"errors"
)

func (s *service) GetTransactionByCampaignID(req request.ListTransactionRequest) ([]model.Transaction, error) {
	//Authorization user
	campaign, err := s.repository.GetCampaignByID(req.CampaignID)
	if err != nil {
		return []model.Transaction{}, err
	}

	if campaign.UserID != req.User.UUID.String() {
		return []model.Transaction{}, errors.New("This user does not have permission to create campaign images")
	}

	transaction, err := s.repository.GetTransactionByCampaignID(req.CampaignID)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
