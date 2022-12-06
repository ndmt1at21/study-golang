package models

import "time"

type User struct {
	Id        int64
	Email     string
	Name      string
	Password  string
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type CreateUserData struct {
	Email    string
	Name     string
	Password string
}
