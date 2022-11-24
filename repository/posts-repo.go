package repository

import (
	"github.com/adykaaa/go-clean-rest/entity"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
