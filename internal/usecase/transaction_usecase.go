package usecase

import (
	"errors"

	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/repository"
	"github.com/shopspring/decimal"
)

type TransactionUsecase interface {
	Insert(newTransaction *entity.Transaction) (*entity.Transaction, error)
	FindByID(id int64) (*entity.Transaction, error)
	FindAll() ([]entity.Transaction, error)
	Update(updateTransaction *entity.Transaction) (*entity.Transaction, error)
	Delete(deletedTransaction *entity.Transaction) error
}

type transactionUsecase struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionUsecase(transactionRepo repository.TransactionRepository) TransactionUsecase {
	return &transactionUsecase{transactionRepo}
}

func (u *transactionUsecase) Insert(newTransaction *entity.Transaction) (*entity.Transaction, error) {
	// Validate ID Borrower
	// if newTransaction.ID_Borrower.ID == 0 {
	// 	return nil, errors.New("borrower id is required")
	// }

	// // Validate ID Lender
	// if newTransaction.ID_Lender.ID == 0 {
	// 	return nil, errors.New("lender id is required")
	// }

	// // Validate ID Loan Product
	// if newTransaction.ID_Loan_Product.ID == 0 {
	// 	return nil, errors.New("loan product id is required")
	// }

	// Validate loan_amount
	if newTransaction.Loan_Amount.LessThanOrEqual(decimal.Zero) {
		return nil, errors.New("the loan amount must be greater than 0")
	}
	// Validate due_date
	// if newTransaction.Due_Date.IsZero() || newTransaction.Due_Date.Before(newTransaction.Transaction_Date) {
	// 	return nil, errors.New("the due date must be filled in and cannot be before the loan date")
	// }
	// Validate alamat
	if newTransaction.Transaction_Status == "" {
		return nil, errors.New("transaction status is required")
	}

	return u.transactionRepo.Insert(newTransaction)
}

func (u *transactionUsecase) FindByID(id int64) (*entity.Transaction, error) {
	return u.transactionRepo.FindByID(id)
}

func (u *transactionUsecase) FindAll() ([]entity.Transaction, error) {
	return u.transactionRepo.FindAll()
}

func (u *transactionUsecase) Update(updateTransaction *entity.Transaction) (*entity.Transaction, error) {
	return u.transactionRepo.Update(updateTransaction)
}

func (u *transactionUsecase) Delete(deletedTransaction *entity.Transaction) error {
	return u.transactionRepo.Delete(deletedTransaction)
}
