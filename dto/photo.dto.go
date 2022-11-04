package dto

type CreatePhotoRequest struct {
	Title     string `json:"title" form:"name" binding:"required,min=1"`
	Caption   string `json:"caption" form:"name" binding:"required"`
	Photo_url string `json:"photo_url" form:"name" binding:"required"`
}

type UpdatePhotoRequest struct {
	ID        int64  `json:"id" form:"id"`
	Title     string `json:"title" form:"name" binding:"required,min=1"`
	Caption   string `json:"caption" form:"name" binding:"required"`
	Photo_url string `json:"photo_url" form:"name" binding:"required"`
}
