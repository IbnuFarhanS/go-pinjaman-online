package entity

import "time"

type Loan_History struct {
	ID             int64       `JSON:"id"`
	ID_Transaction Transaction `JSON:"id_transaction"`
	History_State  string      `JSON:"history_state"`
	Information    string      `JSON:"information"`
	Change_Date    time.Time   `JSON:"change_date"`
	Created_At     time.Time   `JSON:"created_at"`
}
