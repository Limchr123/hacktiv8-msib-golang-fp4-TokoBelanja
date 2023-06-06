package transactionhistory

type Service interface {
	CreateTransaction(input TransactionInput) (TransactionHistory, error)
	// // Login(input LoginInput) (User, error)
	// GetProducts(ID int) ([]Products, error)
	// DeleteProduct(ID int) (Products, error)
	// UpdatedProduct(getUpdatedInput GetinputID, inputProduct UpdatedProduct) (Products, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateTransaction(input TransactionInput) (TransactionHistory, error) {
	transaction := TransactionHistory{}

	transaction.ID = transaction.ID
	transaction.Quantity = input.Quantity

	newProduct, err := s.repository.Save(product)
	if err != nil {
		return newProduct, err
	}
	return newProduct, nil
}

// func (s *service) GetProducts(ID int) ([]Products, error) {
// 	if ID != 0 {
// 		product, err := s.repository.FindByUserId(ID)
// 		if err != nil {
// 			return product, err
// 		}
// 		return product, nil
// 	}

// 	product, err := s.repository.FindAll()
// 	if err != nil {
// 		return product, err
// 	}
// 	return product, nil
// }

// func (s *service) Login(input LoginInput) (User, error) {
// 	email := input.Email
// 	password := input.Password

// 	user, err := s.repository.FindByEmail(email)
// 	if err != nil {
// 		return user, err
// 	}
// 	if user.ID == 0 {
// 		return user, errors.New("User not found that email")
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
// 	if err != nil {
// 		return user, err
// 	}
// 	return user, nil

// }

// func (s *service) UpdatedProduct(getUpdatedInput GetinputID, inputProduct UpdatedProduct) (Products, error) {

// 	cek := inputProduct.CategoryID
// 	cekCategory, err := s.repository.FindById(cek)
// 	if cekCategory.ID != cekCategory.Category.ID {
// 		return cekCategory, err
// 	}

// if product.CategoryID != inputProduct.Category.ID {
// 	return cekCategory, errors.New("not an owner the account")
// }

// product, err := s.repository.FindById(getUpdatedInput.ID)
// if err != nil {
// 	return product, err
// }

// if product.ID != inputProduct.User.ID {
// 	return product, errors.New("not an owner the account")
// }

// 	product.Title = inputProduct.Title
// 	product.Price = inputProduct.Price
// 	product.Stock = inputProduct.Stock
// 	product.CategoryID = inputProduct.CategoryID

// 	productUpdated, err := s.repository.Update(product)
// 	if err != nil {
// 		return productUpdated, err
// 	}

// 	return productUpdated, nil

// }

// func (s *service) DeleteProduct(ID int) (Products, error) {
// 	product, err := s.repository.FindById(ID)
// 	if err != nil {
// 		return product, err
// 	}
// 	productDel, err := s.repository.Delete(product)

// 	if err != nil {
// 		return productDel, err
// 	}
// 	return productDel, nil
// }

// func (s *service) GetUserByid(ID int) (Products, error) {
// 	user, err := s.repository.FindById(ID)

// 	if err != nil {
// 		return user, err
// 	}

// 	if user.ID == 0 {
// 		return user, errors.New("User Not Found With That ID")
// 	}

// 	return user, nil

// }
