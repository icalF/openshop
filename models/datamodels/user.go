package datamodels

import "time"

type User struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" validate:"max=40,excludesall=!@#$%^&*()_+-=:;?/0x2C"`
	Token     string    `json:"token" gorm:"unique"`
	Email     string    `json:"email" validate:"max=40,email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(token string) User {
	return User{
		Token: token,
	}
}