package response

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
