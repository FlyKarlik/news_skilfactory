package service

import (
	"github.com/FlyKarlik/news_skilfactory/internal/model"
	"github.com/FlyKarlik/news_skilfactory/internal/usecases/rss"
)

type RssService struct {
	rss rss.RssInterface
}

func NewRssService(rss rss.RssInterface) *RssService {
	return &RssService{rss: rss}
}

func (r *RssService) ParseURL(url string, posts chan<- []model.Post, errs chan<- error) {
	r.rss.ParseURL(url, posts, errs)
}
