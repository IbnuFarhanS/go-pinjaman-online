package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/usecase"
	"github.com/gin-gonic/gin"
)

type LoanProductHandler interface {
	Insert(c *gin.Context)
	FindByID(c *gin.Context)
	FindAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type loanProductHandler struct {
	loanProductUsecase usecase.LoanProductUsecase
}

// Delete implements LoanProductHandler
func (h *loanProductHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loanProduct := entity.Loan_Product{ID: id}
	err = h.loanProductUsecase.Delete(&loanProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	msg := fmt.Sprintf("User with id %d has been deleted", id)
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

// FindAll implements LoanProductHandler
func (h *loanProductHandler) FindAll(c *gin.Context) {
	result, err := h.loanProductUsecase.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

// FindByID implements LoanProductHandler
func (h *loanProductHandler) FindByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.loanProductUsecase.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

// Insert implements LoanProductHandler
func (h *loanProductHandler) Insert(c *gin.Context) {
	var lp entity.Loan_Product
	if err := c.ShouldBindJSON(&lp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	result, err := h.loanProductUsecase.Insert(&lp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

// Update implements LoanProductHandler
func (h *loanProductHandler) Update(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var lp entity.Loan_Product
	lp.ID = id

	if err := c.ShouldBindJSON(&lp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result, err := h.loanProductUsecase.Update(&lp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func NewLoanProductHandler(loanProductUsecase usecase.LoanProductUsecase) LoanProductHandler {
	return &loanProductHandler{loanProductUsecase}
}
