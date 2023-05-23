package repository

import (
	"database/sql"

	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
)

type PaymentRepository interface {
	Insert(newPayment *entity.Payment) (*entity.Payment, error)
	FindByID(id int64) (*entity.Payment, error)
	FindAll() ([]entity.Payment, error)
	Update(updatePayment *entity.Payment) (*entity.Payment, error)
	Delete(deletedPayment *entity.Payment) error
}

type paymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) PaymentRepository {
	return &paymentRepository{db}
}

// ======================= INSERT ==============================
func (r *paymentRepository) Insert(newPayment *entity.Payment) (*entity.Payment, error) {
	stmt, err := r.db.Prepare("INSERT INTO payment(id_transaction, payment_amount, payment_date, payment_method) VALUES ($1,$2,$3,$4) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(newPayment.ID_Transaction.ID, newPayment.Payment_Amount, newPayment.Payment_Date, newPayment.Payment_Method).Scan(&newPayment.ID)
	if err != nil {
		return nil, err
	}
	return newPayment, nil
}

// ======================= FIND BY ID ==============================
func (r *paymentRepository) FindByID(id int64) (*entity.Payment, error) {
	var payment entity.Payment

	stmt, err := r.db.Prepare("SELECT id, id_transaction, payment_amount, payment_date, payment_method FROM payment WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	stmt.QueryRow(id).Scan(&payment.ID, &payment.ID_Transaction.ID, &payment.Payment_Amount, &payment.Payment_Date, &payment.Payment_Method)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}

// ======================= FIND ALL ==============================
func (r *paymentRepository) FindAll() ([]entity.Payment, error) {
	var payments []entity.Payment
	rows, err := r.db.Query("SELECT id, id_transaction, payment_amount, payment_date, payment_method FROM payment")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var payment entity.Payment
		err := rows.Scan(&payment.ID, &payment.ID_Transaction.ID, &payment.Payment_Amount, &payment.Payment_Date, &payment.Payment_Method)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	return payments, nil
}

// ======================= UPDATE ==============================
func (r *paymentRepository) Update(updatePayment *entity.Payment) (*entity.Payment, error) {
	stmt, err := r.db.Prepare("UPDATE payment SET id_transaction = $1, payment_amount = $2, payment_date = $3, payment_method = $4 WHERE id = $5")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(updatePayment.ID_Transaction.ID, updatePayment.Payment_Amount, updatePayment.Payment_Date, updatePayment.Payment_Method, &updatePayment.ID)
	if err != nil {
		return nil, err
	}

	return updatePayment, err
}

// ======================= DELETE ==============================
func (r *paymentRepository) Delete(deletedPayment *entity.Payment) error {
	stmt, err := r.db.Prepare("DELETE FROM payment WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(deletedPayment.ID)
	if err != nil {
		return err
	}

	return nil
}
