package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/FlyKarlik/news_skilfactory/config"
	"github.com/FlyKarlik/news_skilfactory/internal/gateways/handlers"
	"github.com/FlyKarlik/news_skilfactory/internal/gateways/server"
	"github.com/FlyKarlik/news_skilfactory/internal/model"
	"github.com/FlyKarlik/news_skilfactory/internal/usecases/repository"
	"github.com/FlyKarlik/news_skilfactory/internal/usecases/rss"
	"github.com/FlyKarlik/news_skilfactory/internal/usecases/service"
	"github.com/FlyKarlik/news_skilfactory/pkg/postgres"
	_ "github.com/lib/pq"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) RunApp() error {
	ctx := context.Background()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Printf("Failed to init config: %s", err.Error())
		return err
	}

	db, err := postgres.ConnectPostgres(cfg.DatabaseURL)
	if err != nil {
		log.Printf("Failed connect to postgresql: %s", err.Error())
		return err
	}

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	handlers := handlers.NewHandlers(svc)
	rss := rss.NewRss(cfg)
	srv := server.NewServer(cfg)
	router := server.NewRouter(handlers)

	go func() {
		if err := srv.Run(router); err != nil {
			log.Fatalf("failed to start server: %s", err.Error())
		}
	}()

	posts := make(chan []model.Post)
	errs := make(chan error)

	for _, url := range cfg.Urls {
		go rss.RssPost.ParseURL(url, posts, errs)
	}

	go func() {
		for p := range posts {
			svc.NewsService.AddNews(p)
		}
	}()

	go func() {
		for errors := range errs {
			log.Println(errors.Error())
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shuttdown(ctx); err != nil {

		return err
	}

	if err := db.Close(); err != nil {
		log.Printf("Failed to clode database: %s", err.Error())
		return err
	}
	return nil
}
