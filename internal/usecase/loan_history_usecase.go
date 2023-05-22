package usecase

import (
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/repository"
)

type LoanHistoryUsecase interface {
	Insert(newLoanHistory *entity.Loan_History) (*entity.Loan_History, error)
	FindByID(id int64) (*entity.Loan_History, error)
	FindAll() ([]entity.Loan_History, error)
	Update(updateLoanHistory *entity.Loan_History) (*entity.Loan_History, error)
	Delete(deletedLoanHistory *entity.Loan_History) error
}

type loanHistoryUsecase struct {
	loanHistoryRepo repository.LoanHistoryRepository
}

func NewLoanHistoryUsecase(loanHistoryRepo repository.LoanHistoryRepository) LoanHistoryUsecase {
	return &loanHistoryUsecase{loanHistoryRepo}
}

func (u *loanHistoryUsecase) Insert(newLoanHistory *entity.Loan_History) (*entity.Loan_History, error) {
	return u.loanHistoryRepo.Insert(newLoanHistory)
}

func (u *loanHistoryUsecase) FindByID(id int64) (*entity.Loan_History, error) {
	return u.loanHistoryRepo.FindByID(id)
}

func (u *loanHistoryUsecase) FindAll() ([]entity.Loan_History, error) {
	return u.loanHistoryRepo.FindAll()
}

func (u *loanHistoryUsecase) Update(updateLoanHistory *entity.Loan_History) (*entity.Loan_History, error) {
	return u.loanHistoryRepo.Update(updateLoanHistory)
}

func (u *loanHistoryUsecase) Delete(deletedLoanHistory *entity.Loan_History) error {
	return u.loanHistoryRepo.Delete(deletedLoanHistory)
}
