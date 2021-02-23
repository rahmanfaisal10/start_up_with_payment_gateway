package response

import "bwastartup/pkg/model"

type registerResponse struct {
	Fullname   string `json:"fullname"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Token      string `json:"token"`
}

func FormaterUserResponse(user model.User, token string) registerResponse {
	resp := &registerResponse{
		Fullname:   user.Fullname,
		Occupation: user.Occupation,
		Email:      user.Email,
		Password:   user.Password,
		Token:      token,
	}
	return *resp
}
