package service

import (
	"github.com/FlyKarlik/news_skilfactory/internal/model"
	"github.com/FlyKarlik/news_skilfactory/internal/usecases/repository"
)

type NewsService struct {
	repo repository.NewsInterface
}

func NewNewsService(repo repository.NewsInterface) *NewsService {
	return &NewsService{repo: repo}
}

func (s *NewsService) AddNews(posts []model.Post) error {
	return s.repo.AddNews(posts)
}

func (s *NewsService) GetNews(count int) ([]model.Post, error) {
	return s.repo.GetNews(count)
}
