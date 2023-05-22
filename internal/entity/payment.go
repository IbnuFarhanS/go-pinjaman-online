package entity

import "time"

type Payment struct {
	ID             int64       `JSON:"id"`
	ID_Transaction Transaction `JSON:"id_transaction"`
	Payment_Amount string      `JSON:"payment_amount"`
	Payment_Date   time.Time   `JSON:"payment_date"`
	Payment_Method string      `JSON:"payment_method"`
}
