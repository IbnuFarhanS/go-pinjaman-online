package repository

import (
	"database/sql"

	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
)

type LenderRepository interface {
	Insert(newLender *entity.Lender) (*entity.Lender, error)
	FindByID(id int64) (*entity.Lender, error)
	FindAll() ([]entity.Lender, error)
	Update(updateLender *entity.Lender) (*entity.Lender, error)
	Delete(deletedLender *entity.Lender) error
}

type lenderRepository struct {
	db *sql.DB
}

func NewLenderRepository(db *sql.DB) LenderRepository {
	return &lenderRepository{db}
}

// ======================= INSERT ==============================
func (r *lenderRepository) Insert(newLender *entity.Lender) (*entity.Lender, error) {
	stmt, err := r.db.Prepare("INSERT INTO lender(name, created_at) VALUES ($1,$2) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(newLender.Name, newLender.Created_At).Scan(&newLender.ID)
	if err != nil {
		return nil, err
	}
	return newLender, nil
}

// ======================= FIND BY ID ==============================
func (r *lenderRepository) FindByID(id int64) (*entity.Lender, error) {
	var lender entity.Lender

	stmt, err := r.db.Prepare("SELECT id, name, created_at FROM lender WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	stmt.QueryRow(id).Scan(&lender.ID, &lender.Name, &lender.Created_At)
	if err != nil {
		return nil, err
	}

	return &lender, nil
}

// ======================= FIND ALL ==============================
func (r *lenderRepository) FindAll() ([]entity.Lender, error) {
	var lenders []entity.Lender
	rows, err := r.db.Query("SELECT id, name, created_at FROM lender")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var lender entity.Lender
		err := rows.Scan(&lender.ID, &lender.Name, &lender.Created_At)
		if err != nil {
			return nil, err
		}
		lenders = append(lenders, lender)
	}

	return lenders, nil
}

// ======================= UPDATE ==============================
func (r *lenderRepository) Update(updateLender *entity.Lender) (*entity.Lender, error) {
	stmt, err := r.db.Prepare("UPDATE lender SET name = $1, created_at = $2 WHERE id = $3")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(updateLender.Name, updateLender.Created_At, &updateLender.ID)
	if err != nil {
		return nil, err
	}

	return updateLender, err
}

// ======================= DELETE ==============================
func (r *lenderRepository) Delete(deletedLender *entity.Lender) error {
	stmt, err := r.db.Prepare("DELETE FROM lender WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(deletedLender.ID)
	if err != nil {
		return err
	}

	return nil
}
