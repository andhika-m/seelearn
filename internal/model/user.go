package model

import (
	"seelearn/internal/model/constant"
)

type User struct {
	ID     string `gorm:"primaryKey" json:"user_id"`
	Name   string `json:"name"`
	School string `json:"school"`
	Email  string `gorm:"unique" json:"email"`
	Hash   string `json:"-"`
	Role   string `json:"role"` // "teacher" atau "student"
}

type RegisterRequest struct {
	Name     string                `json:"name"`
	School   string                `json:"school"`
	Email    string                `json:"email"`
	Password string                `json:"password"`
	Role     constant.RoleCategory `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSession struct {
	JWTToken string `json:"jwt_token"`
}
