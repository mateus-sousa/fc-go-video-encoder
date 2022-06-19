package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/mateus-sousa/fc-go-video-encoder/domain"
)

type JobRepository interface {
	Insert(job *domain.Job) (*domain.Job, error)
	Update(job *domain.Job) (*domain.Job, error)
	Find(id string) (*domain.Job, error)
}

type JobRepositoryDb struct {
	Db *gorm.DB
}

func NewJobRepository(db *gorm.DB) *JobRepositoryDb {
	return &JobRepositoryDb{Db: db}
}

func (j *JobRepositoryDb) Insert(job *domain.Job) (*domain.Job, error) {
	err := j.Db.Create(job).Error
	if err != nil {
		return nil, err
	}
	return job, nil
}

func (j *JobRepositoryDb) Update(job *domain.Job) (*domain.Job, error) {
	err := j.Db.Save(job).Error
	if err != nil {
		return nil, err
	}
	return job, nil
}

func (j *JobRepositoryDb) Find(id string) (*domain.Job, error) {
	var job domain.Job
	j.Db.First(&job, "id = ?", id)
	if job.ID == "" {
		return nil, fmt.Errorf("job not found")
	}
	return &job, nil
}
