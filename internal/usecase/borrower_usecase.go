package usecase

import (
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type BorrowerUsecase interface {
	Insert(newBorrower *entity.Borrower) (*entity.Borrower, error)
	FindByID(id int64) (*entity.Borrower, error)
	FindAll() ([]entity.Borrower, error)
	Update(updateBorrower *entity.Borrower) (*entity.Borrower, error)
	Delete(deletedBorrower *entity.Borrower) error
}

type borrowerUsecase struct {
	borrowerRepo repository.BorrowerRepository
}

func NewBorrowerUsecase(borrowerRepo repository.BorrowerRepository) BorrowerUsecase {
	return &borrowerUsecase{borrowerRepo}
}

func (u *borrowerUsecase) Insert(newBorrower *entity.Borrower) (*entity.Borrower, error) {
	// Business Logic
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(newBorrower.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newBorrower.Password = string(encryptedPassword)

	// Tulis ke db melalui repo
	return u.borrowerRepo.Insert(newBorrower)
}

func (u *borrowerUsecase) FindByID(id int64) (*entity.Borrower, error) {
	return u.borrowerRepo.FindByID(id)
}

func (u *borrowerUsecase) FindAll() ([]entity.Borrower, error) {
	return u.borrowerRepo.FindAll()
}

func (u *borrowerUsecase) Update(updateBorrower *entity.Borrower) (*entity.Borrower, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(updateBorrower.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	updateBorrower.Password = string(encryptedPassword)

	return u.borrowerRepo.Update(updateBorrower)
}

func (u *borrowerUsecase) Delete(deletedBorrower *entity.Borrower) error {
	return u.borrowerRepo.Delete(deletedBorrower)
}
