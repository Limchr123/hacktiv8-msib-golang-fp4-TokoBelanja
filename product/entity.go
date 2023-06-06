package product

import (
	"time"
	"tokoBelanja/category"
)

type Products struct {
	ID         int
	Title      string
	Price      int
	Stock      int
	CategoryID int
	Category   category.Categorys `gorm:"foreignKey:CategoryID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
