package data

type DoSomethingRequest struct {
	Name string `json:"name"`
}

type DoSomethingResponse struct {
	Value string `json:"value"`
}

type DocumentReadyRequest struct {
	Id    string `json:"id"`
	IIN   string `json:"iin"`
	Phone string `json:"phone"`
}

type DocumentReadyResponse struct {
	Message string `json:"message"`
}

type CheckIINRequest struct {
	IIN string `json:"iin"`
}

type CheckIINResponse struct {
	IsExists bool `json:"is_exists"`
}

type CreateOrderRequest struct {
	Region         string `json:"region"`
	City           string `json:"city"`
	Street         string `json:"street"`
	House          string `json:"house"`
	Entrance       string `json:"entrance"`
	Floor          string `json:"floor"`
	Corpus         string `json:"corpus"`
	Rc             string `json:"rc"`
	AdditionalData string `json:"additional_data"`
	TrustedFaceIin string `json:"trusted_face_iin"`
}

type CreateOrderResponse struct {
	Price    float64 `json:"price"`
	Time     int     `json:"time"`
	Distance int     `json:"distance"`
}

type ConfirmOrder struct {
}

type DistanceResponse struct {
	DestinationAddresses []string `json:"destination_addresses"`
	OriginAddresses      []string `json:"origin_addresses"`
	Rows                 []struct {
		Elements []struct {
			Distance struct {
				Text  string `json:"text"`
				Value int    `json:"value"`
			} `json:"distance"`
			Duration struct {
				Text  string `json:"text"`
				Value int    `json:"value"`
			} `json:"duration"`
			Status string `json:"status"`
		} `json:"elements"`
	} `json:"rows"`
	Status string `json:"status"`
}

type GetCoordinatesRequest struct {
	Street string `json:"street"`
}

type GetCoordinatesResponse struct {
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
	Distance string  `json:"distance"`
	Time     string  `json:"time"`
}

type GeocodingResponse struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
}
