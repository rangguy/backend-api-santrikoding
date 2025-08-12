package model

import "time"

type User struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Username  string    `json:"username" gorm:"not null;unique"`
	Email     string    `json:"email" gorm:"not null;unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
