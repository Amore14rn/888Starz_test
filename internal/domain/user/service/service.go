package service

import (
	"context"
	"github.com/Amore14rn/888Starz/internal/domain/user/model"
	"github.com/Amore14rn/888Starz/pkg/errors"
)

type repository interface {
	All(ctx context.Context) ([]model.User, error)
	Create(ctx context.Context, req model.CreateUser) error
	GetUser(ctx context.Context, id string) (model.User, error)
	GetUserByName(ctx context.Context, name string) (model.User, error)
	Update(ctx context.Context, req model.UpdateUser) error
	Delete(ctx context.Context, id string) error
	CreateOrder(ctx context.Context, req model.CreateOrder) error
	AreProductsAvailable(ctx context.Context, productID string, products []model.OrderProduct) bool
}

type UserService struct {
	repository repository
}

func NewUserService(repository repository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) All(ctx context.Context) ([]model.User, error) {
	users, err := s.repository.All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "repository.All")
	}

	return users, nil
}

func (u *UserService) CreateUser(ctx context.Context, req model.CreateUser) (model.User, error) {
	// Проверка возраста пользователя
	if req.Age < 18 {
		return model.User{}, errors.New("Пользователь должен быть не младше 18 лет")
	}

	// Проверка длины пароля
	if len(req.Password) < 8 {
		return model.User{}, errors.New("Пароль должен содержать не менее 8 символов")
	}

	// Проверка наличия цифр в пароле
	var isDigit bool
	for _, char := range req.Password {
		if char >= '0' && char <= '9' {
			isDigit = true
			break
		}
	}
	if !isDigit {
		return model.User{}, errors.New("Пароль должен содержать хотя бы одну цифру")
	}

	// Проверка наличия заглавных букв в пароле
	var isUpper bool
	for _, char := range req.Password {
		if char >= 'A' && char <= 'Z' {
			isUpper = true
			break
		}
	}
	if !isUpper {
		return model.User{}, errors.New("Пароль должен содержать хотя бы одну заглавную букву")
	}

	err := u.repository.Create(ctx, req)
	if err != nil {
		return model.User{}, err
	}
	return model.NewUser(
		req.ID,
		req.FirstName,
		req.LastName,
		req.Age,
		req.IsMarried,
		req.Password,
		req.Order,
		req.CreatedAt,
		nil), nil
}

func (u *UserService) GetUser(ctx context.Context, id string) (model.User, error) {
	user, err := u.repository.GetUser(ctx, id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *UserService) GetUserByName(ctx context.Context, name string) (model.User, error) {
	user, err := u.repository.GetUserByName(ctx, name)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *UserService) UpdateUser(ctx context.Context, req model.UpdateUser) (model.User, error) {
	err := u.repository.Update(ctx, req)
	if err != nil {
		return model.User{}, err
	}
	return model.NewUser(
		req.ID,
		req.FirstName,
		req.LastName,
		req.Age,
		req.IsMarried,
		req.Password,
		req.Order,
		req.UpdatedAt,
		nil), nil
}

func (u *UserService) DeleteUser(ctx context.Context, id string) error {
	err := u.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) CreateOrder(ctx context.Context, req model.CreateOrder) (model.CreateOrder, error) {
	err := u.repository.CreateOrder(ctx, req)
	if err != nil {
		return model.CreateOrder{}, err
	}
	return model.NewCreateOrder(
		req.ID,
		req.UserID,
		req.ProductID,
		req.Products,
		req.TimeStamp,
	), nil

}

func (u *UserService) AreProductsAvailable(ctx context.Context, productID string, products []model.OrderProduct) bool {
	return u.repository.AreProductsAvailable(ctx, productID, products)
}
