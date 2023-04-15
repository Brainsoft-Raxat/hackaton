package models

import "time"

type Auth struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

type Person struct {
	Iin          string    `json:"iin"`
	LastName     string    `json:"lastName"`
	FirstName    string    `json:"firstName"`
	MiddleName   string    `json:"middleName"`
	EngFirstName string    `json:"engFirstName"`
	EngSurname   string    `json:"engSurname"`
	DateOfBirth  time.Time `json:"dateOfBirth"`
	Nationality  struct {
		Code   string `json:"code"`
		NameRu string `json:"nameRu"`
		NameKz string `json:"nameKz"`
	} `json:"nationality"`
	Gender struct {
		Code   string `json:"code"`
		NameRu string `json:"nameRu"`
		NameKz string `json:"nameKz"`
	} `json:"gender"`
	RegAddress struct {
		Address string `json:"address"`
		Country struct {
			Code   string `json:"code"`
			NameRu string `json:"nameRu"`
			NameKz string `json:"nameKz"`
		} `json:"country"`
		District struct {
			Code   string `json:"code"`
			NameRu string `json:"nameRu"`
			NameKz string `json:"nameKz"`
		} `json:"district"`
		Region struct {
			Code   string `json:"code"`
			NameRu string `json:"nameRu"`
			NameKz string `json:"nameKz"`
		} `json:"region"`
		StreetLocation    string    `json:"streetLocation"`
		HouseLocation     string    `json:"houseLocation"`
		ApartmentLocation string    `json:"apartmentLocation"`
		BeginDate         time.Time `json:"beginDate"`
		Status            struct {
			Code   string `json:"code"`
			NameRu string `json:"nameRu"`
			NameKz string `json:"nameKz"`
		} `json:"status"`
		Invalidity struct {
			Code   string `json:"code"`
			NameRu string `json:"nameRu"`
			NameKz string `json:"nameKz"`
		} `json:"invalidity"`
		ArCode string `json:"arCode"`
	} `json:"regAddress"`
	BirthPlace struct {
		Country struct {
			Code   string `json:"code"`
			NameRu string `json:"nameRu"`
			NameKz string `json:"nameKz"`
		} `json:"country"`
		District struct {
			Code   string `json:"code"`
			NameRu string `json:"nameRu"`
			NameKz string `json:"nameKz"`
		} `json:"district"`
		Region struct {
			Code   string `json:"code"`
			NameRu string `json:"nameRu"`
			NameKz string `json:"nameKz"`
		} `json:"region"`
		City          string      `json:"city"`
		BirthTeCodeAR interface{} `json:"birthTeCodeAR"`
	} `json:"birthPlace"`
	Documents []struct {
		DocTypeCode                string    `json:"docTypeCode"`
		DocTypeNameRu              string    `json:"docTypeNameRu"`
		DocTypeNameKz              string    `json:"docTypeNameKz"`
		DocStatusCode              string    `json:"docStatusCode"`
		DocStatusNameRu            string    `json:"docStatusNameRu"`
		DocStatusNameKz            string    `json:"docStatusNameKz"`
		DocIssueOrganizationCode   string    `json:"docIssueOrganizationCode"`
		DocIssueOrganizationNameRu string    `json:"docIssueOrganizationNameRu"`
		DocIssueOrganizationNameKz string    `json:"docIssueOrganizationNameKz"`
		DocNumber                  string    `json:"docNumber"`
		BeginDate                  time.Time `json:"beginDate"`
		EndDate                    time.Time `json:"endDate"`
	} `json:"documents"`
}

type GetPersonDataResponse struct {
	Person Person `json:"person"`
}

type SendSMSRequest struct {
	Phone   string `json:"phone"`
	SmsText string `json:"smsText"`
}

type SendSMSResponse struct {
	Status        string `json:"status"`
	StatusMessage string `json:"statusMessage"`
}

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
