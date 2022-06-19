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

func TestJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	videoRepo := repository.NewVideoRepository(db)
	videoRepo.Insert(video)

	job, err := domain.NewJob("output", "pending", video)
	require.Nil(t, err)

	jobRepo := repository.NewJobRepository(db)
	jobRepo.Insert(job)

	storedJob, err := jobRepo.Find(job.ID)

	require.Nil(t, err)
	require.NotNil(t, storedJob)
	assert.Equal(t, job.ID, storedJob.ID)
}

func TestJobRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	videoRepo := repository.NewVideoRepository(db)
	videoRepo.Insert(video)

	job, err := domain.NewJob("output", "pending", video)
	require.Nil(t, err)

	jobRepo := repository.NewJobRepository(db)
	jobRepo.Insert(job)

	storedJob, err := jobRepo.Find(job.ID)
	assert.Equal(t, storedJob.Status, "pending")

	storedJob.Status = "complete"
	jobRepo.Update(storedJob)
	storedJob, err = jobRepo.Find(storedJob.ID)
	assert.Equal(t, storedJob.Status, "complete")
}
