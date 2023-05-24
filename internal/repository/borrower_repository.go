package repository

import (
	"time"

	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
	"gorm.io/gorm"
)

type BorrowerRepository interface {
	Insert(newBorrower *entity.Borrower) (*entity.Borrower, error)
	FindByID(id int64) (*entity.Borrower, error)
	FindAll() ([]entity.Borrower, error)
	Update(updateBorrower *entity.Borrower) (*entity.Borrower, error)
	Delete(deletedBorrower *entity.Borrower) error
}

type borrowerRepository struct {
	db *gorm.DB
}

func NewBorrowerRepository(db *gorm.DB) BorrowerRepository {
	return &borrowerRepository{db}
}

// ======================= INSERT ==============================
func (r *borrowerRepository) Insert(newBorrower *entity.Borrower) (*entity.Borrower, error) {
	currentTime := time.Now()
	newBorrower.Created_At = currentTime
	if err := r.db.Create(newBorrower).Error; err != nil {
		return nil, err
	}
	return newBorrower, nil
}

// ======================= FIND BY ID ==============================
func (r *borrowerRepository) FindByID(id int64) (*entity.Borrower, error) {
	var borrower entity.Borrower

	if err := r.db.Where("id = ?", id).Find(&borrower).Error; err != nil {
		return nil, err
	}
	return &borrower, nil
}

// ======================= FIND ALL ==============================
func (r *borrowerRepository) FindAll() ([]entity.Borrower, error) {
	var borrowers []entity.Borrower

	if err := r.db.Find(&borrowers).Error; err != nil {
		return nil, err
	}
	return borrowers, nil
}

// ======================= UPDATE ==============================
func (r *borrowerRepository) Update(updateBorrower *entity.Borrower) (*entity.Borrower, error) {
	var bor entity.Borrower
	if err := r.db.Where("id = ?", updateBorrower.ID).First(&bor).Error; err != nil {
		return nil, err
	}

	create_at := bor.Created_At

	bor.Username = updateBorrower.Username
	bor.Password = updateBorrower.Password
	bor.Name = updateBorrower.Name
	bor.Alamat = updateBorrower.Alamat
	bor.Phone_Number = updateBorrower.Phone_Number
	bor.Created_At = create_at

	if err := r.db.Save(&bor).Error; err != nil {
		return nil, err
	}
	return &bor, nil
}

// ======================= DELETE ==============================
func (r *borrowerRepository) Delete(deletedBorrower *entity.Borrower) error {
	if err := r.db.Delete(deletedBorrower).Error; err != nil {
		return err
	}
	return nil
}
