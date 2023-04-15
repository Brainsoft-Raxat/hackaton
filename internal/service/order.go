package service

import (
	"context"
	"fmt"
	"github.com/Brainsoft-Raxat/hacknu/internal/models"
	"github.com/Brainsoft-Raxat/hacknu/internal/repository"
	"github.com/Brainsoft-Raxat/hacknu/pkg/data"
	"time"
)

type orderService struct {
	orderRepo *repository.Repository
}

func (s *orderService) DocumentReady(ctx context.Context, req data.DocumentReadyRequest) (resp data.DocumentReadyResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	msg := models.SendSMSRequest{
		Phone:   req.Phone,
		SmsText: fmt.Sprintf(models.SMSTEMPLATE, req.Id, models.URLTEMPLATE+req.Id, req.Id, models.URLTEMPLATE+req.Id),
	}

	fmt.Println(msg.SmsText)
	//err = s.orderRepo.Egov.SendSMS(ctx, msg)
	//if err != nil {
	//	return
	//}

	return
}

func (s *orderService) CheckIIN(ctx context.Context, req data.CheckIINRequest) (resp data.CheckIINResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	eResp, err := s.orderRepo.CheckIIN(ctx, req.IIN)
	if err != nil {
		return
	}

	return data.CheckIINResponse{IsExists: eResp.IsExists}, nil
}

func NewOrderService(repo *repository.Repository) OrderService {
	return &orderService{
		orderRepo: repo,
	}
}
