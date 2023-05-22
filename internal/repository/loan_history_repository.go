package repository

import (
	"database/sql"

	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
)

type LoanHistoryRepository interface {
	Insert(newLoanHistory *entity.Loan_History) (*entity.Loan_History, error)
	FindByID(id int64) (*entity.Loan_History, error)
	FindAll() ([]entity.Loan_History, error)
	Update(updateHistory *entity.Loan_History) (*entity.Loan_History, error)
	Delete(deletedLoanHistory *entity.Loan_History) error
}

type loanHistoryRepository struct {
	db *sql.DB
}

func newLoanHistoryRepository(db *sql.DB) LoanHistoryRepository {
	return &loanHistoryRepository{db}
}

// ======================= INSERT ==============================
func (r *loanHistoryRepository) Insert(newLoanHistory *entity.Loan_History) (*entity.Loan_History, error) {
	stmt, err := r.db.Prepare("INSERT INTO loan_history(id_transaction, history_state, information, change_date, created_at) VALUES ($1,$2,$3,$4,$5) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(newLoanHistory.ID_Transaction.ID, newLoanHistory.History_State, newLoanHistory.Information, newLoanHistory.Change_Date, newLoanHistory.Created_At).Scan(&newLoanHistory.ID)
	if err != nil {
		return nil, err
	}
	return newLoanHistory, nil
}

// ======================= FIND BY ID ==============================
func (r *loanHistoryRepository) FindByID(id int64) (*entity.Loan_History, error) {
	var loanHistory entity.Loan_History

	stmt, err := r.db.Prepare("SELECT id, id_transaction, history_state, information, change_date, created_at FROM loan_history WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	stmt.QueryRow(id).Scan(&loanHistory.ID, &loanHistory.ID_Transaction.ID, &loanHistory.History_State, &loanHistory.Information, &loanHistory.Change_Date, &loanHistory.Created_At)
	if err != nil {
		return nil, err
	}

	return &loanHistory, nil
}

// ======================= FIND ALL ==============================
func (r *loanHistoryRepository) FindAll() ([]entity.Loan_History, error) {
	var loanHistorys []entity.Loan_History
	rows, err := r.db.Query("SELECT id, id_transaction, history_state, information, change_date, created_at FROM loan_history")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var loanHistory entity.Loan_History
		err := rows.Scan(&loanHistory.ID, &loanHistory.ID_Transaction.ID, &loanHistory.History_State, &loanHistory.Information, &loanHistory.Change_Date, &loanHistory.Created_At)
		if err != nil {
			return nil, err
		}
		loanHistorys = append(loanHistorys, loanHistory)
	}

	return loanHistorys, nil
}

// ======================= UPDATE ==============================
func (r *loanHistoryRepository) Update(updateHistory *entity.Loan_History) (*entity.Loan_History, error) {
	stmt, err := r.db.Prepare("UPDATE loan_history SET id_transaction = $1, history_state = $2, information = $3, change_date = $4, created_at = $5 WHERE id = $6")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(updateHistory.ID_Transaction.ID, updateHistory.History_State, updateHistory.Information, updateHistory.Change_Date, updateHistory.Created_At, &updateHistory.ID)
	if err != nil {
		return nil, err
	}

	return updateHistory, err
}

// ======================= DELETE ==============================
func (r *loanHistoryRepository) Delete(deletedLoanHistory *entity.Loan_History) error {
	stmt, err := r.db.Prepare("DELETE FROM loan_history WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(deletedLoanHistory.ID)
	if err != nil {
		return err
	}

	return nil
}
