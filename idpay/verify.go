package idpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nothyphen/idpays/serializers"
)

func (i * idpays) Verify(SANDBOX string, id string, order_id string) (*serializers.VerifyRequest, error) {
	verify_url := "https://api.idpay.ir/v1.1/payment/verify"
	
	data := map[string]string{
		"id" : id,
		"order_id" : order_id,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("POST", verify_url, bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", i.apikey)
	request.Header.Add("X-SANDBOX", SANDBOX)

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		msg, _ := ioutil.ReadAll(response.Body)
		buf := map[string]interface{}{}
		err = json.Unmarshal(msg, &buf)
		fmt.Println(buf["error_code"], buf["error_message"])
		return nil, err
	}

	bindRequest := &serializers.VerifyRequest{}
	derr := json.NewDecoder(response.Body).Decode(bindRequest)
	if derr != nil {
		panic(derr)
	}
	return bindRequest, nil
}