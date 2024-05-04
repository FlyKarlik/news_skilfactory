package server

import (
	"github.com/FlyKarlik/news_skilfactory/internal/gateways/handlers"
	"github.com/FlyKarlik/news_skilfactory/internal/gateways/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *handlers.Handlers) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	news := router.Group("GoNews")
	{
		news.StaticFile("/", "./webapp")
		news.Use(middleware.SetHeader())
		news.GET("/news/:count", handler.GetNews)
	}

	return router
}
