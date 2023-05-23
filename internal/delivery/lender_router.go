package delivery

import (
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/usecase"
	"github.com/gin-gonic/gin"
)

type LenderRouter struct {
	lenderHandler LenderHandler
	publicRoute   *gin.RouterGroup
}

func (r *LenderRouter) SetupRouter() {
	lenders := r.publicRoute.Group("/lenders")
	{
		lenders.POST("/", r.lenderHandler.Insert)
		lenders.PUT("/:id", r.lenderHandler.Update)
		lenders.DELETE("/:id", r.lenderHandler.Delete)
		lenders.GET("/:id", r.lenderHandler.FindByID)
		lenders.GET("/", r.lenderHandler.FindAll)
	}
}

func NewLenderRouter(publicRoute *gin.RouterGroup, lenderUsecase usecase.LenderUsecase) {
	lenderHandler := NewLenderHandler(lenderUsecase)
	router := LenderRouter{
		lenderHandler,
		publicRoute,
	}
	router.SetupRouter()
}
