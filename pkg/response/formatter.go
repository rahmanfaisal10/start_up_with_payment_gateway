package response

import (
	"bwastartup/pkg/model"
	"strings"
)

//master response json
func ResponseAPI(message, status string, code int, formatter interface{}) BaseResponse {
	return BaseResponse{
		Meta: Meta{
			Message: message,
			Code:    code,
			Status:  status,
		},
		Data: formatter,
	}
}

//get first campaign response
func FormaterCampaignResponse(campaign model.Campaign) campaignResponse {
	imageURL := ""
	if len(campaign.CampaignImages) > 0 {
		imageURL = campaign.CampaignImages[0].FileName
	}

	resp := &campaignResponse{
		UUID:             campaign.UUID.String(),
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescriptions,
		ImageURL:         imageURL,
		Slug:             campaign.Slug,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
	}

	return *resp
}

//get many campign response
func FormatterAllCampaignResponse(campaigns []model.Campaign) []campaignResponse {
	data := make([]campaignResponse, 0)
	for _, v := range campaigns {
		data = append(data, FormaterCampaignResponse(v))
	}

	return data
}

//detail  campaign response
func FormatterDetailCampaignResponse(campaign model.Campaign) CampaignDetailResponse {
	imageURL := ""
	if len(campaign.CampaignImages) > 0 {
		imageURL = campaign.CampaignImages[0].FileName
	}

	userResp := &CampaignUserResponse{
		Name:     campaign.Users.Fullname,
		ImageURL: campaign.Users.Avatar,
	}

	imagesResp := make([]CampaignImageResponse, 0)
	for _, v := range campaign.CampaignImages {
		isPrimaryBool := false
		if v.IsPrimary == 1 {
			isPrimaryBool = true
		}

		imagesResp = append(imagesResp, CampaignImageResponse{
			ImageURL:  v.FileName,
			IsPrimary: isPrimaryBool,
		})
	}

	resp := &CampaignDetailResponse{
		UUID:             campaign.UUID.String(),
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescriptions,
		Description:      campaign.Descriptions,
		ImageURL:         imageURL,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
		UserID:           campaign.UserID,
		Perks:            strings.Split(campaign.Perks, ","),
		User:             *userResp,
		Images:           imagesResp,
	}

	return *resp
}

//formater response user
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
