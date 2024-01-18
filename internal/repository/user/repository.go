package user

import "seelearn/internal/model"

type Repository interface {
	RegisterUser(userData model.User) (model.User, error)
	CheckRegistered(email string) (bool, error)
	GenerateUserHash(password string) (hash string, err error)
	VerifyLogin(email, password string, userData model.User) (bool, error)
	GetUserData(email string) (model.User, error)
	CreateUserSession(userID string) (model.UserSession, error)
	CheckSession(data model.UserSession) (userID string, err error)
}