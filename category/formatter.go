package category

import "time"

// import "time"

type CategoryFormatter struct {
	ID                int       `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount int       `json:"sold_product_amount"`
	CreatedAt         time.Time `json:"created_at"`
}

func FormatterCategory(category Categorys) CategoryFormatter {
	formatterCategory := CategoryFormatter{
		ID:                category.ID,
		Type:              category.Type,
		SoldProductAmount: category.SoldProductAmount,
		CreatedAt:         category.CreatedAt,
	}
	return formatterCategory

}

type UserFormatter struct {
	Token string `json:"token"`
}

func FormatterUser(Token string) UserFormatter {
	formatterLogin := UserFormatter{
		Token: Token,
	}
	return formatterLogin
}

type UpdatedCategoryFormatter struct {
	ID                int       `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount int       `json:"sold_product_amount"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func FormatterCategoryUpdated(category Categorys) UpdatedCategoryFormatter {
	formatterCategory := UpdatedCategoryFormatter{
		ID:                category.ID,
		Type:              category.Type,
		SoldProductAmount: category.SoldProductAmount,
		UpdatedAt:         category.UpdatedAt,
	}
	return formatterCategory

}

// type DeletedUserFormatter struct {
// 	Message string `json:"message"`
// }

// func FormatterDeletedUser(user string) DeletedUserFormatter {
// 	formatterDeletedUser := DeletedUserFormatter{
// 		Message: user,
// 	}
// 	return formatterDeletedUser
// }
