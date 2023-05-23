package delivery

import (
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/usecase"
	"github.com/gin-gonic/gin"
)

type TransactionRouter struct {
	transactionHandler TransactionHandler
	publicRoute        *gin.RouterGroup
}

func (r *TransactionRouter) SetupRouter() {
	tra := r.publicRoute.Group("/transactions")
	{
		tra.POST("/", r.transactionHandler.Insert)
		tra.PUT("/:id", r.transactionHandler.Update)
		tra.DELETE("/:id", r.transactionHandler.Delete)
		tra.GET("/:id", r.transactionHandler.FindByID)
		tra.GET("/", r.transactionHandler.FindAll)
	}
}

func NewTransactionRouter(publicRoute *gin.RouterGroup, transactionUsecase usecase.TransactionUsecase) {
	transactionHandler := NewTransactionHandler(transactionUsecase)
	router := TransactionRouter{
		transactionHandler,
		publicRoute,
	}
	router.SetupRouter()
}
