package product

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Products, error)
	Save(product Products) (Products, error)
	FindById(ID int) (Products, error)
	Update(product Products) (Products, error)
	Delete(product Products) (Products, error)
	FindByUserId(userID int) ([]Products, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Products, error) {
	var product []Products

	err := r.db.Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) Save(product Products) (Products, error) {
	err := r.db.Create(&product).Error

	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) FindByUserId(ID int) ([]Products, error) {
	var product []Products
	err := r.db.Where("id = ?", ID).Find(&product).Error

	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) FindById(ID int) (Products, error) {
	var product Products

	err := r.db.Where("id = ?", ID).Find(&product).Error

	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) Update(product Products) (Products, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil

}

func (r *repository) Delete(product Products) (Products, error) {
	err := r.db.Delete(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}
