package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/CaiqueRibeiro/product-api/internal/dto"
	"github.com/CaiqueRibeiro/product-api/internal/entity"
	"github.com/CaiqueRibeiro/product-api/internal/infra/database"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Repository database.ProductRepositoryInterface
}

func NewProductHandlers(repository database.ProductRepositoryInterface) *ProductHandler {
	return &ProductHandler{Repository: repository}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product dto.CreateProductInput
	if c.Bind(&product) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	p, err := entity.NewProduct(product.Name, product.Price, time.Now())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Repository.Create(p) // Salva no banco
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *ProductHandler) FindAll(c *gin.Context) {
	var page string
	if page = c.Query("page"); page == "" {
		page = "1"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	var limit string
	if limit = c.Query("limit"); limit == "" {
		limit = "10"
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}

	var order string
	if order = c.Query("order"); order == "" {
		order = "10"
	}

	products, err := h.Repository.FindAll(pageInt, limitInt, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}
