package repository_test

import (
	"github.com/mateus-sousa/fc-go-video-encoder/app/repository"
	"github.com/mateus-sousa/fc-go-video-encoder/domain"
	"github.com/mateus-sousa/fc-go-video-encoder/infra/database"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repository.NewVideoRepository(db)
	repo.Insert(video)

	storedVideo, err := repo.Find(video.ID)
	require.Nil(t, err)
	require.NotNil(t, storedVideo)
	assert.Equal(t, video.ID, storedVideo.ID)
}
