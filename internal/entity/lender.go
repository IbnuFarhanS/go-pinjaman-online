package entity

import "time"

type Lender struct {
	ID         int64     `JSON:"id"`
	Name       string    `JSON:"name"`
	Created_At time.Time `JSON:"created_at"`
}
