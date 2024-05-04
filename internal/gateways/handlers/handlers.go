package handlers

import "github.com/FlyKarlik/news_skilfactory/internal/usecases/service"

type Handlers struct {
	svc *service.Service
}

func NewHandlers(svc *service.Service) *Handlers {
	return &Handlers{svc: svc}
}
