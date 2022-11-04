package repo

import (
	"github.com/ydhnwb/golang_heroku/entity"
	"gorm.io/gorm"
)

type CommentRepository interface {
	All(userID string) ([]entity.Comments, error)
	InserComment(comment entity.Comments) (entity.Comments, error)
	UpdateComment(comment entity.Comments) (entity.Comments, error)
	DeleteComment(commentID string) error
	FindOneCommentByID(ID string) (entity.Comments, error)
	FindAllComment(userID string) ([]entity.Comments, error)
}

type commentRepo struct {
	connection *gorm.DB
}

func NewCommentRepo(connection *gorm.DB) CommentRepository {
	return &commentRepo{
		connection: connection,
	}
}

// All implements CommentRepository
func (c *commentRepo) All(userID string) ([]entity.Comments, error) {
	comments := []entity.Comments{}
	c.connection.Preload("User").Where("user_id", userID).Find(&comments)
	return comments, nil
}

// DeleteComment implements CommentRepository

// FindAllComment implements CommentRepository
func (c *commentRepo) FindAllComment(userID string) ([]entity.Comments, error) {
	comments := []entity.Comments{}
	c.connection.Where("user_id = ?", userID).Find(&comments)
	return comments, nil
}

// FindOneCommentByID implements CommentRepository
func (c *commentRepo) FindOneCommentByID(commentID string) (entity.Comments, error) {
	var comment entity.Comments
	res := c.connection.Preload("User").Where("id = ?", commentID).Take(&comment)
	if res.Error != nil {
		return comment, res.Error
	}
	return comment, nil
}

// InserComment implements CommentRepository
func (c *commentRepo) InserComment(comment entity.Comments) (entity.Comments, error) {
	c.connection.Save(&comment)
	c.connection.Preload("User").Find(&comment)
	return comment, nil
}

// UpdateComment implements CommentRepository
func (c *commentRepo) UpdateComment(comment entity.Comments) (entity.Comments, error) {
	c.connection.Save(&comment)
	c.connection.Preload("User").Find(&comment)
	return comment, nil
}

func (c *commentRepo) DeleteComment(commentID string) error {
	var comment entity.Comments
	res := c.connection.Preload("User").Where("id = ?", commentID).Take(&comment)
	if res.Error != nil {
		return res.Error
	}
	c.connection.Delete(&comment)
	return nil
}
