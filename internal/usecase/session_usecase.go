package usecase

import (
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/repository"
)

type SessionUsecase interface {
	Insert(newSession *entity.Session) (*entity.Session, error)
	FindByID(id int64) (*entity.Session, error)
	FindAll() ([]entity.Session, error)
	Update(updateSession *entity.Session) (*entity.Session, error)
	Delete(deletedSession *entity.Session) error
}

type sessionUsecase struct {
	sessionRepo repository.SessionRepository
}

func NewSessionUsecase(sessionRepo repository.SessionRepository) SessionUsecase {
	return &sessionUsecase{sessionRepo}
}

func (u *sessionUsecase) Insert(newSession *entity.Session) (*entity.Session, error) {
	return u.sessionRepo.Insert(newSession)
}

func (u *sessionUsecase) FindByID(id int64) (*entity.Session, error) {
	return u.sessionRepo.FindByID(id)
}

func (u *sessionUsecase) FindAll() ([]entity.Session, error) {
	return u.sessionRepo.FindAll()
}

func (u *sessionUsecase) Update(updateLender *entity.Session) (*entity.Session, error) {
	return u.sessionRepo.Update(updateLender)
}

func (u *sessionUsecase) Delete(deletedLender *entity.Session) error {
	return u.sessionRepo.Delete(deletedLender)
}
