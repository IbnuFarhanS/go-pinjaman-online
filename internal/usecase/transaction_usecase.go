package usecase

import (
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/repository"
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
