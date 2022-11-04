package repo

import (
	"github.com/ydhnwb/golang_heroku/entity"
	"gorm.io/gorm"
)

type PhotoRepository interface {
	All(userID string) ([]entity.Photo, error)
	InsertPhoto(photo entity.Photo) (entity.Photo, error)
	UpdatePhoto(photo entity.Photo) (entity.Photo, error)
	DeletePhoto(photoID string) error
	FindOnePhotoByID(ID string) (entity.Photo, error)
	FindAllPhoto(userID string) ([]entity.Photo, error)
}

type photoRepo struct {
	connection *gorm.DB
}

func NewPhotoRepo(connection *gorm.DB) PhotoRepository {
	return &photoRepo{
		connection: connection,
	}
}

func (c *photoRepo) All(userID string) ([]entity.Photo, error) {
	photos := []entity.Photo{}
	c.connection.Preload("User").Where("user_id = ?", userID).Find(&photos)
	return photos, nil
}

func (c *photoRepo) InsertPhoto(photo entity.Photo) (entity.Photo, error) {
	c.connection.Save(&photo)
	c.connection.Preload("User").Find(&photo)
	return photo, nil
}

func (c *photoRepo) UpdatePhoto(photo entity.Photo) (entity.Photo, error) {
	c.connection.Save(&photo)
	c.connection.Preload("User").Find(&photo)
	return photo, nil
}

func (c *photoRepo) FindOnePhotoByID(photoID string) (entity.Photo, error) {
	var photo entity.Photo
	res := c.connection.Preload("User").Where("id = ?", photoID).Take(&photo)
	if res.Error != nil {
		return photo, res.Error
	}
	return photo, nil
}

func (c *photoRepo) FindAllPhoto(userID string) ([]entity.Photo, error) {
	photo := []entity.Photo{}
	c.connection.Where("user_id = ?", userID).Find(&photo)
	return photo, nil
}

func (c *photoRepo) DeletePhoto(photoID string) error {
	var photo entity.Photo
	res := c.connection.Preload("User").Where("id = ?", photoID).Take(&photo)
	if res.Error != nil {
		return res.Error
	}
	c.connection.Delete(&photo)
	return nil
}
