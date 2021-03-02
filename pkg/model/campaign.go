package model

import (
	"time"

	"github.com/google/uuid"
)

type Campaign struct {
	UUID              uuid.UUID       `json:"uuid" db:"UUID"`
	UserID            string          `json:"user_id" db:"user_id"`
	Name              string          `json:"name" db:"name"`
	ShortDescriptions string          `json:"short_descriptions" db:"short_descriptions"`
	Descriptions      string          `json:"descriptions" db:"descriptions"`
	Perks             string          `json:"perks" db:"perks"`
	BeckerCount       float64         `json:"backer_count" db:"backer_count"`
	GoalAmount        float64         `json:"goal_amount" db:"goal_amount"`
	CurrentAmount     float64         `json:"current_amount" db:"current_amount"`
	Slug              string          `json:"slug" db:"slug"`
	CreatedBy         string          `json:"created_by" db:"created_by"`
	CreatedDate       time.Time       `json:"created_date" db:"created_date"`
	UpdatedBy         string          `json:"updated_by" db:"updated_by"`
	UpdatedDate       time.Time       `json:"updated_date" db:"updated_date"`
	CampaignImages    []CampaignImage `gorm:"foreignKey:CampaignID;references:UUID"`
	Users             User            `gorm:"foreignKey:UUID;references:UserID"`
}

type CampaignImage struct {
	UUID        uuid.UUID `json:"uuid" db:"UUID"`
	CampaignID  string    `json:"campaign_id" db:"campaign_id"`
	FileName    string    `json:"file_name" db:"file_name"`
	IsPrimary   int       `json:"is_primary" db:"is_primary"`
	CreatedBy   string    `json:"created_by" db:"created_by"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
	UpdatedBy   string    `json:"updated_by" db:"updated_by"`
	UpdatedDate time.Time `json:"updated_date" db:"updated_date"`
}
