package transaction

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
	Product    product.Products `gorm:"foreignKey:ProductID;preload"`
	User       user.User        `gorm:"foreignKey:UserID;preload"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
