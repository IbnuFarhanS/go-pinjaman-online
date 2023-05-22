package entity

import "time"

type Borrower struct {
	ID           int64     `JSON:"id"`
	Username     string    `JSON:"username"`
	Password     string    `JSON:"password"`
	Name         string    `JSON:"name"`
	Alamat       string    `JSON:"alamat"`
	Phone_Number string    `JSON:"phone_number"`
	Created_At   time.Time `JSON:"created_at"`
}
