package service

import (
	"bwastartup/pkg/request"
)

func (s *service) CheckEmailAvailabilityService(req request.CheckEmailAvailable) (bool, error) {
	_, err := s.repository.GetUserByEmail(req.Email)
	if err != nil {
		if err.Error() == "record not found" {
			return true, nil
		}
		return false, err
	}

	return false, nil
}
