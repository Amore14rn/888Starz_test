package user

import (
	"context"
	"github.com/Amore14rn/888Starz/internal/domain/user/model"
	"github.com/Amore14rn/888Starz/internal/domain/user/service"
	"github.com/Amore14rn/888Starz/pkg/common/core/clock"
	"github.com/Amore14rn/888Starz/pkg/errors"
	"time"
)

type IdentityGenerator interface {
	GenerateUUIDv4String() string
}

type Clock interface {
	Now() time.Time
}

type Policy struct {
	userService *service.UserService

	identity IdentityGenerator
	clock    Clock
}

func NewUserPolicy(userService *service.UserService, identity IdentityGenerator, clock clock.Clock) *Policy {
	return &Policy{
		userService: userService,
		identity:    identity,
		clock:       clock,
	}
}

func (u *Policy) CreateUser(ctx context.Context, input CreateUserInput) (CreateUserOutput, error) {
	// Check user's age
	if input.Age < 18 {
		return CreateUserOutput{}, errors.New("Пользователь должен быть не младше 18 лет")
	}

	// Validate password
	if err := u.validatePassword(input.Password); err != nil {
		return CreateUserOutput{}, err
	}

	createUser := model.NewCreateUser(
		u.identity.GenerateUUIDv4String(),
		input.FirstName,
		input.LastName,
		input.Age,
		input.IsMarried,
		input.Password,
		input.Order,
		u.clock.Now(),
	)

	user, err := u.userService.CreateUser(ctx, createUser)
	if err != nil {
		return CreateUserOutput{}, errors.Wrap(err, "Error when creating a user")
	}

	return CreateUserOutput{
		User: user,
	}, nil
}

func (u *Policy) validatePassword(password string) error {
	// Check password length
	if len(password) < 8 {
		return errors.New("Пароль должен содержать не менее 8 символов")
	}

	// Check for the presence of digits in the password
	var hasDigit bool
	for _, char := range password {
		if char >= '0' && char <= '9' {
			hasDigit = true
			break
		}
	}
	if !hasDigit {
		return errors.New("Пароль должен содержать хотя бы одну цифру")
	}

	// Check for the presence of uppercase letters in the password
	var hasUpper bool
	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			hasUpper = true
			break
		}
	}
	if !hasUpper {
		return errors.New("Пароль должен содержать хотя бы одну заглавную букву")
	}

	return nil
}

func (u *Policy) All(ctx context.Context) ([]model.User, error) {
	users, err := u.userService.All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error when getting all users")
	}

	return users, nil
}

func (u *Policy) GetUser(ctx context.Context, input GetUserInput) (GetUserOutput, error) {
	user, err := u.userService.GetUser(ctx, input.ID)
	if err != nil {
		return GetUserOutput{}, errors.Wrap(err, "Error when getting user")
	}

	return GetUserOutput{
		User: user,
	}, nil
}

func (u *Policy) GetUserByName(ctx context.Context, input GetUserByNameInput) (GetUserByNameOutput, error) {
	user, err := u.userService.GetUserByName(ctx, input.FirstName)
	if err != nil {
		return GetUserByNameOutput{}, errors.Wrap(err, "Error when getting user")
	}

	return GetUserByNameOutput{
		User: []model.User{user},
	}, nil
}

func (u *Policy) UpdateUser(ctx context.Context, input UpdateUserInput) (UpdateUserOutput, error) {
	// Check user's age
	if input.Age < 18 {
		return UpdateUserOutput{}, errors.New("Пользователь должен быть не младше 18 лет")
	}

	// Validate password
	if err := u.validatePassword(input.Password); err != nil {
		return UpdateUserOutput{}, err
	}

	updateUser := model.NewUpdateUser(
		input.FirstName,
		input.LastName,
		input.FullName,
		input.Age,
		input.IsMarried,
		input.Password,
		input.Order,
		input.UpdatedAt,
	)

	user, err := u.userService.UpdateUser(ctx, updateUser)
	if err != nil {
		return UpdateUserOutput{}, errors.Wrap(err, "Error when updating user")
	}

	return UpdateUserOutput{
		User: user,
	}, nil
}

func (u *Policy) DeleteUser(ctx context.Context, input DeleteUserInput) error {
	err := u.userService.DeleteUser(ctx, input.ID)
	if err != nil {
		return errors.Wrap(err, "Error when deleting user")
	}

	return nil
}

func (u *Policy) CreateOrder(ctx context.Context, input CreateOrderInput) (CreateOrderOutput, error) {

	if !u.userService.AreProductsAvailable(ctx, input.ProductID, input.Products) {
		return CreateOrderOutput{}, errors.New("Products not available in stock")
	}

	createOrder := model.NewCreateOrder(
		input.ID,
		input.UserID,
		input.ProductID,
		input.Products,
		input.TimeStamp,
	)

	createdOrder, err := u.userService.CreateOrder(ctx, createOrder)
	if err != nil {
		return CreateOrderOutput{}, errors.Wrap(err, "Error when creating an order")
	}

	order := createdOrder.ToOrder()

	return CreateOrderOutput{
		Order: order,
	}, nil
}
