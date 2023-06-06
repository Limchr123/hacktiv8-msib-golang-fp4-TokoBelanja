package transactionhistory

import (
	"time"
	"tokoBelanja/product"
	"tokoBelanja/user"
)

type TransactionHistory struct {
	ID         int
	ProductID  int
	UserID     int
	Quantity   int
	TotalPrice int
	User       user.User        `gorm:"foreignKey:UserID"`
	Product    product.Products `gorm:"foreignKey:ProductID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
