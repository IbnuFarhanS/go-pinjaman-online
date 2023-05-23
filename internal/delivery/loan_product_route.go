package delivery

import (
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/usecase"
	"github.com/gin-gonic/gin"
)

type LoanProductRouter struct {
	loanProductHandler LoanProductHandler
	publicRoute        *gin.RouterGroup
}

func (r *LoanProductRouter) SetupRouter() {
	loanProduct := r.publicRoute.Group("/loanproducts")
	{
		loanProduct.POST("/", r.loanProductHandler.Insert)
		loanProduct.PUT("/:id", r.loanProductHandler.Update)
		loanProduct.DELETE("/:id", r.loanProductHandler.Delete)
		loanProduct.GET("/:id", r.loanProductHandler.FindByID)
		loanProduct.GET("/", r.loanProductHandler.FindAll)
	}
}

func NewLoanProductRouter(publicRoute *gin.RouterGroup, loanProductUsecase usecase.LoanProductUsecase) {
	loanProductHandler := NewLoanProductHandler(loanProductUsecase)
	router := LoanProductRouter{
		loanProductHandler,
		publicRoute,
	}
	router.SetupRouter()
}
