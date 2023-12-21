package model

import (
	"time"
)

type PostRegisterRequest struct {
	Id        uint      `json:"id"`
	Fullname  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type PostCategoriesRequest struct {
	Id          uint      `json:"id"`
	Type        string    `json:"type"`
	SoldProduct int       `json:"sold_product_amount"`
	CreatedAt   time.Time `json:"created_at"`
}

type ProductInCategoryRequest struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Price     int       `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetCategoriesRequest struct {
	Id          uint                       `json:"id"`
	Type        string                     `json:"type"`
	SoldProduct int                        `json:"sold_product_amount"`
	CreatedAt   time.Time                  `json:"created_at"`
	UpdatedAt   time.Time                  `json:"updated_at"`
	Products    []ProductInCategoryRequest `json:"Products"`
}

type PatchCategoriesRequest struct {
	Id          uint      `json:"id"`
	Type        string    `json:"type"`
	SoldProduct int       `json:"sold_product_amount"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PostProductsRequest struct {
	Id         uint      `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type GetProductsRequest struct {
	Id         uint      `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type PutProductsRequest struct {
	Id         uint      `json:"id"`
	Title      string    `json:"title"`
	Price      string    `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID int       `json:"CategoryId"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type PostTransactionsRequest struct {
	TotalPrice   int    `json:"total_price"`
	Quantity     int    `json:"quantity"`
	ProductTitle string `json:"product_title"`
}

type ProductInTransactionRequest struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Price     int       `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetTransactionsRequestResponse struct {
	Id         uint `json:"id"`
	ProductID  int  `json:"product_id"`
	UserID     int  `json:"user_id"`
	Quantity   int  `json:"quantity"`
	TotalPrice int  `json:"total_price"`
	Product    ProductInTransactionRequest
}

type UserInTransactionRequest struct {
	Id        uint      `json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetTransactionsAdminRequestResponse struct {
	Id         uint `json:"id"`
	ProductID  int  `json:"product_id"`
	Quantity   int  `json:"quantity"`
	UserID     int  `json:"user_id"`
	TotalPrice int  `json:"total_price"`
	Product    ProductInTransactionRequest
	User       UserInTransactionRequest
}
