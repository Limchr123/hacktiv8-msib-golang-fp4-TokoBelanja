package category

import (
	"time"
)

type Categorys struct {
	ID                int
	Type              string
	SoldProductAmount int
	CreatedAt         time.Time
	UpdatedAt         time.Time
	// Product  Products
}
