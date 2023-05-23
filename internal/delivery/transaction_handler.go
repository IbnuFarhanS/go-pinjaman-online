package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/usecase"
	"github.com/gin-gonic/gin"
)

type TransactionHandler interface {
	Insert(c *gin.Context)
	FindByID(c *gin.Context)
	FindAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type transactionHandler struct {
	transactionUsecase usecase.TransactionUsecase
}

// Delete implements TransactionHandler
func (h *transactionHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tra := entity.Transaction{ID: id}
	err = h.transactionUsecase.Delete(&tra)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	msg := fmt.Sprintf("User with id %d has been deleted", id)
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

// FindAll implements TransactionHandler
func (h *transactionHandler) FindAll(c *gin.Context) {
	result, err := h.transactionUsecase.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

// FindByID implements TransactionHandler
func (h *transactionHandler) FindByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.transactionUsecase.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

// Insert implements TransactionHandler
func (h *transactionHandler) Insert(c *gin.Context) {
	var tra entity.Transaction
	if err := c.ShouldBindJSON(&tra); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	result, err := h.transactionUsecase.Insert(&tra)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

// Update implements TransactionHandler
func (h *transactionHandler) Update(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var tra entity.Transaction
	tra.ID = id

	if err := c.ShouldBindJSON(&tra); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result, err := h.transactionUsecase.Update(&tra)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func NewTransactionHandler(transactionUsecase usecase.TransactionUsecase) TransactionHandler {
	return &transactionHandler{transactionUsecase}
}
