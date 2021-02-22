package model

import "time"

type User struct {
	ID             int32     `json:"id" db:"id"`
	Fullname       string    `json:"fullname" db:"fullname"`
	Occupation     string    `json:"occupation" db:"occupation"`
	Password       string    `json:"password" db:"password"`
	Email          string    `json:"email" db:"email"`
	AvatarFilename string    `json:"avatar_filename" db:"avatar_filename"`
	Role           string    `json:"role" db:"role"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}
