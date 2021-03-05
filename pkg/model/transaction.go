package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          int       `json:"id" db:"id"`
	CampaignID  uuid.UUID `json:"campaign_id" db:"campaign_id"`
	UserID      uuid.UUID `json:"user_id" db:"user_id"`
	Amount      float64   `json:"amount" db:"amount"`
	Status      string    `json:"status" db:"status"`
	Code        string    `json:"code" db:"code"`
	CreatedBy   string    `json:"created_by" db:"created_by"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
	UpdatedBy   string    `json:"updated_by" db:"updated_by"`
	UpdatedDate time.Time `json:"updated_date" db:"updated_date"`
}
