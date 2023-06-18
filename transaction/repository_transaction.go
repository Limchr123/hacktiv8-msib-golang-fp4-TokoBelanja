package transaction

import (
	"gorm.io/gorm"
)

type RepositoryTransaction interface {
	FindAll() ([]TransactionHistory, error)
	Save(transaction TransactionHistory) (TransactionHistory, error)
	FindById(ID int) (TransactionHistory, error)
	Update(transaction TransactionHistory) (TransactionHistory, error)
	Delete(transaction TransactionHistory) (TransactionHistory, error)
	FindByUserId(userID int) ([]TransactionHistory, error)
}

type repositoryTransaction struct {
	db *gorm.DB
}

func NewRepositoryTransaction(db *gorm.DB) *repositoryTransaction {
	return &repositoryTransaction{db}
}

func (r *repositoryTransaction) FindAll() ([]TransactionHistory, error) {
	var transaction []TransactionHistory

	err := r.db.Preload("Product").Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repositoryTransaction) Save(transaction TransactionHistory) (TransactionHistory, error) {
	err := r.db.Create(&transaction).Error

	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repositoryTransaction) FindByUserId(ID int) ([]TransactionHistory, error) {
	var transaction []TransactionHistory
	err := r.db.Where("id = ?", ID).Find(&transaction).Error

	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repositoryTransaction) FindById(ID int) (TransactionHistory, error) {
	var transaction TransactionHistory

	err := r.db.Where("id = ?", ID).Find(&transaction).Error

	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repositoryTransaction) Update(transaction TransactionHistory) (TransactionHistory, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil

}

func (r *repositoryTransaction) Delete(transaction TransactionHistory) (TransactionHistory, error) {
	err := r.db.Delete(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
