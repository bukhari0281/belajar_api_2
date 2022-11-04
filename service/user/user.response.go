package _user

import "github.com/ydhnwb/golang_heroku/entity"

type UserResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int64  `json:"age"`
	Token string `json:"token,omitempty"`
}

func NewUserResponse(user entity.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
		Age:   user.Age,
	}
}
