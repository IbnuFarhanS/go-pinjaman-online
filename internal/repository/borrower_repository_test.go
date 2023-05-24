package repository

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func prepareTestData() *entity.Borrower {
	currentTime := time.Now()
	newBorrower := &entity.Borrower{
		ID:           1,
		Username:     "user1",
		Password:     "password1",
		Name:         "John Doe",
		Alamat:       "123 Main St",
		Phone_Number: "123456789",
		Created_At:   currentTime,
	}
	return newBorrower
}

func TestBorrowerRepository_Insert_success(t *testing.T) {
	newBorrower := prepareTestData()

	// Prepare GORM Mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	// Create borrower repository instance
	repo := NewBorrowerRepository(gormDB)

	// Unit test
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO borrowers").
		WithArgs(newBorrower.ID, newBorrower.Username, newBorrower.Password, newBorrower.Name, newBorrower.Alamat, newBorrower.Phone_Number, newBorrower.Created_At).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	insertedBorrower, err := repo.Insert(newBorrower)

	assert.NoError(t, err)
	assert.Equal(t, newBorrower, insertedBorrower)
}

func TestBorrowerRepository_Insert_failed(t *testing.T) {
	newBorrower := prepareTestData()

	// Prepare GORM Mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	// Create borrower repository instance
	repo := NewBorrowerRepository(gormDB)

	// Unit test
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO borrowers").
		WithArgs(newBorrower.ID, newBorrower.Username, newBorrower.Password, newBorrower.Name, newBorrower.Alamat, newBorrower.Phone_Number, newBorrower.Created_At).
		WillReturnError(errors.New("some error"))
	mock.ExpectRollback()

	insertedBorrower, err := repo.Insert(newBorrower)

	assert.Error(t, err)
	assert.Nil(t, insertedBorrower)
}
