package usecase

import (
	"errors"

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
	// Validate username
	if newBorrower.Username == "" {
		return nil, errors.New("username is required")
	}
	// Check if username already exists
	existingBorrower, err := u.borrowerRepo.FindByUsername(newBorrower.Username)
	if err != nil {
		return nil, err
	}
	if existingBorrower != nil {
		return nil, errors.New("username is already use")
	}

	// Validate password
	if newBorrower.Password == "" {
		return nil, errors.New("password is required")
	}
	// Validate name
	if newBorrower.Name == "" {
		return nil, errors.New("nama is required")
	}
	// Validate alamat
	if newBorrower.Alamat == "" {
		return nil, errors.New("alamat is required")
	}
	// Validate phone_number
	if newBorrower.Phone_Number == "" {
		return nil, errors.New("nomor telepon is required")
	}

	// Business Logic Password bcrypt
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(newBorrower.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newBorrower.Password = string(encryptedPassword)

	return u.borrowerRepo.Insert(newBorrower)
}

func (u *borrowerUsecase) FindByID(id int64) (*entity.Borrower, error) {
	return u.borrowerRepo.FindByID(id)
}

func (u *borrowerUsecase) FindAll() ([]entity.Borrower, error) {
	return u.borrowerRepo.FindAll()
}

func (u *borrowerUsecase) Update(updateBorrower *entity.Borrower) (*entity.Borrower, error) {
	// Validate username
	if updateBorrower.Username == "" {
		return nil, errors.New("username is required")
	}
	// Check if username already exists
	existingBorrower, err := u.borrowerRepo.FindByUsername(updateBorrower.Username)
	if err != nil {
		return nil, err
	}
	if existingBorrower != nil {
		return nil, errors.New("username is already use")
	}

	// Validate password
	if updateBorrower.Password == "" {
		return nil, errors.New("password is required")
	}
	// Validate name
	if updateBorrower.Name == "" {
		return nil, errors.New("nama is required")
	}
	// Validate alamat
	if updateBorrower.Alamat == "" {
		return nil, errors.New("alamat is required")
	}
	// Validate phone_number
	if updateBorrower.Phone_Number == "" {
		return nil, errors.New("nomor telepon is required")
	}

	// Business Logic Password bcrypt
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
