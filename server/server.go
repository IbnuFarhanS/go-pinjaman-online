package server

import (
	"net/http"

	"github.com/IbnuFarhanS/go-pinjaman-online/internal/delivery"
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/repository"
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/usecase"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type server struct {
	router         *gin.Engine
	borrowerRouter delivery.BorrowerRouter
}

func NewServer() *server {
	return &server{}
}

func (s *server) Init(ConnStr string) error {
	//init db conn
	dsn := ConnStr
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&entity.Borrower{})

	borrowerRepo := repository.NewBorrowerRepository(db)
	borrowerUsecase := usecase.NewBorrowerUsecase(borrowerRepo)

	// lenderRepo := repository.NewLenderRepository(db)
	// lenderUsecase := usecase.NewLenderUsecase(lenderRepo)

	// loanHistoryRepo := repository.NewLoanHistoryRepository(db)
	// loanHistoryUsecase := usecase.NewLoanHistoryUsecase(loanHistoryRepo)

	// loanProductRepo := repository.NewLoanProductRepository(db)
	// loanProductUsecase := usecase.NewLoanProductUsecase(loanProductRepo)

	// paymentRepo := repository.NewPaymentRepository(db)
	// paymentUsecase := usecase.NewPaymentUsecase(paymentRepo)

	// traRepo := repository.NewTransactionRepository(db)
	// traUsecase := usecase.NewTransactionUsecase(traRepo)

	r := gin.Default()
	api := r.Group("/api")
	r.POST("/auth/login", delivery.LoginHandler)
	api.Use(delivery.AuthMiddleware())

	delivery.NewBorrowerRouter(api, borrowerUsecase)
	// delivery.NewLenderRouter(api, lenderUsecase)
	// delivery.NewLoanHistoryRouter(api, loanHistoryUsecase)
	// delivery.NewLoanProductRouter(api, loanProductUsecase)
	// delivery.NewPaymentRouter(api, paymentUsecase)
	// delivery.NewTransactionRouter(api, traUsecase)

	s.router = r

	return nil
}

func (s *server) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
