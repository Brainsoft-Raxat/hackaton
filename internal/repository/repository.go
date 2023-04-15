package repository

import (
	"context"
	"github.com/Brainsoft-Raxat/hacknu/internal/app/config"
	"github.com/Brainsoft-Raxat/hacknu/internal/app/conn"
	"github.com/Brainsoft-Raxat/hacknu/internal/models"
)

type Repository struct {
	Postgres
	Egov
}

type Postgres interface {
	GetOrders(ctx context.Context, id int) (Orders models.Orders, err error)
	GetDeliveryServices(ctx context.Context, id int) (DeliveryServices models.DeliveryServices, err error)
	GetCouriers(ctx context.Context, id int) (Courier models.Couriers, err error)
	SaveCouriers(ctx context.Context, Couriers models.Couriers) (value int, err error)
	SaveOrder(ctx context.Context, Order models.Orders) (value int, err error)
	SaveDeliveryServices(ctx context.Context, DeliveryServices models.DeliveryServices) (value int, err error)
}

type Egov interface {
	GetPersonData(ctx context.Context, iin string) (person models.Person, err error)
	SendSMS(ctx context.Context, msg models.SendSMSRequest) (err error)
	GetRequestData(ctx context.Context, request models.GetRequestDataRequest) (response models.GetRequestDataResponse, err error)
	CheckIIN(ctx context.Context, iin string) (response models.CheckIINResponse, err error)
}

func New(conn conn.Conn, cfg *config.Config) *Repository {
	return &Repository{
		Postgres: NewPostgresRepository(conn.DB, cfg.Postgres),
		Egov:     NewEgov(cfg),
	}
}
