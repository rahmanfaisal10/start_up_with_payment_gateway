package response

import (
	"bwastartup/pkg/model"
	"strings"
)

type campaignResponse struct {
	UUID             string  `json:"uuid"`
	UserID           string  `json:"user_id"`
	Name             string  `json:"name"`
	ShortDescription string  `json:"short_description"`
	ImageURL         string  `json:"image_url"`
	GoalAmount       float64 `json:"goal_amount"`
	CurrentAmount    float64 `json:"current_amount"`
	Slug             string  `json:"slug"`
}

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

func FormatterAllCampaignResponse(campaigns []model.Campaign) []campaignResponse {
	data := make([]campaignResponse, 0)
	for _, v := range campaigns {
		data = append(data, FormaterCampaignResponse(v))
	}

	return data
}

type CampaignDetailResponse struct {
	UUID             string                  `json:"uuid"`
	Name             string                  `json:"name"`
	ShortDescription string                  `json:"short_description"`
	Description      string                  `json:"description"`
	ImageURL         string                  `json:"image_url"`
	GoalAmount       float64                 `json:"goal_amount"`
	CurrentAmount    float64                 `json:"current_amount"`
	UserID           string                  `json:"user_id"`
	Slug             string                  `json:"slug"`
	Perks            []string                `json:"perks"`
	User             CampaignUserResponse    `json:"user"`
	Images           []CampaignImageResponse `json:"images"`
}

type CampaignUserResponse struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImageResponse struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

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
