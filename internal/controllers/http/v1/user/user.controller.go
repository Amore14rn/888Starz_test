package user

import (
	"github.com/Amore14rn/888Starz_test/internal/domain/policy/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	policy *user.Policy
}

func NewUserHandler(policy *user.Policy) *UserHandler {
	return &UserHandler{
		policy: policy,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var input user.CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userOutput, err := h.policy.CreateUser(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": userOutput.User})
}

func (h *UserHandler) All(c *gin.Context) {
	users, err := h.policy.All(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	var input user.GetUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userOutput, err := h.policy.GetUser(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": userOutput.User})
}

func (h *UserHandler) GetUserByName(c *gin.Context) {
	var input user.GetUserByNameInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userOutput, err := h.policy.GetUserByName(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": userOutput.User})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var input user.UpdateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userOutput, err := h.policy.UpdateUser(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": userOutput.User})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	var input user.DeleteUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.policy.DeleteUser(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *UserHandler) CreateOrder(c *gin.Context) {
	var input user.CreateOrderInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderOutput, err := h.policy.CreateOrder(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"order": orderOutput.Order})
}
