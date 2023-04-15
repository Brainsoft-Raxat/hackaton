package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
)

type DistanceResponse struct {
	Rows []struct {
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

func GetDistance(originCity string, originAddress string, originHouse string, destinationAddress string, destinationHouse string) (float64, error) {

	originAddress = strings.ReplaceAll(originAddress, " ", "_")
	destinationAddress = strings.ReplaceAll(destinationAddress, " ", "_")

	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/distancematrix/json?units=metric&origins=%s%s%%20%s&destinations=%s%s%%20%s&key=AIzaSyCUf6GIt3soIsxHxfGmg7jBoh8yN2A57z8", originCity, originAddress, originHouse, originCity, destinationAddress, destinationHouse)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	var distanceResponse DistanceResponse
	err = json.Unmarshal([]byte(body), &distanceResponse)
	if err != nil {
		panic(err)
	}

	// get the distance value
	distanceValue := distanceResponse.Rows[0].Elements[0].Distance.Value
	price := math.Round(float64(distanceValue+50)/100) * 10
	return price, err
}
