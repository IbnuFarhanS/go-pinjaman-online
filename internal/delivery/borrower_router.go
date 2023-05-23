package delivery

import (
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/usecase"
	"github.com/gin-gonic/gin"
)

type BorrowerRouter struct {
	borrowerHandler BorrowerHandler
	publicRoute     *gin.RouterGroup
}

func (r *BorrowerRouter) SetupRouter() {
	borrowers := r.publicRoute.Group("/borrowers")
	{
		borrowers.POST("/", r.borrowerHandler.Insert)
		borrowers.PUT("/:id", r.borrowerHandler.Update)
		borrowers.DELETE("/:id", r.borrowerHandler.Delete)
		borrowers.GET("/:id", r.borrowerHandler.FindByID)
		borrowers.GET("/", r.borrowerHandler.FindAll)
	}
}

func NewBorrowerRouter(publicRoute *gin.RouterGroup, borrowerUsecase usecase.BorrowerUsecase) {
	borrowerHandler := NewBorrowerHandler(borrowerUsecase)
	router := BorrowerRouter{
		borrowerHandler,
		publicRoute,
	}
	router.SetupRouter()
}
