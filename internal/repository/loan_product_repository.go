package repository

import (
	"database/sql"

	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
)

type LoanProductRepository interface {
	Insert(newLoanProduct *entity.Loan_Product) (*entity.Loan_Product, error)
	FindByID(id int64) (*entity.Loan_Product, error)
	FindAll() ([]entity.Loan_Product, error)
	Update(updateLoanProduct *entity.Loan_Product) (*entity.Loan_Product, error)
	Delete(deletedLoanProduct *entity.Loan_Product) error
}

type loanProductRepository struct {
	db *sql.DB
}

func NewLoanProductRepository(db *sql.DB) LoanProductRepository {
	return &loanProductRepository{db}
}

// ======================= INSERT ==============================
func (r *loanProductRepository) Insert(newLoanProduct *entity.Loan_Product) (*entity.Loan_Product, error) {
	stmt, err := r.db.Prepare("INSERT INTO loan_product(name, description, persyaratan, created_at) VALUES ($1,$2,$3,$4) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(newLoanProduct.Name, newLoanProduct.Description, newLoanProduct.Persyaratan, newLoanProduct.Created_At).Scan(&newLoanProduct.ID)
	if err != nil {
		return nil, err
	}
	return newLoanProduct, nil
}

// ======================= FIND BY ID ==============================
func (r *loanProductRepository) FindByID(id int64) (*entity.Loan_Product, error) {
	var loan_product entity.Loan_Product

	stmt, err := r.db.Prepare("SELECT id, name, description, persyaratan, created_at FROM loan_product WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	stmt.QueryRow(id).Scan(&loan_product.ID, &loan_product.Name, &loan_product.Description, &loan_product.Persyaratan, &loan_product.Created_At)
	if err != nil {
		return nil, err
	}

	return &loan_product, nil
}

// ======================= FIND ALL ==============================
func (r *loanProductRepository) FindAll() ([]entity.Loan_Product, error) {
	var loanProducts []entity.Loan_Product
	rows, err := r.db.Query("SELECT id, name, description, persyaratan, created_at FROM loan_product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var loan_product entity.Loan_Product
		err := rows.Scan(&loan_product.ID, &loan_product.Name, &loan_product.Description, &loan_product.Persyaratan, &loan_product.Created_At)
		if err != nil {
			return nil, err
		}
		loanProducts = append(loanProducts, loan_product)
	}

	return loanProducts, nil
}

// ======================= UPDATE ==============================
func (r *loanProductRepository) Update(updateLoanProduct *entity.Loan_Product) (*entity.Loan_Product, error) {
	stmt, err := r.db.Prepare("UPDATE loan_product SET name = $1, description = $2, persyaratan = $3, created_at = $4 WHERE id = $5")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(updateLoanProduct.Name, updateLoanProduct.Description, updateLoanProduct.Persyaratan, updateLoanProduct.Created_At, &updateLoanProduct.ID)
	if err != nil {
		return nil, err
	}

	return updateLoanProduct, err
}

// ======================= DELETE ==============================
func (r *loanProductRepository) Delete(deletedLoanProduct *entity.Loan_Product) error {
	stmt, err := r.db.Prepare("DELETE FROM loan_product WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(deletedLoanProduct.ID)
	if err != nil {
		return err
	}

	return nil
}
