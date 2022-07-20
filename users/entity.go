package users

import "time"

type User struct {
	ID             int
	Name           string
	Occupation     string
	Email          string
	Password       string
	AvatarFileName string
	Role           string
	Token          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
