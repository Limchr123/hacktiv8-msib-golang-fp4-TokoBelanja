package user

type RegisterUserInput struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     int    `json:"role" binding:"required"`
	Balance  int    `json:"balance" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetinputID struct {
	ID int `uri:"id" binding:"required"`
}

type UpdatedUser struct {
	Balance int `json:"balance" binding:"required"`
	User    User
}
