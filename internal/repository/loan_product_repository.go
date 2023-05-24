package repository

import (
	"errors"
	"time"

	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
	"gorm.io/gorm"
)

type LoanProductRepository interface {
	Insert(newLoanProduct *entity.Loan_Product) (*entity.Loan_Product, error)
	FindByID(id int64) (*entity.Loan_Product, error)
	FindByName(name string) (*entity.Loan_Product, error)
	FindAll() ([]entity.Loan_Product, error)
	Update(updateLoanProduct *entity.Loan_Product) (*entity.Loan_Product, error)
	Delete(deletedLoanProduct *entity.Loan_Product) error
}

type loanProductRepository struct {
	db *gorm.DB
}

func NewLoanProductRepository(db *gorm.DB) LoanProductRepository {
	return &loanProductRepository{db}
}

// ======================= INSERT ==============================
func (r *loanProductRepository) Insert(newLoanProduct *entity.Loan_Product) (*entity.Loan_Product, error) {
	currentTime := time.Now()
	newLoanProduct.Created_At = currentTime
	if err := r.db.Create(newLoanProduct).Error; err != nil {
		return nil, err
	}
	return newLoanProduct, nil
}

// ======================= FIND BY ID ==============================
func (r *loanProductRepository) FindByID(id int64) (*entity.Loan_Product, error) {
	var loan_product entity.Loan_Product

	if err := r.db.Where("id = ?", id).Find(&loan_product).Error; err != nil {
		return nil, err
	}
	return &loan_product, nil
}

// ======================= FIND BY NAME ==============================
func (r *loanProductRepository) FindByName(name string) (*entity.Loan_Product, error) {
	var loan_product entity.Loan_Product

	if err := r.db.Where("name = ?", name).First(&loan_product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &loan_product, nil
}

// ======================= FIND ALL ==============================
func (r *loanProductRepository) FindAll() ([]entity.Loan_Product, error) {
	var loan_products []entity.Loan_Product

	if err := r.db.Find(&loan_products).Error; err != nil {
		return nil, err
	}
	return loan_products, nil
}

// ======================= UPDATE ==============================
func (r *loanProductRepository) Update(updateLoanProduct *entity.Loan_Product) (*entity.Loan_Product, error) {
	var lp entity.Loan_Product
	if err := r.db.Where("id = ?", updateLoanProduct.ID).First(&lp).Error; err != nil {
		return nil, err
	}

	create_at := lp.Created_At

	lp.Name = updateLoanProduct.Name
	lp.Description = updateLoanProduct.Description
	lp.Persyaratan = updateLoanProduct.Persyaratan
	lp.Created_At = create_at

	if err := r.db.Save(&lp).Error; err != nil {
		return nil, err
	}
	return &lp, nil
}

// ======================= DELETE ==============================
func (r *loanProductRepository) Delete(deletedLender *entity.Loan_Product) error {
	if err := r.db.Delete(deletedLender).Error; err != nil {
		return err
	}
	return nil
}
