package user

import (
	"github.com/Amore14rn/888Starz/internal/domain/user/model"
	"time"
)

type CreateUserInput struct {
	ID        string
	FirstName string
	LastName  string
	FullName  string
	Age       uint32
	IsMarried bool
	Password  string
	CreatedAt time.Time
	Order     model.Order
}

func NewCreateUserInput(firstName string, lastName string, fullname string, age uint32, isMarried bool, password string, order model.Order) CreateUserInput {
	return CreateUserInput{
		FirstName: firstName,
		LastName:  lastName,
		FullName:  fullname,
		Age:       age,
		IsMarried: isMarried,
		Password:  password,
		Order:     order,
	}
}

type CreateUserOutput struct {
	User model.User
}

type GetUserInput struct {
	ID string
}

func NewGetUserInput(id string) GetUserInput {
	return GetUserInput{
		ID: id,
	}
}

type GetUserOutput struct {
	User model.User
}

type GetUserByNameInput struct {
	FirstName string
}

func NewGetUsersInput(firstName string) GetUserByNameInput {
	return GetUserByNameInput{
		FirstName: firstName,
	}
}

type GetUserByNameOutput struct {
	User []model.User
}

type UpdateUserInput struct {
	ID        string
	FirstName string
	LastName  string
	FullName  string
	Age       uint32
	IsMarried bool
	Password  string
	Order     model.Order
	UpdatedAt time.Time
}

func NewUpdateUserInput(
	id string,
	firstName string,
	lastName string,
	fullname string,
	age uint32,
	isMarried bool,
	password string,
	order model.Order,
	updatedAt time.Time) UpdateUserInput {
	return UpdateUserInput{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		FullName:  fullname,
		Age:       age,
		IsMarried: isMarried,
		Password:  password,
		Order:     order,
		UpdatedAt: updatedAt,
	}
}

type UpdateUserOutput struct {
	User model.User
}

type DeleteUserInput struct {
	ID string
}

func NewDeleteUserInput(id string) DeleteUserInput {
	return DeleteUserInput{
		ID: id,
	}
}

type DeleteUserOutput struct {
}

type CreateOrderInput struct {
	ID        string
	UserID    string
	ProductID string
	Products  []model.OrderProduct
	TimeStamp time.Time
}

func NewCreateOrderInput(id string, userID string, productID string, products []model.OrderProduct, timestamp time.Time) CreateOrderInput {
	return CreateOrderInput{
		ID:        id,
		UserID:    userID,
		ProductID: productID,
		Products:  products,
		TimeStamp: timestamp,
	}
}

type CreateOrderOutput struct {
	Order model.Order
}
