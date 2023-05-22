package usecase

import (
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/repository"
)

type LenderUsecase interface {
	Insert(newLender *entity.Lender) (*entity.Lender, error)
	FindByID(id int64) (*entity.Lender, error)
	FindAll() ([]entity.Lender, error)
	Update(updateLender *entity.Lender) (*entity.Lender, error)
	Delete(deletedLender *entity.Lender) error
}

type lenderUsecase struct {
	lenderRepo repository.LenderRepository
}

func NewLenderUsecase(lenderRepo repository.LenderRepository) LenderUsecase {
	return &lenderUsecase{lenderRepo}
}

func (u *lenderUsecase) Insert(newLender *entity.Lender) (*entity.Lender, error) {
	return u.lenderRepo.Insert(newLender)
}

func (u *lenderUsecase) FindByID(id int64) (*entity.Lender, error) {
	return u.lenderRepo.FindByID(id)
}

func (u *lenderUsecase) FindAll() ([]entity.Lender, error) {
	return u.lenderRepo.FindAll()
}

func (u *lenderUsecase) Update(updateLender *entity.Lender) (*entity.Lender, error) {
	return u.lenderRepo.Update(updateLender)
}

func (u *lenderUsecase) Delete(deletedLender *entity.Lender) error {
	return u.lenderRepo.Delete(deletedLender)
}
