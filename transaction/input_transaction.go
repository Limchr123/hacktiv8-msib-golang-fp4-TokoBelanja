package transaction

type TransactionInput struct {
	ProductID int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required"`
	UserID    int `json:"user_id"`
}

// type LoginInput struct {
// 	Email    string `json:"email" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }

type GetinputTransactionID struct {
	ID int `uri:"id" binding:"required"`
}

// type UpdatedProduct struct {
// 	Title      string `json:"title" binding:"required"`
// 	Price      int    `json:"price" binding:"required"`
// 	Stock      int    `json:"stock" binding:"required"`
// 	CategoryID int    `json:"category_id" binding:"required"`
// 	User       user.User
// 	Category   category.Categorys
// }
