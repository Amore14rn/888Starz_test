package products

import (
	"context"
	"github.com/Amore14rn/888Starz_test/internal/domain/products/model"
	"github.com/Amore14rn/888Starz_test/internal/domain/products/service"
	"github.com/Amore14rn/888Starz_test/pkg/common/core/clock"
	"github.com/Amore14rn/888Starz_test/pkg/errors"
	"time"
)

type IdentityGenerator interface {
	GenerateUUIDv4String() string
}

type Clock interface {
	Now() time.Time
}

type Policy struct {
	productService *service.ProductService

	identity IdentityGenerator
	clock    Clock
}

func NewProductPolicy(productService *service.ProductService, identity IdentityGenerator, clock clock.Clock) *Policy {
	return &Policy{
		productService: productService,
		identity:       identity,
		clock:          clock,
	}
}

func (p *Policy) CreateProduct(ctx context.Context, input CreateProductInput) (CreateProductOutput, error) {

	if input.Description == "" {
		return CreateProductOutput{}, errors.New("Описание продукта обязательно")
	}

	// Проверка количества
	if input.Quantity <= 0 {
		return CreateProductOutput{}, errors.New("Количество продукта должно быть положительным числом")
	}

	// Создание продукта
	createProduct := model.NewCreateProducts(
		input.ID,
		input.Description,
		input.Quantity,
		input.Tags,
		input.CreatedAt,
	)

	product, err := p.productService.CreateProduct(ctx, createProduct)
	if err != nil {
		return CreateProductOutput{}, errors.Wrap(err, "Error when creating a product")
	}

	return CreateProductOutput{
		Product: product,
	}, nil
}

func (p *Policy) All(ctx context.Context) ([]model.Products, error) {
	products, err := p.productService.All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error when getting all products")
	}

	return products, nil
}

func (p *Policy) GetProduct(ctx context.Context, input GetProductInput) (GetProductOutput, error) {
	product, err := p.productService.GetProduct(ctx, input.ID)
	if err != nil {
		return GetProductOutput{}, errors.Wrap(err, "Error when getting product")
	}

	return GetProductOutput{
		Product: product,
	}, nil
}

func (p *Policy) UpdateProduct(ctx context.Context, input UpdateProductInput) (UpdateProductOutput, error) {
	// Проверка на существование продукта
	_, err := p.productService.GetProduct(ctx, input.ID)
	if err != nil {
		return UpdateProductOutput{}, errors.Wrap(err, "Error when getting product")
	}

	// Проверка на количество
	if input.Quantity <= 0 {
		return UpdateProductOutput{}, errors.New("Количество продукта должно быть положительным числом")
	}

	updateProduct := model.NewUpdateProducts(
		input.ID,
		input.Description,
		input.Quantity,
		input.Tags,
		input.UpdatedAt,
	)

	err = p.productService.UpdateProduct(ctx, updateProduct)
	if err != nil {
		return UpdateProductOutput{}, errors.Wrap(err, "Error when updating product")
	}

	return UpdateProductOutput{}, nil
}

func (p *Policy) DeleteProduct(ctx context.Context, input DeleteProductInput) (DeleteProductOutput, error) {
	// Проверка на существование продукта
	_, err := p.productService.GetProduct(ctx, input.ID)
	if err != nil {
		return DeleteProductOutput{}, errors.Wrap(err, "Error when getting product")
	}

	err = p.productService.DeleteProduct(ctx, input.ID)
	if err != nil {
		return DeleteProductOutput{}, errors.Wrap(err, "Error when deleting product")
	}

	return DeleteProductOutput{}, nil
}
