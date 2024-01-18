package video

import (
	"errors"
	"seelearn/internal/model"
	"seelearn/internal/repository/user"
	"seelearn/internal/repository/video"
)

type videoUsecase struct {
	videoRepo video.Repository
	userRepo  user.Repository
}

func NewVideoUsecase(videoRepo video.Repository, userRepo user.Repository) Usecase {
	return &videoUsecase{
		videoRepo: videoRepo,
		userRepo:  userRepo,
	}
}

func (m *videoUsecase) GetVideoByID(videoID string) (model.VideoDescription, error) {
	return m.videoRepo.GetVideoDataByID(videoID)
}

func (m *videoUsecase) GetVideoList(videoKategori, videoKelas string) ([]model.VideoDescription, error) {
	videoData, err := m.videoRepo.GetVideoList(videoKategori, videoKelas)
	if err != nil {
		return nil, err
	}

	for i, video := range videoData {
		videoFile, err := m.videoRepo.GetVideoFileByID(video.VideoFileID)
		if err != nil {
			return nil, err
		}

		videoData[i].VideoFile = []model.VideoFile{videoFile}
	}

	return videoData, nil
}

func (m *videoUsecase) VideoDescription(request model.VideoDescription) (model.VideoDescription, error) {

	createdVideoDescription, err := m.videoRepo.CreateVideoDescription(request)
	if err != nil {
		return model.VideoDescription{}, err
	}

	return createdVideoDescription, nil
}

func (m *videoUsecase) GetVideoFilePathByID(videoID string) (string, error) {
	return m.videoRepo.GetVideoFilePathByID(videoID)
}

func (m *videoUsecase) GetFileNameByID(id string) (string, error) {
	return m.videoRepo.GetFileNameByID(id)
}

func (m *videoUsecase) VideoCreate(request model.VideoFile) (model.VideoFile, error) {

	createdVideoFile, err := m.videoRepo.CreateVideo(request)
	if err != nil {
		return model.VideoFile{}, err
	}

	return createdVideoFile, nil
}

func (m *videoUsecase) UpdateVideo(video model.VideoDescription) (model.VideoDescription, error) {
	return m.videoRepo.UpdateVideo(video)
}

func (m *videoUsecase) DeleteVideo(id string) error {
	return m.videoRepo.DeleteVideo(id)
}

func (vu *videoUsecase) GetVideos(request model.VideoFile) (model.VideoDescription, error) {
	videoFile, err := vu.videoRepo.GetVideosWithDescriptions(request.ID)
	if err != nil {
		return videoFile, err
	}

	if videoFile.VideoFileID != request.ID {
		return model.VideoDescription{}, errors.New("unauthorized")
	}

	if err != nil {
		return videoFile, err
	}

	return videoFile, nil
}
