package data

type DoSomethingRequest struct {
	Name string `json:"name"`
}

type DoSomethingResponse struct {
	Value string `json:"value"`
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
	Price  int    `json:"price"`
	Branch string `json:"branch"`
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
