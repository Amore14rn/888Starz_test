package service

import (
	"context"
	"github.com/Amore14rn/888Starz/internal/domain/products/model"
	"github.com/Amore14rn/888Starz/pkg/errors"
)

type repository interface {
	All(ctx context.Context) ([]model.Products, error)
	Create(ctx context.Context, req model.CreateProducts) (model.Products, error)
	GetProduct(ctx context.Context, id string) (model.Products, error)
	Update(ctx context.Context, req model.UpdateProducts) error
	Delete(ctx context.Context, id string) error
}

type ProductService struct {
	repository repository
}

func NewProductService(repository repository) *ProductService {
	return &ProductService{
		repository: repository,
	}
}

func (s *ProductService) All(ctx context.Context) ([]model.Products, error) {
	products, err := s.repository.All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "repository.All")
	}

	return products, nil
}

func (s *ProductService) CreateProduct(ctx context.Context, req model.CreateProducts) (model.Products, error) {
	product, err := s.repository.Create(ctx, req)
	if err != nil {
		return model.Products{}, errors.Wrap(err, "repository.CreateProduct")
	}

	return product, nil
}

func (s *ProductService) GetProduct(ctx context.Context, id string) (model.Products, error) {
	product, err := s.repository.GetProduct(ctx, id)
	if err != nil {
		return model.Products{}, errors.Wrap(err, "repository.GetProduct")
	}

	return product, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, req model.UpdateProducts) error {
	if err := s.repository.Update(ctx, req); err != nil {
		return errors.Wrap(err, "repository.UpdateProduct")
	}

	return nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id string) error {
	if err := s.repository.Delete(ctx, id); err != nil {
		return errors.Wrap(err, "repository.DeleteProduct")
	}

	return nil
}
