package product

import (
	"github.com/Amore14rn/888Starz_test/internal/domain/policy/products"
	"github.com/gin-gonic/gin"

	"net/http"
)

type ProductHandler struct {
	policy *products.Policy
}

func NewProductHandler(policy *products.Policy) *ProductHandler {
	return &ProductHandler{
		policy: policy,
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var input products.CreateProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productOutput, err := h.policy.CreateProduct(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"product": productOutput.Product})
}

func (h *ProductHandler) All(c *gin.Context) {
	products, err := h.policy.All(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	var input products.GetProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productOutput, err := h.policy.GetProduct(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": productOutput.Product})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	var input products.UpdateProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productOutput, err := h.policy.UpdateProduct(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": productOutput.Product})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	var input products.DeleteProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productOutput, err := h.policy.DeleteProduct(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": productOutput.Product})
}
