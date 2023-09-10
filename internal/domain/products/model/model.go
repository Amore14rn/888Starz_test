package model

import "time"

type Products struct {
	ID          string
	Description string
	Quantity    int
	Tags        []string
	CreatedAt   time.Time
	UpdatedAt   *time.Time // Если есть поле "updated_at"
}

func NewProduct(id, description string, quantity int, tags []string, createdAt time.Time, updatedAt *time.Time) Products {
	return Products{
		ID:          id,
		Description: description,
		Quantity:    quantity,
		Tags:        tags,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

type CreateProducts struct {
	ID          string
	Description string
	Quantity    int
	Tags        []string
	CreatedAt   time.Time
}

func NewCreateProducts(id, description string, quantity int, tags []string, createdAt time.Time) CreateProducts {
	return CreateProducts{
		ID:          id,
		Description: description,
		Quantity:    quantity,
		Tags:        tags,
		CreatedAt:   createdAt,
	}
}

type ProductHistory struct {
	ProductID string
	Price     float64
	Timestamp time.Time
}

type Order struct {
	ID        string
	UserID    string
	Products  []OrderProduct
	Timestamp time.Time
}

type OrderProduct struct {
	ProductID string
	Quantity  int
	Price     float64
}

func (o *Order) AddProduct(product OrderProduct) {
	o.Products = append(o.Products, product)
}

type UpdateProducts struct {
	ID          string
	Description string
	Quantity    int
	Tags        []string
	UpdatedAt   time.Time
}

func NewUpdateProducts(id,
	description string,
	quantity int,
	tags []string,
	updatedAt time.Time) UpdateProducts {
	return UpdateProducts{
		ID:          id,
		Description: description,
		Quantity:    quantity,
		Tags:        tags,
		UpdatedAt:   updatedAt,
	}
}
