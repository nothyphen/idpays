package idpay

import "github.com/nothyphen/idpays/serializers"

type idPays interface {
	Payment(SANDBOX string, order_id int, amount int, name string, phone string, mail string, desc string, callback string) (string, string, error)
	Verify(SANDBOX string, id string, order_id string) (*serializers.VerifyRequest, error)
	Inquiry(id string, order_id string) (interface{}, error)
}

type idpays struct {
	apikey string
}

func NewIdPays(API_KEY string) *idpays{
	return &idpays{apikey: API_KEY}
}