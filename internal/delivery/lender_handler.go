package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/IbnuFarhanS/go-pinjaman-online/internal/entity"
	"github.com/IbnuFarhanS/go-pinjaman-online/internal/usecase"
	"github.com/gin-gonic/gin"
)

type LenderHandler interface {
	Insert(c *gin.Context)
	FindByID(c *gin.Context)
	FindAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type lenderHandler struct {
	lenderUsecase usecase.LenderUsecase
}

// Delete implements LenderHandler
func (h *lenderHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lender := entity.Lender{ID: id}
	err = h.lenderUsecase.Delete(&lender)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	msg := fmt.Sprintf("User with id %d has been deleted", id)
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

// FindAll implements LenderHandler
func (h *lenderHandler) FindAll(c *gin.Context) {
	result, err := h.lenderUsecase.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

// FindByID implements LenderHandler
func (h *lenderHandler) FindByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.lenderUsecase.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

// Insert implements LenderHandler
func (h *lenderHandler) Insert(c *gin.Context) {
	var len entity.Lender
	if err := c.ShouldBindJSON(&len); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	result, err := h.lenderUsecase.Insert(&len)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

// Update implements LenderHandler
func (h *lenderHandler) Update(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var len entity.Lender
	len.ID = id

	if err := c.ShouldBindJSON(&len); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result, err := h.lenderUsecase.Update(&len)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func NewLenderHandler(lenderUsecase usecase.LenderUsecase) LenderHandler {
	return &lenderHandler{lenderUsecase}
}
