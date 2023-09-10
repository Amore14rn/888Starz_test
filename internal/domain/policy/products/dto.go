package products

import (
	"github.com/Amore14rn/888Starz_test/internal/domain/products/model"
	"time"
)

type CreateProductInput struct {
	ID          string
	Description string
	Quantity    int
	Tags        []string
	CreatedAt   time.Time
}

func NewCreateProductInput(id, description string, quantity int, tags []string, createdAt time.Time) CreateProductInput {
	return CreateProductInput{
		ID:          id,
		Description: description,
		Quantity:    quantity,
		Tags:        tags,
		CreatedAt:   createdAt,
	}
}

type CreateProductOutput struct {
	Product model.Products
}

type UpdateProductInput struct {
	ID          string
	Description string
	Quantity    int
	Tags        []string
	UpdatedAt   time.Time
}

func NewUpdateProductInput(id, description string, quantity int, tags []string, updatedAt time.Time) UpdateProductInput {
	return UpdateProductInput{
		ID:          id,
		Description: description,
		Quantity:    quantity,
		Tags:        tags,
		UpdatedAt:   updatedAt,
	}
}

type UpdateProductOutput struct {
	Product model.Products
}

type GetProductInput struct {
	ID string
}

func NewGetProductInput(id string) GetProductInput {
	return GetProductInput{
		ID: id,
	}
}

type GetProductOutput struct {
	Product model.Products
}

type DeleteProductInput struct {
	ID string
}

func NewDeleteProductInput(id string) DeleteProductInput {
	return DeleteProductInput{
		ID: id,
	}
}

type DeleteProductOutput struct {
	Product model.Products
}
