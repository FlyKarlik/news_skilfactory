package repository

import (
	"github.com/FlyKarlik/news_skilfactory/internal/model"
	"github.com/jmoiron/sqlx"
)

type NewsInterface interface {
	AddNews(posts []model.Post) error
	GetNews(count int) ([]model.Post, error)
}

type Repository struct {
	NewsRepo NewsInterface
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		NewsRepo: NewNewsRepository(db),
	}
}
