package idpays

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type IdPays interface {
	Payment(SANDBOX string, API_KEY string, order_id int, amount int, name string, phone string, mail string, desc string, callback string) (string, string, error)
	Verify(id string, order_id string) (interface{}, error)
	Inquiry(id string, order_id string) (interface{}, error)
}

type idpays struct {}

type PaymentRequest struct {
	Id string `json:"id"`
	Link string `json:"link"`
}

func NewIdPays() IdPays{
	return &idpays{}
}

func (i *idpays) Payment(SANDBOX string, API_KEY string, order_id int, amount int, name string, phone string, mail string, desc string, callback string) (string, string, error) {
	payment_url := "https://api.idpay.ir/v1.1/payment"

	data := map[string]interface{}{
		"order_id" : order_id,
		"amount" : amount,
		"name" : name,
		"phone" : phone,
		"mail" : mail,
		"desc" : desc,
		"callback" : callback,
	}

	payload, _ := json.Marshal(data)


	request, err := http.NewRequest("POST", payment_url, bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", API_KEY)
	request.Header.Add("X-SANDBOX", SANDBOX)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	bindRequest := &PaymentRequest{}
	derr := json.NewDecoder(response.Body).Decode(bindRequest)
	if derr != nil {
		panic(derr)
	}

	return bindRequest.Id, bindRequest.Link, nil
}

func (i * idpays) Verify(id string, order_id string) (interface{}, error) {
	return "ok", nil
}

func (i * idpays) Inquiry(id string, order_id string) (interface{}, error) {
	return "ok", nil
}