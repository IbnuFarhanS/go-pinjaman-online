package repository

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BorrowerRepositoryTestSuite struct {
	suite.Suite
	dbMock   sqlmock.Sqlmock
	gormDB   *gorm.DB
	repo     BorrowerRepository
	borrower *entity.Borrower
}

func (suite *BorrowerRepositoryTestSuite) SetupTest() {
	db, sqlMock, err := sqlmock.New()
	if err != nil {
		suite.T().Fatal("db connection error:", err)
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		suite.T().Fatal("gorm open error:", err)
	}

	suite.dbMock = sqlMock
	suite.gormDB = gormDB
	repo := NewBorrowerRepository(gormDB)
	suite.repo = repo

	// Prepare test data
	suite.borrower = &entity.Borrower{
		ID:           1,
		Username:     "borrower 1",
		Password:     "password 1",
		Name:         "user 1",
		Alamat:       "Bandung",
		Phone_Number: "084758789568",
		Created_At:   time.Now(),
	}
}

func (suite *BorrowerRepositoryTestSuite) TearDownTest() {
	err := suite.dbMock.ExpectationsWereMet()
	if err != nil {
		suite.T().Errorf("expectations were not met: %s", err.Error())
	}

	sqlDB, err := suite.gormDB.DB()
	if err != nil {
		suite.T().Errorf("failed to get SQL DB: %s", err.Error())
	}

	err = sqlDB.Close()
	if err != nil {
		suite.T().Errorf("failed to close gormDB: %s", err.Error())
	}
}

func TestBorrowerRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(BorrowerRepositoryTestSuite))
}

func (suite *BorrowerRepositoryTestSuite) TestBorrowerRepository_Insert_Success() {
	suite.dbMock.ExpectBegin()
	suite.dbMock.ExpectExec("INSERT INTO `borrowers`").
		WithArgs(suite.borrower.ID, suite.borrower.Username, suite.borrower.Password, suite.borrower.Name, suite.borrower.Alamat, suite.borrower.Phone_Number, suite.borrower.Created_At).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.dbMock.ExpectCommit()

	result, err := suite.repo.Insert(suite.borrower)

	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), suite.borrower, result)
}

func (suite *BorrowerRepositoryTestSuite) TestBorrowerRepository_Insert_Failed() {
	suite.dbMock.ExpectBegin()
	suite.dbMock.ExpectExec("INSERT INTO `borrowers`").
		WithArgs(suite.borrower.ID, suite.borrower.Username, suite.borrower.Password, suite.borrower.Name, suite.borrower.Alamat, suite.borrower.Phone_Number, suite.borrower.Created_At).
		WillReturnError(errors.New("database error"))
	suite.dbMock.ExpectRollback()

	result, err := suite.repo.Insert(suite.borrower)

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
}

//
