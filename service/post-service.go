package service

import (
	"errors"
	"math/rand"

	"github.com/adykaaa/go-clean-rest/entity"
	"github.com/adykaaa/go-clean-rest/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

var (
	repo repository.PostRepository = repository.NewFirestoreRepository()
)

func NewPostService() PostService {
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		return errors.New("The post is nil!")
	}

	if post.Title == "" {
		return errors.New("The post title is empty!")
	}
	return nil
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}

func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
