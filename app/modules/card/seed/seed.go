package cardSeed

import (
	"bytes"
	"encoding/json"
	"io"

	"net/http"

	"github.com/Daizaikun/drive-back/app/lib"
	cardModel "github.com/Daizaikun/drive-back/app/modules/card/model"
)

type responseData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	User uint   `json:"user"`
}

type response struct {
	Status string       `json:"status"`
	Data   responseData `json:"data"`
}

func sendPostRequest(jsonStr []byte) responseData {

	url := "https://sandbox.wompi.co/v1/tokens/cards"

	bearer := lib.SetWompiKey().Public

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic("error creating seed card")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic("error creating seed card")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("error creating seed card")
	}

	var responseObj response
	err = json.Unmarshal(body, &responseObj)
	if err != nil {
		panic("error creating seed card")
	}

	return responseObj.Data
}

func GenerateSeed(ids []uint) []*cardModel.Model {
	jsonStr1 := []byte(`{
		"number": "4111111111111111",
		"exp_month": "06",
		"exp_year": "29",
		"cvc": "123",
		"card_holder": "Pedro Pérez"
	}`)

	jsonStr2 := []byte(`{
		"number": "4242424242424242",
		"exp_month": "06",
		"exp_year": "29",
		"cvc": "123",
		"card_holder": "Pedro Pérez"
	}`)

	seed1 := sendPostRequest(jsonStr1)
	seed2 := sendPostRequest(jsonStr2)

	Seed := []*cardModel.Model{
		{
			Token:   seed1.ID,
			Name:    seed1.Name,
			RiderID: ids[0],
		},
		{
			Token:   seed2.ID,
			Name:    seed2.Name,
			RiderID: ids[1],
		},
	}
	return Seed

}
