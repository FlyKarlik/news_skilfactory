package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/FlyKarlik/news_skilfactory/internal/model"
	"github.com/jmoiron/sqlx"
)

func TestNewsRepository_AddNews(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewRepository(sqlx.NewDb(db, "sqlmock"))

	posts := []model.Post{
		{
			Title:   "Test Title 1",
			Content: "Test Content 1",
			PubTime: time.Now(),
			Link:    "http://example.com/1",
		},
		{
			Title:   "Test Title 2",
			Content: "Test Content 2",
			PubTime: time.Now(),
			Link:    "http://example.com/2",
		},
	}

	mock.ExpectExec("INSERT INTO GoNews").WithArgs(posts[0].Title, posts[0].Content, posts[0].PubTime, posts[0].Link).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO GoNews").WithArgs(posts[1].Title, posts[1].Content, posts[1].PubTime, posts[1].Link).WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.NewsRepo.AddNews(posts)

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestNewsRepository_GetNews(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := &NewsRepository{sqlx.NewDb(db, "sqlmock")}

	expectedPosts := []model.Post{
		{
			ID:      "1",
			Title:   "Test Title 1",
			Content: "Test Content 1",
			PubTime: time.Now(),
			Link:    "http://example.com/1",
		},
		{
			ID:      "2",
			Title:   "Test Title 2",
			Content: "Test Content 2",
			PubTime: time.Now(),
			Link:    "http://example.com/2",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "content", "pub_date", "link"}).
		AddRow(expectedPosts[0].ID, expectedPosts[0].Title, expectedPosts[0].Content, expectedPosts[0].PubTime, expectedPosts[0].Link).
		AddRow(expectedPosts[1].ID, expectedPosts[1].Title, expectedPosts[1].Content, expectedPosts[1].PubTime, expectedPosts[1].Link)

	mock.ExpectQuery("SELECT id,title,content,pub_date,link FROM GoNews ORDER BY pub_date DESC LIMIT").WithArgs(2).WillReturnRows(rows)

	posts, err := repo.GetNews(2)

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if len(posts) != len(expectedPosts) {
		t.Errorf("expected %d posts, got %d", len(expectedPosts), len(posts))
	}

	for i := range posts {
		if posts[i].ID != expectedPosts[i].ID ||
			posts[i].Title != expectedPosts[i].Title ||
			posts[i].Content != expectedPosts[i].Content ||
			!posts[i].PubTime.Equal(expectedPosts[i].PubTime) ||
			posts[i].Link != expectedPosts[i].Link {
			t.Errorf("expected post %v, got %v", expectedPosts[i], posts[i])
		}
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
