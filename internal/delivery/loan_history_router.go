package delivery

import (
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/usecase"
	"github.com/gin-gonic/gin"
)

type LoanHistoryRouter struct {
	loanHistoryHandler LoanHistoryHandler
	publicRoute        *gin.RouterGroup
}

func (r *LoanHistoryRouter) SetupRouter() {
	loanHistory := r.publicRoute.Group("/loanhistorys")
	{
		loanHistory.POST("/", r.loanHistoryHandler.Insert)
		loanHistory.PUT("/:id", r.loanHistoryHandler.Update)
		loanHistory.DELETE("/:id", r.loanHistoryHandler.Delete)
		loanHistory.GET("/:id", r.loanHistoryHandler.FindByID)
		loanHistory.GET("/", r.loanHistoryHandler.FindAll)
	}
}

func NewLoanHistoryRouter(publicRoute *gin.RouterGroup, loanHistoryUsecase usecase.LoanHistoryUsecase) {
	loanHistoryHandler := NewLoanHistoryHandler(loanHistoryUsecase)
	router := LoanHistoryRouter{
		loanHistoryHandler,
		publicRoute,
	}
	router.SetupRouter()
}
