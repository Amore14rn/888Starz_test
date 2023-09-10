package dao

import (
	"github.com/Amore14rn/888Starz/internal/domain/products/model"
	"time"
)

type ProductStorage struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	Tags        []string  `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
}

func (ps *ProductStorage) ToDomain() model.Products {
	return model.Products{
		ID:          ps.ID,
		Description: ps.Description,
		Quantity:    ps.Quantity,
		Tags:        ps.Tags,
		CreatedAt:   ps.CreatedAt,
	}
}
