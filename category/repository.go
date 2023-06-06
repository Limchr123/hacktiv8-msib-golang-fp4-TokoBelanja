package category

import "gorm.io/gorm"

type Repository interface {
	//create User
	Save(category Categorys) (Categorys, error)
	FindById(ID int) (Categorys, error)
	// FindByEmail(email string) (Category, error)
	Update(category Categorys) (Categorys, error)
	Delete(category Categorys) (Categorys, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(category Categorys) (Categorys, error) {
	err := r.db.Create(&category).Error

	if err != nil {
		return category, err
	}
	return category, nil
}

// func (r *repository) FindByEmail(email string) (User, error) {
// 	var user User
// 	err := r.db.Where("email = ?", email).Find(&user).Error

// 	if err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }

func (r *repository) FindById(ID int) (Categorys, error) {
	var category Categorys

	err := r.db.Where("id = ?", ID).Find(&category).Error

	if err != nil {
		return category, err
	}
	return category, nil
}

func (r *repository) Update(category Categorys) (Categorys, error) {
	err := r.db.Save(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil

}

func (r *repository) Delete(category Categorys) (Categorys, error) {
	err := r.db.Delete(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}
