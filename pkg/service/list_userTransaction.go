package service

import (
	"bwastartup/pkg/model"
	"bwastartup/pkg/request"
)

func (s *service) GetTransactionByUserID(request request.ListUserTransactionRequest) ([]model.Transaction, error) {
	transaction, err := s.repository.GetTransactionByUserID(request.UserID)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}
