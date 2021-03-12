package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          int       `json:"id" gorm:"primaryKey;type:int;autoIncrement" db:"id"`
	CampaignID  uuid.UUID `json:"campaign_id" gorm:"type:varchar(50)" db:"campaign_id"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:varchar(50)" db:"user_id"`
	Name        string    `json:"name" gorm:"type:varchar(50)" db:"name"`
	Amount      float64   `json:"amount" gorm:"type:double" db:"amount"`
	Status      string    `json:"status" gorm:"type:varchar(50)" db:"status"`
	Code        string    `json:"code" gorm:"type:varchar(50)" db:"code"`
	Users       User      `gorm:"foreignKey:UserID;references:UUID"`
	Campaign    Campaign  `gorm:"foreignKey:CampaignID;references:UUID"`
	CreatedBy   string    `json:"created_by" gorm:"type:varchar(50)" db:"created_by"`
	CreatedDate time.Time `json:"created_date" gorm:"type:datetime" db:"created_date"`
	UpdatedBy   string    `json:"updated_by" gorm:"type:varchar(50)" db:"updated_by"`
	UpdatedDate time.Time `json:"updated_date" gorm:"type:datetime" db:"updated_date"`
}
