package entity

import "time"

type Session struct {
	ID         int64     `JSON:"id"`
	Username   string    `JSON:"username"`
	Token      string    `JSON:"token"`
	Expired_At time.Time `JSON:"expired_at"`
	Created_At time.Time `JSON:"created_at"`
}
