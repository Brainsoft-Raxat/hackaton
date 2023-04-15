package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"hackaton/internal/app/config"
	"hackaton/internal/models"
)

type postgres struct {
	db  *pgxpool.Pool
	cfg config.Postgres
}

// всё взаимодействие с постгрес
func (r *postgres) GetOrders(ctx context.Context, id int) (Orders models.Orders, err error) {

	q := `SELECT * FROM Orders WHERE id = $1 `
	err = r.db.QueryRow(ctx, q, id).Scan(&Orders.Id, &Orders.Iin, &Orders.Request_id, &Orders.Service_name, &Orders.Organization_code, &Orders.Organization_name, &Orders.Recipient_name, &Orders.Recipient_surname, &Orders.Recipient_phone, &Orders.Recipient_phone, &Orders.Region, &Orders.City, &Orders.Street, &Orders.House, &Orders.Entrance, &Orders.Floor, &Orders.Corpus, &Orders.Rc, &Orders.Additional_data, &Orders.Trusted_face_iin, &Orders.Delivery_service_id, &Orders.Delivery_price, &Orders.Courier_id, &Orders.Status)
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

	err = r.db.QueryRow(ctx, q, Order.Id, Order.Iin, Order.Request_id, Order.Service_name, Order.Organization_code, Order.Organization_name, Order.Recipient_name, Order.Recipient_surname, Order.Recipient_phone, Order.Recipient_phone, Order.Region, Order.City, Order.Street, Order.House, Order.Entrance, Order.Floor, Order.Corpus, Order.Rc, Order.Additional_data, Order.Trusted_face_iin, Order.Delivery_service_id, Order.Delivery_price, Order.Courier_id, Order.Status).Scan(&id)

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

	err = r.db.QueryRow(ctx, q, Couriers.Id, Couriers.Delivery_service_id, Couriers.Firstname, Couriers.Surname, Couriers.Phone).Scan(&Couriers.Id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *postgres) GetCouriers(ctx context.Context, id int) (Courier models.Couriers, err error) {
	q := `SELECT * FROM Orders WHERE id = $1 `
	err = r.db.QueryRow(ctx, q, id).Scan(&Courier.Id, &Courier.Delivery_service_id, &Courier.Firstname, &Courier.Surname, &Courier.Phone)
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
