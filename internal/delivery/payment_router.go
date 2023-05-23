package delivery

import (
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/usecase"
	"github.com/gin-gonic/gin"
)

type PaymentRouter struct {
	paymentHandler PaymentHandler
	publicRoute    *gin.RouterGroup
}

func (r *PaymentRouter) SetupRouter() {
	payment := r.publicRoute.Group("/payments")
	{
		payment.POST("/", r.paymentHandler.Insert)
		payment.PUT("/:id", r.paymentHandler.Update)
		payment.DELETE("/:id", r.paymentHandler.Delete)
		payment.GET("/:id", r.paymentHandler.FindByID)
		payment.GET("/", r.paymentHandler.FindAll)
	}
}

func NewPaymentRouter(publicRoute *gin.RouterGroup, paymentUsecase usecase.PaymentUsecase) {
	paymentHandler := NewPaymentHandler(paymentUsecase)
	router := PaymentRouter{
		paymentHandler,
		publicRoute,
	}
	router.SetupRouter()
}
