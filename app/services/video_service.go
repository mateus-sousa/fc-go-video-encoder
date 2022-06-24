package services

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/mateus-sousa/fc-go-video-encoder/app/repository"
	"github.com/mateus-sousa/fc-go-video-encoder/domain"
	"io/ioutil"
	"log"
	"os"
)

type VideoService struct {
	Video           *domain.Video
	VideoRepository repository.VideoRepository
}

func NewVideoService() VideoService {
	return VideoService{}
}

func (v *VideoService) Download(bucketPath string) error {

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}

	bkt := client.Bucket(bucketPath)
	obj := bkt.Object(v.Video.FilePath)

	r, err := obj.NewReader(ctx)
	if err != nil {
		return err
	}
	defer r.Close()

	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	f, err := os.Create(os.Getenv("LOCAL_STORAGE_PATH") + v.Video.ID + ".mp4")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(body)
	if err != nil {
		return err
	}
	log.Printf("video %v has been stored", v.Video.ID)
	return nil
}
