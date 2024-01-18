package video

import (
	"errors"
	"seelearn/internal/model"

	"github.com/google/uuid"
	"seelearn/internal/model/constant"
)

func (vr *videoUsecase) RegisterUser(request model.RegisterRequest) (model.User, error) {
	userRegistered, err := vr.userRepo.CheckRegistered(request.Email)
	if err != nil {
		return model.User{}, err
	}
	if userRegistered {
		return model.User{}, errors.New("email already registered")
	}

	userHash, err := vr.userRepo.GenerateUserHash(request.Password)
	if err != nil {
		return model.User{}, err
	}

	userData, err := vr.userRepo.RegisterUser(model.User{
		ID:     uuid.New().String(),
		Name:   request.Name,
		School: request.School,
		Email:  request.Email,
		Hash:   userHash,
		Role:   constant.RoleCategoryTeacher,
	})

	if err != nil {
		return model.User{}, err
	}

	return userData, nil
}

func (vr *videoUsecase) Login(request model.LoginRequest) (model.UserSession, error) {
	userData, err := vr.userRepo.GetUserData(request.Email)
	if err != nil {
		return model.UserSession{}, err
	}

	verified, err := vr.userRepo.VerifyLogin(request.Email, request.Password, userData)
	if err != nil {
		return model.UserSession{}, err
	}

	if !verified {
		return model.UserSession{}, errors.New("can't verify user login")
	}

	userSession, err := vr.userRepo.CreateUserSession(userData.ID)
	if err != nil {
		return model.UserSession{}, err
	}

	return userSession, nil
}

func (vr *videoUsecase) CheckSession(data model.UserSession) (userID string, err error) {
	userID, err = vr.userRepo.CheckSession(data)
	if err != nil {
		return "", err
	}

	return userID, nil
}