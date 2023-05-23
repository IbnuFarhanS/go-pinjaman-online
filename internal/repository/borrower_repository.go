package repository

import (
	"database/sql"
	"time"

	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
)

type BorrowerRepository interface {
	Insert(newBorrower *entity.Borrower) (*entity.Borrower, error)
	FindByID(id int64) (*entity.Borrower, error)
	FindAll() ([]entity.Borrower, error)
	Update(updateBorrower *entity.Borrower) (*entity.Borrower, error)
	Delete(deletedBorrower *entity.Borrower) error
}

type borrowerRepository struct {
	db *sql.DB
}

func NewBorrowerRepository(db *sql.DB) BorrowerRepository {
	return &borrowerRepository{db}
}

// ======================= INSERT ==============================
func (r *borrowerRepository) Insert(newBorrower *entity.Borrower) (*entity.Borrower, error) {
	stmt, err := r.db.Prepare("INSERT INTO borrower(username, password, name, alamat, phone_number, created_at) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	currentTime := time.Now()
	newBorrower.Created_At = currentTime

	err = stmt.QueryRow(newBorrower.Username, newBorrower.Password, newBorrower.Name, newBorrower.Alamat, newBorrower.Phone_Number, newBorrower.Created_At).Scan(&newBorrower.ID)
	if err != nil {
		return nil, err
	}
	return newBorrower, nil
}

// ======================= FIND BY ID ==============================
func (r *borrowerRepository) FindByID(id int64) (*entity.Borrower, error) {
	var borrower entity.Borrower

	stmt, err := r.db.Prepare("SELECT id, username, password, name, alamat, phone_number FROM borrower WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	stmt.QueryRow(id).Scan(&borrower.ID, &borrower.Username, &borrower.Password, &borrower.Name, &borrower.Alamat, &borrower.Phone_Number)
	if err != nil {
		return nil, err
	}

	return &borrower, nil
}

// ======================= FIND ALL ==============================
func (r *borrowerRepository) FindAll() ([]entity.Borrower, error) {
	var borrowers []entity.Borrower
	rows, err := r.db.Query("SELECT id, username, password, name, alamat, phone_number, created_at FROM borrower")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var borrower entity.Borrower
		err := rows.Scan(&borrower.ID, &borrower.Username, &borrower.Password, &borrower.Name, &borrower.Alamat, &borrower.Phone_Number, &borrower.Created_At)
		if err != nil {
			return nil, err
		}
		borrowers = append(borrowers, borrower)
	}

	return borrowers, nil
}

// ======================= UPDATE ==============================
func (r *borrowerRepository) Update(updateBorrower *entity.Borrower) (*entity.Borrower, error) {
	stmt, err := r.db.Prepare("UPDATE borrower SET username = $1, password = $2, name = $3, alamat = $4, phone_number = $5 WHERE id = $6")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(updateBorrower.Username, updateBorrower.Password, updateBorrower.Name, updateBorrower.Alamat, updateBorrower.Phone_Number, &updateBorrower.ID)
	if err != nil {
		return nil, err
	}

	return updateBorrower, err
}

// ======================= DELETE ==============================
func (r *borrowerRepository) Delete(deletedBorrower *entity.Borrower) error {
	stmt, err := r.db.Prepare("DELETE FROM borrower WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(deletedBorrower.ID)
	if err != nil {
		return err
	}

	return nil
}
