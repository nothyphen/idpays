package serializers

import (
	"time"
)


type VerifyRequest struct {
	Status		int			`json:"status"`
	Track_id	int			`json:"track_id"`
	Id			string		`json:"id"`
	Order_id	string		`json:"order_id"`
	Amount		int		`json:"amount"`
	Data		*time.Time	`json:"date"`
	Payment		struct {
		Track_id		string		`json:"track_id"`
		Amount			int			`json:"amount"`
		Cardno			int			`json:"card_no"`
		Hashed_card_no	string		`json:"hashed_card_no"`
		Date			*time.Time	`json:"date"`
	}
	Verify		struct {
		Date			*time.Time	`json:"date"`
	}
}