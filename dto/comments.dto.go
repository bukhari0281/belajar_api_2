package dto

type CreateCommentsRequest struct {
	Message string `json:"message" form:"message" binding:"required"`
}

type UpdateCommentsRequest struct {
	ID      int64  `json:"id" form:"id"`
	Message string `json:"message" form:"message" binding:"required"`
}
