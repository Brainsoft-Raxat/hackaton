package repository

import (
	"context"
	"github.com/Brainsoft-Raxat/hacknu/internal/app/config"
	"github.com/Brainsoft-Raxat/hacknu/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type postgres struct {
	db  *pgxpool.Pool
	cfg config.Postgres
}

// всё взаимодействие с постгрес
func (r *postgres) GetOrders(ctx context.Context, id int) (Orders models.Orders, err error) {

	q := `SELECT * FROM Orders WHERE id = $1 `
	err = r.db.QueryRow(ctx, q, id).Scan(&Orders.Id, &Orders.Iin, &Orders.RequestId, &Orders.ServiceName, &Orders.OrganizationCode, &Orders.OrganizationName, &Orders.RecipientName, &Orders.RecipientSurname, &Orders.RecipientPhone, &Orders.RecipientPhone, &Orders.Region, &Orders.City, &Orders.Street, &Orders.House, &Orders.Entrance, &Orders.Floor, &Orders.Corpus, &Orders.Rc, &Orders.AdditionalData, &Orders.TrustedFaceIin, &Orders.DeliveryServiceId, &Orders.DeliveryPrice, &Orders.CourierId, &Orders.Status)
	if err != nil {
		return Orders, err
	}
	return Orders, nil
}

func (r *postgres) SaveOrder(ctx context.Context, Order models.Orders) (value int, err error) {
	var id int
	q := `insert into orders (id, iin, request_id, service_name, organization_code, organization_name, recipient_name,
                    recipient_surname, recipient_phone, region, city, street, house, entrance, floor, corpus, rc,
                    additional_data, trusted_face_iin, delivery_service_id, delivery_price, courier_id, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23) RETURNING id`

	err = r.db.QueryRow(ctx, q, Order.Id, Order.Iin, Order.RequestId, Order.ServiceName, Order.OrganizationCode, Order.OrganizationName, Order.RecipientName, Order.RecipientSurname, Order.RecipientPhone, Order.RecipientPhone, Order.Region, Order.City, Order.Street, Order.House, Order.Entrance, Order.Floor, Order.Corpus, Order.Rc, Order.AdditionalData, Order.TrustedFaceIin, Order.DeliveryServiceId, Order.DeliveryPrice, Order.CourierId, Order.Status).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *postgres) SaveDeliveryServices(ctx context.Context, DeliveryServices models.DeliveryServices) (value int, err error) {
	var id int
	q := `insert into delivery_services (id, name) values ($1, $2) RETURNING id`

	err = r.db.QueryRow(ctx, q, DeliveryServices.Id, DeliveryServices.Name).Scan(&DeliveryServices.Id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *postgres) GetDeliveryServices(ctx context.Context, id int) (DeliveryServices models.DeliveryServices, err error) {

	q := `SELECT * FROM Orders WHERE id = $1 `
	err = r.db.QueryRow(ctx, q, id).Scan(&DeliveryServices.Id, &DeliveryServices.Name)
	if err != nil {
		return DeliveryServices, err
	}
	return DeliveryServices, nil
}

func (r *postgres) SaveCouriers(ctx context.Context, Couriers models.Couriers) (value int, err error) {

	var id int
	q := `insert into couriers (id, delivery_service_id, firstname, surname, phone) values ($1, $2) RETURNING id`

	err = r.db.QueryRow(ctx, q, Couriers.Id, Couriers.DeliveryServiceId, Couriers.Firstname, Couriers.Surname, Couriers.Phone).Scan(&Couriers.Id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *postgres) GetCouriers(ctx context.Context, id int) (Courier models.Couriers, err error) {
	q := `SELECT * FROM Orders WHERE id = $1 `
	err = r.db.QueryRow(ctx, q, id).Scan(&Courier.Id, &Courier.DeliveryServiceId, &Courier.Firstname, &Courier.Surname, &Courier.Phone)
	if err != nil {
		return Courier, err
	}
	return Courier, nil
}

func NewPostgresRepository(db *pgxpool.Pool, cfg config.Postgres) Postgres {
	return &postgres{
		db:  db,
		cfg: cfg,
	}
}
