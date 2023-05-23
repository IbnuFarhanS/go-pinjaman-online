package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/usecase"
	"github.com/gin-gonic/gin"
)

type BorrowerHandler interface {
	Insert(c *gin.Context)
	FindByID(c *gin.Context)
	FindAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type borrowerHandler struct {
	borrowerUsecase usecase.BorrowerUsecase
}

// Delete implements BorrowerHandler
func (h *borrowerHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	borrower := entity.Borrower{ID: id}
	err = h.borrowerUsecase.Delete(&borrower)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	msg := fmt.Sprintf("User with id %d has been deleted", id)
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

// FindAll implements BorrowerHandler
func (h *borrowerHandler) FindAll(c *gin.Context) {
	result, err := h.borrowerUsecase.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

// FindByID implements BorrowerHandler
func (h *borrowerHandler) FindByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.borrowerUsecase.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

// Insert implements BorrowerHandler
func (h *borrowerHandler) Insert(c *gin.Context) {
	var bor entity.Borrower
	if err := c.ShouldBindJSON(&bor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	result, err := h.borrowerUsecase.Insert(&bor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

// Update implements BorrowerHandler
func (h *borrowerHandler) Update(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var bor entity.Borrower
	bor.ID = id

	if err := c.ShouldBindJSON(&bor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result, err := h.borrowerUsecase.Update(&bor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func NewBorrowerHandler(borrowerUsecase usecase.BorrowerUsecase) BorrowerHandler {
	return &borrowerHandler{borrowerUsecase}
}
