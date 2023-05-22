package usecase

import (
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/repository"
)

type PaymentUsecase interface {
	Insert(newPayment *entity.Payment) (*entity.Payment, error)
	FindByID(id int64) (*entity.Payment, error)
	FindAll() ([]entity.Payment, error)
	Update(updatePayment *entity.Payment) (*entity.Payment, error)
	Delete(deletedPayment *entity.Payment) error
}

type paymentUsecase struct {
	paymentRepo repository.PaymentRepository
}

func NewPaymentUsecase(paymentRepo repository.PaymentRepository) PaymentUsecase {
	return &paymentUsecase{paymentRepo}
}

func (u *paymentUsecase) Insert(newPayment *entity.Payment) (*entity.Payment, error) {
	return u.paymentRepo.Insert(newPayment)
}

func (u *paymentUsecase) FindByID(id int64) (*entity.Payment, error) {
	return u.paymentRepo.FindByID(id)
}

func (u *paymentUsecase) FindAll() ([]entity.Payment, error) {
	return u.paymentRepo.FindAll()
}

func (u *paymentUsecase) Update(updatePayment *entity.Payment) (*entity.Payment, error) {
	return u.paymentRepo.Update(updatePayment)
}

func (u *paymentUsecase) Delete(deletedPayment *entity.Payment) error {
	return u.paymentRepo.Delete(deletedPayment)
}
