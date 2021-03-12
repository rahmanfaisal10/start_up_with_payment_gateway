package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID        uuid.UUID `json:"uuid" gorm:"primaryKey;type:varchar(50)" db:"uuid"`
	Fullname    string    `json:"fullname" gorm:"type:varchar(50)" db:"fullname"`
	Occupation  string    `json:"occupation" gorm:"type:varchar(191)" db:"occupation"`
	Password    string    `json:"password" gorm:"type:varchar(191)" db:"password"`
	Email       string    `json:"email" gorm:"type:varchar(191)" db:"email"`
	Avatar      string    `json:"avatar" gorm:"type:varchar(191)" db:"avatar"`
	Role        string    `json:"role" gorm:"type:varchar(191)" db:"role"`
	Token       string    `json:"token" gorm:"type:varchar(191)" db:"token"`
	CreatedBy   string    `json:"created_by" gorm:"type:varchar(191)" db:"created_by"`
	CreatedDate time.Time `json:"created_date" gorm:"type:datetime" db:"created_date"`
	UpdatedBy   string    `json:"updated_by" gorm:"type:varchar(191)" db:"updated_by"`
	UpdatedDate time.Time `json:"updated_date" gorm:"type:datetime" db:"updated_date"`
}
