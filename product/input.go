package product

import (
	"tokoBelanja/category"
	"tokoBelanja/user"
)

type ProductInput struct {
	Title      string `json:"title" binding:"required"`
	Price      int    `json:"price" binding:"required"`
	Stock      int    `json:"stock" binding:"required"`
	CategoryID int    `json:"category_id" binding:"required"`
}

// type LoginInput struct {
// 	Email    string `json:"email" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }

type GetinputID struct {
	ID int `uri:"id" binding:"required"`
}

type UpdatedProduct struct {
	Title      string `json:"title" binding:"required"`
	Price      int    `json:"price" binding:"required"`
	Stock      int    `json:"stock" binding:"required"`
	CategoryID int    `json:"category_id" binding:"required"`
	User       user.User
	Category   category.Categorys
}
