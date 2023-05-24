package repository

import (
	"errors"
	"time"

	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
	"gorm.io/gorm"
)

type LenderRepository interface {
	Insert(newLender *entity.Lender) (*entity.Lender, error)
	FindByID(id int64) (*entity.Lender, error)
	FindByName(name string) (*entity.Lender, error)
	FindAll() ([]entity.Lender, error)
	Update(updateLender *entity.Lender) (*entity.Lender, error)
	Delete(deletedLender *entity.Lender) error
}

type lenderRepository struct {
	db *gorm.DB
}

func NewLenderRepository(db *gorm.DB) LenderRepository {
	return &lenderRepository{db}
}

// ======================= INSERT ==============================
func (r *lenderRepository) Insert(newLender *entity.Lender) (*entity.Lender, error) {
	currentTime := time.Now()
	newLender.Created_At = currentTime
	if err := r.db.Create(newLender).Error; err != nil {
		return nil, err
	}
	return newLender, nil
}

// ======================= FIND BY ID ==============================
func (r *lenderRepository) FindByID(id int64) (*entity.Lender, error) {
	var lender entity.Lender

	if err := r.db.Where("id = ?", id).Find(&lender).Error; err != nil {
		return nil, err
	}
	return &lender, nil
}

// ======================= FIND BY NAME ==============================
func (r *lenderRepository) FindByName(name string) (*entity.Lender, error) {
	var lender entity.Lender

	if err := r.db.Where("name = ?", name).First(&lender).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &lender, nil
}

// ======================= FIND ALL ==============================
func (r *lenderRepository) FindAll() ([]entity.Lender, error) {
	var lenders []entity.Lender

	if err := r.db.Find(&lenders).Error; err != nil {
		return nil, err
	}
	return lenders, nil
}

// ======================= UPDATE ==============================
func (r *lenderRepository) Update(updateLender *entity.Lender) (*entity.Lender, error) {
	var len entity.Lender
	if err := r.db.Where("id = ?", updateLender.ID).First(&len).Error; err != nil {
		return nil, err
	}

	create_at := len.Created_At

	len.Name = updateLender.Name
	len.Created_At = create_at

	if err := r.db.Save(&len).Error; err != nil {
		return nil, err
	}
	return &len, nil
}

// ======================= DELETE ==============================
func (r *lenderRepository) Delete(deletedLender *entity.Lender) error {
	if err := r.db.Delete(deletedLender).Error; err != nil {
		return err
	}
	return nil
}
