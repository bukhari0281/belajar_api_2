package _photo

import (
	"github.com/ydhnwb/golang_heroku/entity"
	_user "github.com/ydhnwb/golang_heroku/service/user"
)

type PhotoResponse struct {
	ID        int64              `json:"id"`
	Title     string             `json:"title"`
	Caption   string             `json:"caption"`
	Photo_url string             `json:"photo_url"`
	User      _user.UserResponse `json:"user,omitempty"`
}

func NewPhotoResponse(photo entity.Photo) PhotoResponse {
	return PhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		Photo_url: photo.Photo_url,
		User:      _user.NewUserResponse(photo.User),
	}
}

func NewPhotoArrayResponse(photos []entity.Photo) []PhotoResponse {
	photoRes := []PhotoResponse{}
	for _, v := range photos {
		p := PhotoResponse{
			ID:        v.ID,
			Title:     v.Title,
			Caption:   v.Caption,
			Photo_url: v.Photo_url,
			User:      _user.NewUserResponse(v.User),
		}
		photoRes = append(photoRes, p)
	}
	return photoRes
}
