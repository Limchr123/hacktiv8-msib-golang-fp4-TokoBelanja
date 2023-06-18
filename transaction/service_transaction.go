package transaction

import (
	"errors"
	"fmt"
	"tokoBelanja/product"
	"tokoBelanja/user"
)

type ServiceTransaction interface {
	CreateTransaction(input TransactionInput) (TransactionHistory, error)
	// // Login(input LoginInput) (User, error)
	GetTransaction(productID int,userID int) ([]TransactionHistory, error)
	// GetTransactionUser(ID int) ([]TransactionHistory, error)
	// DeleteProduct(ID int) (Products, error)
	// UpdatedProduct(getUpdatedInput GetinputID, inputProduct UpdatedProduct) (Products, error)
}

type serviceTransaction struct {
	repository        RepositoryTransaction
	repositoryProduct product.RepositoryProduct
	repositoryUser    user.RepositoryUser
}

func NewService(repository RepositoryTransaction, repositoryProduct product.RepositoryProduct, repositoryUser user.RepositoryUser) *serviceTransaction {
	return &serviceTransaction{repository, repositoryProduct, repositoryUser}
}


func (s *serviceTransaction) CreateTransaction(input TransactionInput) (TransactionHistory, error) {
	transaction := TransactionHistory{}

	transaction.ProductID = input.ProductID

	cek, err := s.repositoryProduct.FindById(input.ProductID)
	if err != nil {
		print(err)
		return TransactionHistory{}, err
	}
	if cek.Stock < input.Quantity {
		print("stock ga muat")
		return TransactionHistory{}, errors.New("error")
	}

	cekSaldo, err := s.repositoryUser.FindById(input.UserID)
	if err != nil {
		print(err)
		return TransactionHistory{}, err
	}
	if cekSaldo.Balance < (input.Quantity * cek.Price) {
		print("saldo ga cukup")
		fmt.Println(cekSaldo.Balance)
		return TransactionHistory{}, errors.New("error")
	}

	cekSaldo.Balance = cekSaldo.Balance - (input.Quantity * cek.Price)

	_, err = s.repositoryUser.Update(cekSaldo)
	if err != nil {
		print(err)
		return TransactionHistory{}, err
	}

	cek.Stock = cek.Stock - input.Quantity

	_, err = s.repositoryProduct.Update(cek)
	if err != nil {
		print(err)
		return TransactionHistory{}, err
	}

	transaction.Quantity = cek.Stock - input.Quantity

	newProduct, err := s.repository.Save(transaction)
	if err != nil {
		print(err)
		return newProduct, err
	}
	return newProduct, nil
}

// func (s *serviceTransaction) GetTransactionUser(ID int) ([]TransactionHistory, error){
	
// }

func (s *serviceTransaction) GetTransaction(productID int ,userID int) ([]TransactionHistory, error) {
	// if ID != 0 {
	// 	transaction, err := s.repository.FindByUserId(ID)
	// 	if err != nil {
	// 		return transaction, err
	// 	}
	// 	return transaction, nil
	// }

	// transaction, err := s.repository.FindByUserId(productID, userID)
	// if err != nil {
	// 	return transaction, err
	// }
	// return transaction, nil

	transaction, err := s.repository.FindAll()
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

// func (s *serviceTransaction) Login(input LoginInput) (User, error) {
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
