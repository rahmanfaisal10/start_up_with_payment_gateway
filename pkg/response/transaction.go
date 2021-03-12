package response

import (
	"time"
)

type CampaignTransactionFormatter struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Amount      float64   `json:"amount"`
	CreatedDate time.Time `json:"created_date"`
}

type userTransactionFormatter struct {
	Id          int       `json:"id"`
	Amount      float64   `json:"amount"`
	Status      string    `json:"status"`
	CreatedDate time.Time `json:"created_date"`
	Campaign    struct {
		Name     string `json:"name"`
		ImageURL string `json:"image_url"`
	} `json:"campaign"`
}
