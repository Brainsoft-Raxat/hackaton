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
	deliveries := map[string]int{
		"DHL":                 1,
		"Pony Express":        2,
		"Exline":              3,
		"CDEK":                4,
		"Garant Post Service": 5,
		"Алем-Тат":            6,
	}

	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	slc := strings.Split(req.Address, ",")
	home := strings.Split(strings.TrimSpace(slc[2]), " ")

	response, err := s.orderRepo.GetDistance(ctx, home[0], home[1])
	if err != nil {
		return
	}

	// get the distance value
	time := response.Rows[0].Elements[0].Duration.Text

	timeIn := strings.TrimSuffix(time, " mins") // Remove " mins" from the end of the string 	// Extract the first character of the string
	timeInMinute, err := strconv.Atoi(timeIn)
	distanceValue := response.Rows[0].Elements[0].Distance.Value
	price := math.Round(float64(distanceValue+50)/100) * 100

	requestDataResp, err := s.orderRepo.Egov.GetRequestData(ctx, models.GetRequestDataRequest{
		RequestID: req.RequestID,
		IIN:       req.IIN,
	})
	if err != nil {
		return
	}

	_, err = s.orderRepo.Postgres.SaveOrder(ctx, models.Orders{
		Iin:               req.IIN,
		RequestId:         req.RequestID,
		ServiceName:       requestDataResp.Data.ServiceType.NameRu,
		OrganizationCode:  requestDataResp.Data.Organization.Code,
		OrganizationName:  requestDataResp.Data.Organization.NameRu,
		RecipientName:     req.FirstName,
		RecipientSurname:  req.LastName,
		RecipientPhone:    req.Phone,
		Region:            strings.TrimSpace(slc[0]),
		City:              strings.TrimSpace(slc[1]),
		Street:            home[0],
		House:             home[2],
		Entrance:          "",
		Floor:             "",
		Corpus:            "",
		Rc:                "",
		AdditionalData:    req.AdditionalData,
		TrustedFaceIin:    req.TrustedFaceIin,
		DeliveryServiceId: deliveries[req.DeliveryService],
		DeliveryPrice:     int(price),
		CourierId:         0,
		Status:            "CREATED",
	})
	if err != nil {
		return
	}

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
