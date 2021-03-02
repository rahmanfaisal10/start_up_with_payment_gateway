package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID        uuid.UUID `json:"uuid" db:"UUID"`
	Fullname    string    `json:"fullname" db:"fullname"`
	Occupation  string    `json:"occupation" db:"occupation"`
	Password    string    `json:"password" db:"password"`
	Email       string    `json:"email" db:"email"`
	Avatar      string    `json:"avatar" db:"avatar"`
	Role        string    `json:"role" db:"role"`
	Token       string    `json:"token" db:"token"`
	CreatedBy   string    `json:"created_by" db:"created_by"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
	UpdatedBy   string    `json:"updated_by" db:"updated_by"`
	UpdatedDate time.Time `json:"updated_date" db:"updated_date"`
}
