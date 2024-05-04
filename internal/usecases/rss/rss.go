package rss

import (
	"github.com/FlyKarlik/news_skilfactory/config"
	"github.com/FlyKarlik/news_skilfactory/internal/model"
)

type RssInterface interface {
	ParseURL(url string, posts chan<- []model.Post, errs chan<- error)
}

type Rss struct {
	RssPost RssInterface
}

func NewRss(cfg *config.Config) *Rss {
	return &Rss{
		RssPost: NewParse(cfg),
	}
}
