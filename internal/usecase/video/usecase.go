package video

import (
	"seelearn/internal/model"
)

type Usecase interface {
	VideoDescription(request model.VideoDescription) (model.VideoDescription, error)
	VideoCreate(request model.VideoFile) (model.VideoFile, error)

	//use
	GetVideos(request model.VideoFile) (model.VideoDescription, error)
	GetVideoList(videoKategori, videoKelas string) ([]model.VideoDescription, error)
	GetVideoByID(videoID string) (model.VideoDescription, error)
	// try
	GetVideoFilePathByID(videoID string) (string, error)
	GetFileNameByID(id string) (string, error)
	UpdateVideo(video model.VideoDescription) (model.VideoDescription, error)
	DeleteVideo(id string) (error)

	//auth
	RegisterUser(request model.RegisterRequest) (model.User, error)
	Login(request model.LoginRequest) (model.UserSession, error)
	CheckSession(data model.UserSession) (userID string, err error)
	// CheckSession(sessionData model.UserSession) (userID string, err error)
}
