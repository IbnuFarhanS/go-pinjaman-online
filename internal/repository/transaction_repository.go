package repository

import (
	"database/sql"

	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
)

type TransactionRepository interface {
	Insert(newTransaction *entity.Transaction) (*entity.Transaction, error)
	FindByID(id int64) (*entity.Transaction, error)
	FindAll() ([]entity.Transaction, error)
	Update(updateTransaction *entity.Transaction) (*entity.Transaction, error)
	Delete(deletedTransaction *entity.Transaction) error
}

type transactionRepository struct {
	db *sql.DB
}

func newTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{db}
}

// ======================= INSERT ==============================
func (r *transactionRepository) Insert(newTransaction *entity.Transaction) (*entity.Transaction, error) {
	stmt, err := r.db.Prepare("INSERT INTO transaction(id_borrower, id_lender, id_loan_product, loan_amount, transaction_date, due_date, transaction_status) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(newTransaction.ID_Borrower.ID, newTransaction.ID_Lender.ID, newTransaction.ID_Loan_Product.ID, newTransaction.Loan_Amount, newTransaction.Transaction_Date, newTransaction.Due_Date, newTransaction.Transaction_Status).Scan(&newTransaction.ID)
	if err != nil {
		return nil, err
	}
	return newTransaction, nil
}

// ======================= FIND BY ID ==============================
func (r *transactionRepository) FindByID(id int64) (*entity.Transaction, error) {
	var transaction entity.Transaction

	stmt, err := r.db.Prepare("SELECT id, id_borrower, id_lender, id_loan_product, loan_amount, transaction_date, due_date, transaction_status FROM transaction WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	stmt.QueryRow(id).Scan(&transaction.ID, &transaction.ID_Borrower.ID, &transaction.ID_Lender.ID, &transaction.ID_Loan_Product.ID, &transaction.Loan_Amount, &transaction.Transaction_Date, &transaction.Due_Date, &transaction.Transaction_Status)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

// ======================= FIND ALL ==============================
func (r *transactionRepository) FindAll() ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	rows, err := r.db.Query("SELECT id, id_borrower, id_lender, id_loan_product, loan_amount, transaction_date, due_date, transaction_status FROM transaction")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction entity.Transaction
		err := rows.Scan(&transaction.ID, &transaction.ID_Borrower.ID, &transaction.ID_Lender.ID, &transaction.ID_Loan_Product.ID, &transaction.Loan_Amount, &transaction.Transaction_Date, &transaction.Due_Date, &transaction.Transaction_Status)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

// ======================= UPDATE ==============================
func (r *transactionRepository) Update(updateTransaction *entity.Transaction) (*entity.Transaction, error) {
	stmt, err := r.db.Prepare("UPDATE transaction SET id_borrower = $1, id_lender = $2, id_loan_product = $3, loan_amount = $4, transaction_date = $5, due_date = $6, transaction_status = $7 WHERE id = $8")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(updateTransaction.ID_Borrower.ID, updateTransaction.ID_Lender.ID, updateTransaction.ID_Loan_Product.ID, updateTransaction.Loan_Amount, updateTransaction.Transaction_Date, updateTransaction.Due_Date, updateTransaction.Transaction_Status, &updateTransaction.ID)
	if err != nil {
		return nil, err
	}

	return updateTransaction, err
}

// ======================= DELETE ==============================
func (r *transactionRepository) Delete(deletedTransaction *entity.Transaction) error {
	stmt, err := r.db.Prepare("DELETE FROM transaction WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(deletedTransaction.ID)
	if err != nil {
		return err
	}

	return nil
}
