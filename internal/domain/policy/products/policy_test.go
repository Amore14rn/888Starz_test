package products

import (
	"context"
	"github.com/Amore14rn/888Starz_test/internal/domain/products/model"
	"github.com/Amore14rn/888Starz_test/internal/domain/products/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) All(ctx context.Context) ([]model.Products, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.Products), args.Error(1)
}

func (m *MockRepository) Create(ctx context.Context, req model.CreateProducts) (model.Products, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(model.Products), args.Error(1)
}

func (m *MockRepository) GetProduct(ctx context.Context, id string) (model.Products, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(model.Products), args.Error(1)
}

func (m *MockRepository) Update(ctx context.Context, req model.UpdateProducts) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type MockIdentityGenerator struct {
}

func (m MockIdentityGenerator) GenerateUUIDv4String() string {
	return "some_mocked_uuid"
}

type MockClock struct {
}

func (m MockClock) After(d time.Duration) <-chan time.Time {
	//TODO implement me
	panic("implement me")
}

func (m MockClock) Since(t time.Time) time.Duration {
	//TODO implement me
	panic("implement me")
}

func (m MockClock) Until(t time.Time) time.Duration {
	//TODO implement me
	panic("implement me")
}

func (m MockClock) Sleep(d time.Duration) {
	//TODO implement me
	panic("implement me")
}

func (m MockClock) Tick(d time.Duration) <-chan time.Time {
	//TODO implement me
	panic("implement me")
}

func (m MockClock) Now() time.Time {
	return time.Now()
}

func TestCreateProduct(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("Create", mock.Anything, mock.Anything).Return(model.Products{}, nil)

	policy := NewProductPolicy(NewProductService(mockRepo), MockIdentityGenerator{}, MockClock{})

	input := CreateProductInput{
		ID:          "mockedID",
		Description: "Test Product",
		Quantity:    5,
		Tags:        []string{"tag1", "tag2"},
		CreatedAt:   time.Now(),
	}

	output, err := policy.CreateProduct(context.Background(), input)

	assert.NoError(t, err)
	assert.NotNil(t, output.Product)
}

func NewProductService(repo *MockRepository) *service.ProductService {

	return &service.ProductService{
		Repository: repo,
	}
}

func TestAll(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("All", mock.Anything).Return([]model.Products{}, nil)

	policy := NewProductPolicy(NewProductService(mockRepo), MockIdentityGenerator{}, MockClock{})

	products, err := policy.All(context.Background())

	assert.NoError(t, err)
	assert.NotNil(t, products)
}

func TestGetProduct(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("GetProduct", mock.Anything, mock.Anything).Return(model.Products{}, nil)

	policy := NewProductPolicy(NewProductService(mockRepo), MockIdentityGenerator{}, MockClock{})

	input := GetProductInput{
		ID: "mockedID",
	}

	output, err := policy.GetProduct(context.Background(), input)

	assert.NoError(t, err)
	assert.NotNil(t, output.Product)
}

func TestUpdateProduct(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("GetProduct", mock.Anything, mock.Anything).Return(model.Products{}, nil)
	mockRepo.On("Update", mock.Anything, mock.Anything).Return(nil)

	policy := NewProductPolicy(NewProductService(mockRepo), MockIdentityGenerator{}, MockClock{})

	input := UpdateProductInput{
		ID:          "mockedID",
		Description: "Updated Test Product",
		Quantity:    10,
		Tags:        []string{"tag1", "tag2"},
		UpdatedAt:   time.Now(),
	}

	_, err := policy.UpdateProduct(context.Background(), input)

	assert.NoError(t, err)
}

func TestDeleteProduct(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("GetProduct", mock.Anything, mock.Anything).Return(model.Products{}, nil)
	mockRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

	policy := NewProductPolicy(NewProductService(mockRepo), MockIdentityGenerator{}, MockClock{})

	input := DeleteProductInput{
		ID: "mockedID",
	}

	_, err := policy.DeleteProduct(context.Background(), input)

	assert.NoError(t, err)
}
