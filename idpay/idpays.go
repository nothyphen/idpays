package idpay

type IdPays interface {
	Payment(SANDBOX string, API_KEY string, order_id int, amount int, name string, phone string, mail string, desc string, callback string) (string, string, error)
	Verify(id string, order_id string) (interface{}, error)
	Inquiry(id string, order_id string) (interface{}, error)
}

type idpays struct {}

func NewIdPays() IdPays{
	return &idpays{}
}