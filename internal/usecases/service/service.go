package service

import (
	"github.com/FlyKarlik/news_skilfactory/internal/model"
	"github.com/FlyKarlik/news_skilfactory/internal/usecases/repository"
)

type NewsInterface interface {
	AddNews(posts []model.Post) error
	GetNews(count int) ([]model.Post, error)
}

type Service struct {
	NewsService NewsInterface
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		NewsService: NewNewsService(repo.NewsRepo),
	}
}
