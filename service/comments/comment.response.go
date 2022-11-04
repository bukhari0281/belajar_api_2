package _comments

import (
	"github.com/ydhnwb/golang_heroku/entity"
	_user "github.com/ydhnwb/golang_heroku/service/user"
)

type CommentResponse struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
	// Photo   _photo.PhotoResponse `json:"photo,omitempty"`
	User _user.UserResponse `json:"user,omitempty"`
}

func NewCommentResponse(comment entity.Comments) CommentResponse {
	return CommentResponse{
		ID:      comment.ID,
		Message: comment.Message,
		// Photo:   _photo.NewPhotoResponse(comment.Photo),
		User: _user.NewUserResponse(comment.User),
	}
}

func NewCommentArrayResponse(comments []entity.Comments) []CommentResponse {
	commentRes := []CommentResponse{}
	for _, v := range comments {
		p := CommentResponse{
			ID:      v.ID,
			Message: v.Message,
			// Photo:   _photo.NewPhotoResponse(v.Photo),
			User: _user.NewUserResponse(v.User),
		}
		commentRes = append(commentRes, p)

	}
	return commentRes
}
