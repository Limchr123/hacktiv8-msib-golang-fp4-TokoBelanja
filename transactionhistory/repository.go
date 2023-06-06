package transactionhistory

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]TransactionHistory, error)
	Save(transaction TransactionHistory) (TransactionHistory, error)
	FindById(ID int) (TransactionHistory, error)
	Update(transaction TransactionHistory) (TransactionHistory, error)
	Delete(transaction TransactionHistory) (TransactionHistory, error)
	FindByUserId(userID int) ([]TransactionHistory, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]TransactionHistory, error) {
	var transaction []TransactionHistory

	err := r.db.Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) Save(transaction TransactionHistory) (TransactionHistory, error) {
	err := r.db.Create(&transaction).Error

	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) FindByUserId(ID int) ([]TransactionHistory, error) {
	var transaction []TransactionHistory
	err := r.db.Where("id = ?", ID).Find(&transaction).Error

	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) FindById(ID int) (TransactionHistory, error) {
	var transaction TransactionHistory

	err := r.db.Where("id = ?", ID).Find(&transaction).Error

	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) Update(transaction TransactionHistory) (TransactionHistory, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil

}

func (r *repository) Delete(transaction TransactionHistory) (TransactionHistory, error) {
	err := r.db.Delete(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
