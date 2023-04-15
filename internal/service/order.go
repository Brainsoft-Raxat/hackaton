package service

import (
	"context"
	"fmt"
	"github.com/Brainsoft-Raxat/hacknu/internal/models"
	"github.com/Brainsoft-Raxat/hacknu/internal/repository"
	"github.com/Brainsoft-Raxat/hacknu/pkg/data"
	"math"
	"strconv"
	"strings"
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

func (s *orderService) CreateOrder(ctx context.Context, req data.CreateOrderRequest) (resp data.CreateOrderResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	response, err := s.orderRepo.GetDistance(ctx, req.Street, req.House)
	if err != nil {
		return
	}

	// get the distance value
	time := response.Rows[0].Elements[0].Duration.Text

	timeIn := strings.TrimSuffix(time, " mins") // Remove " mins" from the end of the string 	// Extract the first character of the string
	timeInMinute, err := strconv.Atoi(timeIn)
	distanceValue := response.Rows[0].Elements[0].Distance.Value
	price := math.Round(float64(distanceValue+50)/100) * 100

	return data.CreateOrderResponse{Price: price, Distance: distanceValue, Time: timeInMinute}, nil

}

func (s *orderService) GetCoordinates(ctx context.Context, req data.GetCoordinatesRequest) (resp data.GetCoordinatesResponse, err error) {

	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	response, err := s.orderRepo.GetCoordinates(ctx, req.Street)

	if err != nil {
		return
	}

	Lat := response.Results[0].Geometry.Location.Lat
	Lng := response.Results[0].Geometry.Location.Lng

	return data.GetCoordinatesResponse{Lng: Lng, Lat: Lat}, nil

}

func (s *orderService) GetClientData(ctx context.Context, req data.GetClientDataRequest) (resp data.GetClientDataResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	iinResp, err := s.orderRepo.CheckIIN(ctx, req.IIN)
	if err != nil {
		return
	}

	clientResp, err := s.orderRepo.GetPersonData(ctx, req.IIN)
	if err != nil {
		return
	}

	return data.GetClientDataResponse{
		FirstName:  clientResp.FirstName,
		MiddleName: clientResp.MiddleName,
		LastName:   clientResp.LastName,
		Phone:      iinResp.Phone,
	}, nil
}

func (s *orderService) GetDeliveryServices(ctx context.Context) (deliveryServices []models.DeliveryServices, err error) {
	return s.orderRepo.GetDeliveryServices(ctx)
}

func NewOrderService(repo *repository.Repository) OrderService {
	return &orderService{
		orderRepo: repo,
	}
}
