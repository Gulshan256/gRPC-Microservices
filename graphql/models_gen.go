// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package main

import (
	"time"
)

type Account struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Order []*Order `json:"order"`
}

type AccountInput struct {
	Name string `json:"name"`
}

type Mutation struct {
}

type OrderInput struct {
	AccountID string                 `json:"accountId"`
	Products  []*OrderedProductInput `json:"products"`
}

type OrderedProduct struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

type OrderedProductInput struct {
	ID       string `json:"id"`
	Quantity int    `json:"quantity"`
}

type PaginationInput struct {
	Skip int `json:"skip"`
	Take int `json:"take"`
}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type ProductInput struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type Query struct {
}

type Order struct {
	ID         string            `json:"id"`
	Products   []*OrderedProduct `json:"products"`
	TotalPrice float64           `json:"totalPrice"`
	CreatedAt  time.Time         `json:"createdAt"`
}
