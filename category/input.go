package category

import "tokoBelanja/user"

type CategoryInput struct {
	Type string `json:"type" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetinputID struct {
	ID int `uri:"id" binding:"required"`
}

type UpdatedCategory struct {
	Type string `json:"type" binding:"required"`
	User user.User
}
