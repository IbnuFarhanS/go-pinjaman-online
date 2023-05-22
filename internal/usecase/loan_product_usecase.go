package usecase

import (
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/repository"
)

type LoanProductUsecase interface {
	Insert(newLoanProduct *entity.Loan_Product) (*entity.Loan_Product, error)
	FindByID(id int64) (*entity.Loan_Product, error)
	FindAll() ([]entity.Loan_Product, error)
	Update(updateLoanProduct *entity.Loan_Product) (*entity.Loan_Product, error)
	Delete(deletedLoanProduct *entity.Loan_Product) error
}

type loanProductUsecase struct {
	loanProductRepo repository.LoanProductRepository
}

func NewLoanProductUsecase(loanProductRepo repository.LoanProductRepository) LoanProductUsecase {
	return &loanProductUsecase{loanProductRepo}
}

func (u *loanProductUsecase) Insert(newLoanProduct *entity.Loan_Product) (*entity.Loan_Product, error) {
	return u.loanProductRepo.Insert(newLoanProduct)
}

func (u *loanProductUsecase) FindByID(id int64) (*entity.Loan_Product, error) {
	return u.loanProductRepo.FindByID(id)
}

func (u *loanProductUsecase) FindAll() ([]entity.Loan_Product, error) {
	return u.loanProductRepo.FindAll()
}

func (u *loanProductUsecase) Update(updateLoanProduct *entity.Loan_Product) (*entity.Loan_Product, error) {
	return u.loanProductRepo.Update(updateLoanProduct)
}

func (u *loanProductUsecase) Delete(deletedLoanProduct *entity.Loan_Product) error {
	return u.loanProductRepo.Delete(deletedLoanProduct)
}
