package model

import (
	"time"

	"github.com/google/uuid"
)

type Campaign struct {
	UUID              uuid.UUID       `json:"uuid" gorm:"primaryKey;type:varchar(50)" db:"uuid"`
	UserID            string          `json:"user_id" gorm:"type:varchar(50)" db:"user_id"`
	Name              string          `json:"name" gorm:"type:varchar(50)" db:"name"`
	ShortDescriptions string          `json:"short_descriptions" gorm:"type:varchar(100)" db:"short_descriptions"`
	Descriptions      string          `json:"descriptions" gorm:"type:longtext" db:"descriptions"`
	Perks             string          `json:"perks" gorm:"type:varchar(191)" db:"perks"`
	BackerCount       float64         `json:"backer_count" gorm:"type:double" db:"backer_count"`
	GoalAmount        float64         `json:"goal_amount" gorm:"type:double" db:"goal_amount"`
	CurrentAmount     float64         `json:"current_amount" gorm:"type:double" db:"current_amount"`
	Slug              string          `json:"slug" gorm:"type:varchar(191)" db:"slug"`
	CreatedBy         string          `json:"created_by" gorm:"type:varchar(191)" db:"created_by"`
	CreatedDate       time.Time       `json:"created_date" gorm:"type:datetime" db:"created_date"`
	UpdatedBy         string          `json:"updated_by" gorm:"type:varchar(191)" db:"updated_by"`
	UpdatedDate       time.Time       `json:"updated_date" gorm:"type:datetime" db:"updated_date"`
	CampaignImages    []CampaignImage `gorm:"foreignkey:CampaignID;references:UUID"`
	Users             User            `gorm:"foreignKey:UserID;references:UUID"`
}

type CampaignImage struct {
	UUID        uuid.UUID `json:"uuid" gorm:"primaryKey;type:varchar(50)" db:"uuid"`
	CampaignID  string    `json:"campaign_id" gorm:"type:varchar(50)" db:"campaign_id"`
	FileName    string    `json:"file_name" gorm:"type:varchar(50)" db:"file_name"`
	IsPrimary   int       `json:"is_primary" gorm:"type:tinyint(1);default:0" db:"is_primary"`
	CreatedBy   string    `json:"created_by" gorm:"type:varchar(191)" db:"created_by"`
	CreatedDate time.Time `json:"created_date" gorm:"type:datetime" db:"created_date"`
	UpdatedBy   string    `json:"updated_by" gorm:"type:varchar(191)" db:"updated_by"`
	UpdatedDate time.Time `json:"updated_date" gorm:"type:datetime" db:"updated_date"`
}
