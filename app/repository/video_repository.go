package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/mateus-sousa/fc-go-video-encoder/domain"
	uuid "github.com/satori/go.uuid"
)

type VideoRepository interface {
	Insert(video *domain.Video) (*domain.Video, error)
	Find(id string) (*domain.Video, error)
}

type VideoRepositoryDb struct {
	Db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *VideoRepositoryDb {
	return &VideoRepositoryDb{Db: db}
}

func (v *VideoRepositoryDb) Insert(video *domain.Video) (*domain.Video, error) {
	if video.ID == "" {
		video.ID = uuid.NewV4().String()
	}
	err := v.Db.Create(video).Error
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (v *VideoRepositoryDb) Find(id string) (*domain.Video, error) {
	var video domain.Video
	v.Db.First(&video, "id = ?", id)
	if video.ID == "" {
		return nil, fmt.Errorf("video not found")
	}
	return &video, nil
}
