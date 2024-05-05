package service

import (
	"github.com/FlyKarlik/news_skilfactory/internal/model"
	"github.com/FlyKarlik/news_skilfactory/internal/usecases/repository"
	"github.com/FlyKarlik/news_skilfactory/internal/usecases/rss"
)

type NewsInterface interface {
	AddNews(posts []model.Post) error
	GetNews(count int) ([]model.Post, error)
}

type RssInterface interface {
	ParseURL(url string, posts chan<- []model.Post, errs chan<- error)
}

type Service struct {
	RssService  RssInterface
	NewsService NewsInterface
}

func NewService(repo *repository.Repository, rss *rss.Rss) *Service {
	return &Service{
		RssService:  NewRssService(rss.RssPost),
		NewsService: NewNewsService(repo.NewsRepo),
	}
}
