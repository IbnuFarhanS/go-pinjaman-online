package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID                 int64           `json:"id"`
	ID_Borrower        Borrower        `json:"borrower"`
	ID_Lender          Lender          `json:"lender"`
	ID_Loan_Product    Loan_Product    `json:"loan_product"`
	Loan_Amount        decimal.Decimal `json:"loan_amount"`
	Transaction_Date   time.Time       `json:"transaction_date"`
	Due_Date           string       `json:"due_date"`
	Transaction_Status string          `json:"transaction_status"`
}
