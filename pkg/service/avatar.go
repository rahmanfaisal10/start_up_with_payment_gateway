package service

import (
	"bwastartup/pkg/model"
)

func (s *service) SaveAvatarService(ID string, fileLocation string) (model.User, error) {
	user, err := s.repository.GetUserByID(ID)
	if err != nil {
		return user, err
	}

	user.Avatar = fileLocation

	updatedUser, err := s.repository.UpdateUser(user)
	if err != nil {
		return user, err
	}

	return updatedUser, nil
}
