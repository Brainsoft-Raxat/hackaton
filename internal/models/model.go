package models

type Orders struct {
	Id                  int    `json:"id" db:"id"`
	Iin                 string `json:"iin" db:"iin"`
	Request_id          string `json:"request_id" db:"request_id"`
	Service_name        string `json:"service_name" db:"service_name"`
	Organization_code   string `json:"organization_code" db:"organization_code"`
	Organization_name   string `json:"organization_name" db:"organization_name"`
	Recipient_name      string `json:"recipient_name" db:"recipient_name"`
	Recipient_surname   string `json:"recipient_surname" db:"recipient_surname"`
	Recipient_phone     string `json:"recipient_phone" db:"recipient_phone"`
	Region              string `json:"region" db:"region"`
	City                string `json:"city" db:"city"`
	Street              string `json:"street" db:"street"`
	House               string `json:"house" db:"house"`
	Entrance            string `json:"entrance" db:"entrance"`
	Floor               string `json:"floor" db:"floor"`
	Corpus              string `json:"corpus" db:"corpus"`
	Rc                  string `json:"rc" db:"rc"`
	Additional_data     string `json:"additional_data" db:"additional_data"`
	Trusted_face_iin    string `json:"trusted_face_iin" db:"trusted_face_iin"`
	Delivery_service_id int    `json:"delivery_service_id" db:"delivery_service_id"`
	Delivery_price      int    `json:"delivery_price" db:"delivery_price"`
	Courier_id          int    `json:"courier_id" db:"courier_id"`
	Status              string `json:"status" db:"status"`
}

type DeliveryServices struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Couriers struct {
	Id                  int    `json:"id" db:"id"`
	Delivery_service_id int    `json:"delivery_service_id" db:"delivery_service_id"`
	Firstname           string `json:"firstname" db:"firstname"`
	Surname             string `json:"surname" db:"surname"`
	Phone               string `json:"phone" db:"phone"`
}
