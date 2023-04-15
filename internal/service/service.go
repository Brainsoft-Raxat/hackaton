package service

import (
	"context"
	"github.com/Brainsoft-Raxat/hacknu/internal/repository"
	"github.com/Brainsoft-Raxat/hacknu/pkg/data"
)

type SomeService interface {
	DoSomething(ctx context.Context, request data.DoSomethingRequest) (response data.DoSomethingResponse, err error)
}

type OrderService interface {
	DocumentReady(ctx context.Context, req data.DocumentReadyRequest) (resp data.DocumentReadyResponse, err error)
	CheckIIN(ctx context.Context, req data.CheckIINRequest) (resp data.CheckIINResponse, err error)
}

type Service struct {
	SomeService
	OrderService
}

func New(repos *repository.Repository) *Service {
	return &Service{
		SomeService:  NewSomeService(repos),
		OrderService: NewOrderService(repos),
	}
}
