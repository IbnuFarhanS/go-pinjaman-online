package entity

import "time"

type Loan_Product struct {
	ID          int64     `JSON:"id"`
	Name        string    `JSON:"name"`
	Description string    `JSON:"description"`
	Persyaratan string    `JSON:"persyaratan"`
	Created_At  time.Time `JSON:"created_at"`
}
