package repository

import (
	"fmt"

	"github.com/FlyKarlik/news_skilfactory/internal/model"
	"github.com/jmoiron/sqlx"
)

const (
	newsTable = "GoNews"
)

type NewsRepository struct {
	db *sqlx.DB
}

func NewNewsRepository(db *sqlx.DB) *NewsRepository {
	return &NewsRepository{db: db}
}

func (d *NewsRepository) AddNews(posts []model.Post) error {
	length := len(posts)
	query := fmt.Sprintf("INSERT INTO %s (title,content,pub_date,link) VALUES ($1,$2,$3,$4)", newsTable)

	for i := 0; i < length; i++ {

		_, err := d.db.Exec(query, posts[i].Title, posts[i].Content, posts[i].PubTime, posts[i].Link)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *NewsRepository) GetNews(count int) ([]model.Post, error) {
	var posts []model.Post

	query := fmt.Sprintf("SELECT id,title,content,pub_date,link FROM %s ORDER BY pub_date DESC LIMIT $1", newsTable)
	if err := d.db.Select(&posts, query, count); err != nil {
		return nil, err
	}

	return posts, nil
}
