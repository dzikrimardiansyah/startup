package campaign

import "time"

type Campaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	GoalAmount       int
	CurrentAmount    int
	Perks            string
	BackerCount      int
	Slug             string
	CreatedAt        time.Time
	UpadateAt        time.Time
	CampaignImages   []CampaignImage
}

type CampaignImage struct {
	ID         int
	CampaignID int
	FIleName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpadateAt  time.Time
}
