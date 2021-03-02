package model

import "time"

type Campaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BeckerCount      int
	GoalAmount       int
	CurrentAmount    int
	slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type CampaignImage struct {
	ID         int
	CampaignID int
	Filename   string
	isPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
