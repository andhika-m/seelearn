package rest

import "seelearn/internal/usecase/video"

type handler struct {
	videoUsecase video.Usecase
}

func NewVideoHandler(videoUsecase video.Usecase) *handler {
	return &handler{
		videoUsecase: videoUsecase,
	}
}
