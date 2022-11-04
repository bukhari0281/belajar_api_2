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
	_comments "github.com/ydhnwb/golang_heroku/service/comments"
)

type CommentService interface {
	All(userID string) (*[]_comments.CommentResponse, error)
	CreateComment(commentRequest dto.CreateCommentsRequest, userID string) (*_comments.CommentResponse, error)
	UpdateComment(updateCommentRequest dto.UpdateCommentsRequest, userID string) (*_comments.CommentResponse, error)
	FindOneCommentByID(commentID string) (*_comments.CommentResponse, error)
	DeleteComment(commentID string, userID string) error
}

type commentService struct {
	commentRepo repo.CommentRepository
}

func NewCommentService(commentRepo repo.CommentRepository) CommentService {
	return &commentService{
		commentRepo: commentRepo,
	}
}

// UpdateComment implements CommentService
func (c *commentService) UpdateComment(updateCommentRequest dto.UpdateCommentsRequest, userID string) (*_comments.CommentResponse, error) {
	comment, err := c.commentRepo.FindOneCommentByID(fmt.Sprintf("%d", updateCommentRequest.ID))
	if err != nil {
		return nil, err
	}

	uid, _ := strconv.ParseInt(userID, 0, 64)
	if comment.UserID != uid {
		return nil, errors.New("comment ini bukan milik anda")
	}

	comment = entity.Comments{}
	err = smapping.FillStruct(&comment, smapping.MapFields(&updateCommentRequest))

	if err != nil {
		return nil, err
	}

	comment.UserID = uid
	comment, err = c.commentRepo.UpdateComment(comment)

	if err != nil {
		return nil, err
	}

	res := _comments.NewCommentResponse(comment)
	return &res, nil
}

// All implements CommentService
func (c *commentService) All(userID string) (*[]_comments.CommentResponse, error) {
	comments, err := c.commentRepo.All(userID)
	if err != nil {
		return nil, err
	}
	comment := _comments.NewCommentArrayResponse(comments)
	return &comment, nil
}

// CreateComment implements CommentService
func (c *commentService) CreateComment(commentRequest dto.CreateCommentsRequest, userID string) (*_comments.CommentResponse, error) {
	comment := entity.Comments{}
	err := smapping.FillStruct(&comment, smapping.MapFields(&commentRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	id, _ := strconv.ParseInt(userID, 0, 64)
	comment.UserID = id
	p, err := c.commentRepo.InserComment(comment)
	if err != nil {
		return nil, err
	}

	res := _comments.NewCommentResponse(p)
	return &res, nil
}

// DeleteComment implements CommentService
func (c *commentService) DeleteComment(commentID string, userID string) error {
	comment, err := c.commentRepo.FindOneCommentByID(commentID)
	if err != nil {
		return err
	}

	if fmt.Sprintf("%d", comment.UserID) != userID {
		return errors.New("photo ini bukan milik anda")
	}

	c.commentRepo.DeleteComment(commentID)
	return nil
}

// FindOneCommentByID implements CommentService
func (c *commentService) FindOneCommentByID(commentID string) (*_comments.CommentResponse, error) {
	comment, err := c.commentRepo.FindOneCommentByID(commentID)

	if err != nil {
		return nil, err
	}

	res := _comments.NewCommentResponse(comment)
	return &res, nil
}
