package repository

import (
	"../entity"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindOne(title string) (*entity.Post, error)
}

// New PostRepository
func NewPostRepository() PostRepository {
	return &repo{}
}