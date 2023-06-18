package transaction

import (
	"time"
)

// import "time"

// type TransactionFormatter struct {
// 	Title      string `json:"title"`
// 	Price      int    `json:"price"`
// 	Stock      int    `json:"stock"`
// 	CategoryID int    `json:"category_id"`
// }

// func FormatterTransaction(product Products) ProductFormatter {
// 	formatterProduct := ProductFormatter{
// 		Title:      product.Title,
// 		Price:      product.Price,
// 		Stock:      product.Stock,
// 		CategoryID: product.CategoryID,
// 	}
// 	return formatterProduct

// }

type ProductsFormatter struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserFormatter struct {
	ID        int `json:"id"`
	FullName  string  `json:"full_name" `
	Email string `json:"email"`
	Balance   int `json:"balance" `
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TransactionGetFormatter struct {
	ID         int       `json:"id"`
	ProductID      int    `json:"product_id"`
	UserID      int       `json:"user_id"`
	Quantity      int       `json:"quantity"`
	TotalPrice int       `json:"total_price"`
	Product ProductsFormatter `json:"Product"`
	User UserFormatter `json:"User"`
	CreatedAt  time.Time `json:"created_at"`
}

func FormatterGet(transaction TransactionHistory) TransactionGetFormatter {
	formatterGet := TransactionGetFormatter{}
	formatterGet.ID = transaction.ID
	formatterGet.ProductID = transaction.ProductID
	formatterGet.UserID = transaction.UserID
	formatterGet.Quantity = transaction.Quantity
	formatterGet.TotalPrice = transaction.TotalPrice
	formatterGet.CreatedAt = transaction.CreatedAt

	newProduct := transaction.Product

	products := ProductsFormatter{}
	products.ID = newProduct.ID
	products.Title = newProduct.Title
	products.Price = newProduct.Price
	products.Stock = newProduct.Stock
	products.CategoryID = newProduct.CategoryID
	products.CreatedAt = newProduct.CreatedAt
	products.UpdatedAt = newProduct.UpdatedAt

	formatterGet.Product = products

	newUserdua := transaction.User

	users := UserFormatter{}

	users.ID = newUserdua.ID
	users.Email = newUserdua.Email
	users.FullName = newUserdua.Email
	users.Balance = newUserdua.Balance
	users.CreatedAt = newUserdua.CreatedAt
	users.UpdatedAt = newUserdua.UpdatedAt

	formatterGet.User = users

	return formatterGet
}

func FormatterGetCampaign(products []TransactionHistory) []TransactionGetFormatter {
	productGetFormatter := []TransactionGetFormatter{}

	for _, product := range products {
		productFormatter := FormatterGet(product)
		productGetFormatter = append(productGetFormatter, productFormatter)
	}

	return productGetFormatter
}

// type ProductUpdateFormatter struct {
// 	ID         int       `json:"id"`
// 	Title      string    `json:"title"`
// 	Price      int       `json:"price"`
// 	Stock      int       `json:"stock"`
// 	CategoryID int       `json:"category_id"`
// 	CreatedAt  time.Time `json:"created_at"`
// 	UpdatedAt  time.Time `json:"updated_at"`
// }

// func FormatterUpdate(product Products) ProductUpdateFormatter {
// 	formatterUpdate := ProductUpdateFormatter{}
// 	formatterUpdate.ID = product.ID
// 	formatterUpdate.Title = product.Title
// 	formatterUpdate.Price = product.Price
// 	formatterUpdate.Stock = product.Stock
// 	formatterUpdate.CategoryID = product.CategoryID
// 	formatterUpdate.CreatedAt = product.CreatedAt
// 	formatterUpdate.UpdatedAt = product.UpdatedAt

// 	return formatterUpdate
// }
