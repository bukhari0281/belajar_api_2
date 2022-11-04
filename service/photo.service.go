package service

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/mashingan/smapping"
	"github.com/ydhnwb/golang_heroku/dto"
	"github.com/ydhnwb/golang_heroku/entity"
	"github.com/ydhnwb/golang_heroku/repo"
	_photo "github.com/ydhnwb/golang_heroku/service/photo"
)

type PhotoService interface {
	All(userID string) (*[]_photo.PhotoResponse, error)
	CreatePhoto(photoRequest dto.CreatePhotoRequest, userID string) (*_photo.PhotoResponse, error)
	UpdatePhoto(updateProductRequest dto.UpdatePhotoRequest, userID string) (*_photo.PhotoResponse, error)
	FindOnePhotoByID(photoID string) (*_photo.PhotoResponse, error)
	DeletePhoto(photoID string, userID string) error
}

type photoService struct {
	photoRepo repo.PhotoRepository
}

func NewPhotoService(photoRepo repo.PhotoRepository) PhotoService {
	return &photoService{
		photoRepo: photoRepo,
	}
}

func (c *photoService) All(userID string) (*[]_photo.PhotoResponse, error) {
	photos, err := c.photoRepo.All(userID)
	if err != nil {
		return nil, err
	}
	phots := _photo.NewPhotoArrayResponse(photos)
	return &phots, nil
}

func (c *photoService) CreatePhoto(photoRequest dto.CreatePhotoRequest, userID string) (*_photo.PhotoResponse, error) {
	photo := entity.Photo{}
	err := smapping.FillStruct(&photo, smapping.MapFields(&photoRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	id, _ := strconv.ParseInt(userID, 0, 64)
	photo.UserID = id
	p, err := c.photoRepo.InsertPhoto(photo)
	if err != nil {
		return nil, err
	}

	res := _photo.NewPhotoResponse(p)
	return &res, nil
}

func (c *photoService) FindOnePhotoByID(photoID string) (*_photo.PhotoResponse, error) {
	photo, err := c.photoRepo.FindOnePhotoByID(photoID)

	if err != nil {
		return nil, err
	}

	res := _photo.NewPhotoResponse(photo)
	return &res, nil
}

func (c *photoService) UpdatePhoto(updatePhotoRequest dto.UpdatePhotoRequest, userID string) (*_photo.PhotoResponse, error) {
	photo, err := c.photoRepo.FindOnePhotoByID(fmt.Sprintf("%d", updatePhotoRequest.ID))
	if err != nil {
		return nil, err
	}

	uid, _ := strconv.ParseInt(userID, 0, 64)
	if photo.UserID != uid {
		return nil, errors.New("photo ini bukan milik anda")
	}

	photo = entity.Photo{}
	err = smapping.FillStruct(&photo, smapping.MapFields(&updatePhotoRequest))

	if err != nil {
		return nil, err
	}

	photo.UserID = uid
	photo, err = c.photoRepo.UpdatePhoto(photo)

	if err != nil {
		return nil, err
	}

	res := _photo.NewPhotoResponse(photo)
	return &res, nil
}

func (c *photoService) DeletePhoto(photoID string, userID string) error {
	photo, err := c.photoRepo.FindOnePhotoByID(photoID)
	if err != nil {
		return err
	}

	if fmt.Sprintf("%d", photo.UserID) != userID {
		return errors.New("photo ini bukan milik anda")
	}

	c.photoRepo.DeletePhoto(photoID)
	return nil

}
