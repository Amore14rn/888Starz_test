package model

import (
	"time"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	FullName  string
	Age       uint32
	IsMarried bool
	Password  string
	Orders    []Order
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func NewUser(
	ID string,
	firstName string,
	lastName string,
	age uint32,
	isMarried bool,
	password string,
	orders Order,
	createdAt time.Time,
	updatedAt *time.Time,
) User {
	fullName := firstName + " " + lastName
	return User{
		ID:        ID,
		FirstName: firstName,
		LastName:  lastName,
		FullName:  fullName,
		Age:       age,
		IsMarried: isMarried,
		Password:  password,
		Orders:    []Order{orders},
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

type Order struct {
	ID        string
	UserID    string
	ProductID string
	Products  []OrderProduct
	Timestamp time.Time
}

type OrderProduct struct {
	ProductID string
	Quantity  int
	Price     float64
}

func (u *User) AddOrder(order Order) {
	u.Orders = append(u.Orders, order)
}

type CreateUser struct {
	ID        string
	FirstName string
	LastName  string
	FullName  string
	Age       uint32
	IsMarried bool
	Password  string
	Order     Order
	CreatedAt time.Time
}

func NewCreateUser(
	ID string,
	firstName string,
	lastName string,
	age uint32,
	isMarried bool,
	password string,
	order Order,
	createdAt time.Time,
) CreateUser {
	fullName := firstName + " " + lastName
	return CreateUser{
		ID:        ID,
		FirstName: firstName,
		LastName:  lastName,
		FullName:  fullName,
		Age:       age,
		IsMarried: isMarried,
		Order:     order,
		Password:  password,
		CreatedAt: createdAt,
	}
}

type UpdateUser struct {
	ID        string
	FirstName string
	LastName  string
	FullName  string
	Age       uint32
	IsMarried bool
	Password  string
	Order     Order
	UpdatedAt time.Time
}

func NewUpdateUser(
	ID string,
	firstName string,
	lastName string,
	age uint32,
	isMarried bool,
	password string,
	order Order,
	updatedAt time.Time,
) UpdateUser {
	fullName := firstName + " " + lastName
	return UpdateUser{
		ID:        ID,
		FirstName: firstName,
		LastName:  lastName,
		FullName:  fullName,
		Age:       age,
		IsMarried: isMarried,
		Password:  password,
		Order:     order,
		UpdatedAt: updatedAt,
	}
}

type CreateOrder struct {
	ID        string
	UserID    string
	ProductID string
	Products  []OrderProduct
	Timestamp time.Time
	TimeStamp time.Time
}

func (co CreateOrder) ToOrder() Order {
	return Order{
		ID:        co.ID,
		UserID:    co.UserID,
		ProductID: co.ProductID,
		Products:  co.Products,
		Timestamp: co.Timestamp,
	}
}

func NewCreateOrder(
	ID string,
	userID string,
	productID string,
	products []OrderProduct,
	timeStamp time.Time,
) CreateOrder {
	return CreateOrder{
		ID:        ID,
		UserID:    userID,
		ProductID: productID,
		Products:  products,
		TimeStamp: timeStamp,
	}
}
