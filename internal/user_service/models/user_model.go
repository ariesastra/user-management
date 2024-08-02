package usermodel

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id           string `json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Role         string `json:"role"`
	Organization string `json:"organization"`
}

type UserId struct {
	Id string `json:"id"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
