package response

import "time"

type CampaignTransactionFormatter struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Amount      float64   `json:"amount"`
	CreatedDate time.Time `json:"created_date"`
}
