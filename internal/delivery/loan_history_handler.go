package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/usecase"
	"github.com/gin-gonic/gin"
)

type LoanHistoryHandler interface {
	Insert(c *gin.Context)
	FindByID(c *gin.Context)
	FindAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type loanHistoryHandler struct {
	loanHistoryUsecase usecase.LoanHistoryUsecase
}

// Delete implements LoanHistoryHandler
func (h *loanHistoryHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loanHistory := entity.Loan_History{ID: id}
	err = h.loanHistoryUsecase.Delete(&loanHistory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	msg := fmt.Sprintf("User with id %d has been deleted", id)
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

// FindAll implements LoanHistoryHandler
func (h *loanHistoryHandler) FindAll(c *gin.Context) {
	result, err := h.loanHistoryUsecase.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

// FindByID implements LoanHistoryHandler
func (h *loanHistoryHandler) FindByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.loanHistoryUsecase.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

// Insert implements LoanHistoryHandler
func (h *loanHistoryHandler) Insert(c *gin.Context) {
	var ls entity.Loan_History
	if err := c.ShouldBindJSON(&ls); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	result, err := h.loanHistoryUsecase.Insert(&ls)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

// Update implements LoanHistoryHandler
func (h *loanHistoryHandler) Update(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var ls entity.Loan_History
	ls.ID = id

	if err := c.ShouldBindJSON(&ls); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result, err := h.loanHistoryUsecase.Update(&ls)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func NewLoanHistoryHandler(loanHistoryUsecase usecase.LoanHistoryUsecase) LoanHistoryHandler {
	return &loanHistoryHandler{loanHistoryUsecase}
}
