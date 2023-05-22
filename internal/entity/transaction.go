package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID                 int64           `JSON:"id"`
	ID_Borrower        Borrower        `JSON:"id_borrower"`
	ID_Lender          Lender          `JSON:"id_lender"`
	ID_Loan_Product    Loan_Product    `JSON:"id_loan_product"`
	Loan_Amount        decimal.Decimal `JSON:"loan_amount"`
	Transaction_Date   time.Time       `JSON:"transaction_date"`
	Due_Date           time.Time       `JSON:"due_date"`
	Transaction_Status string          `JSON:"transaction_status"`
}
