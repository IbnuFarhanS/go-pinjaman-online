package usecase

import (
	"errors"

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
	// Validate name
	if newLoanProduct.Name == "" {
		return nil, errors.New("name is required")
	}
	// Check if name already exists
	existingLoan_Product, err := u.loanProductRepo.FindByName(newLoanProduct.Name)
	if err != nil {
		return nil, err
	}
	if existingLoan_Product != nil {
		return nil, errors.New("name is already use")
	}

	// Validate description
	if newLoanProduct.Description == "" {
		return nil, errors.New("description is required")
	}
	// Validate persyaratan
	if newLoanProduct.Persyaratan == "" {
		return nil, errors.New("persyaratan is required")
	}

	return u.loanProductRepo.Insert(newLoanProduct)
}

func (u *loanProductUsecase) FindByID(id int64) (*entity.Loan_Product, error) {
	return u.loanProductRepo.FindByID(id)
}

func (u *loanProductUsecase) FindAll() ([]entity.Loan_Product, error) {
	return u.loanProductRepo.FindAll()
}

func (u *loanProductUsecase) Update(updateLoanProduct *entity.Loan_Product) (*entity.Loan_Product, error) {
	// Validate name
	if updateLoanProduct.Name == "" {
		return nil, errors.New("name is required")
	}
	// Check if name already exists
	existingLoan_Product, err := u.loanProductRepo.FindByName(updateLoanProduct.Name)
	if err != nil {
		return nil, err
	}
	if existingLoan_Product != nil {
		return nil, errors.New("name is already use")
	}
	// Validate description
	if updateLoanProduct.Description == "" {
		return nil, errors.New("description is required")
	}
	// Validate persyaratan
	if updateLoanProduct.Persyaratan == "" {
		return nil, errors.New("persyaratan is required")
	}

	return u.loanProductRepo.Update(updateLoanProduct)
}

func (u *loanProductUsecase) Delete(deletedLoanProduct *entity.Loan_Product) error {
	return u.loanProductRepo.Delete(deletedLoanProduct)
}
