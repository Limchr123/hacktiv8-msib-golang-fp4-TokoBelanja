package category

import "errors"

type Service interface {
	CreateCategory(input CategoryInput) (Categorys, error)
	// Login(input LoginInput) (User, error)
	// GetUserByid(ID int) (User, error)
	DeleteCategory(ID int) (Categorys, error)
	UpdatedCategory(getUpdatedInput GetinputID, inputUser UpdatedCategory) (Categorys, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateCategory(input CategoryInput) (Categorys, error) {
	category := Categorys{}

	category.Type = input.Type

	newCategory, err := s.repository.Save(category)
	if err != nil {
		return newCategory, err
	}
	return newCategory, nil
}

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

func (s *service) UpdatedCategory(getUpdatedInput GetinputID, inputCategory UpdatedCategory) (Categorys, error) {

	category, err := s.repository.FindById(getUpdatedInput.ID)
	if err != nil {
		return category, err
	}

	if category.ID != inputCategory.User.ID {
		return category, errors.New("not an owner the account")
	}

	category.Type = inputCategory.Type

	categoryUpdated, err := s.repository.Update(category)
	if err != nil {
		return categoryUpdated, err
	}

	return categoryUpdated, nil

}

func (s *service) DeleteCategory(categoryID int) (Categorys, error) {
	category, err := s.repository.FindById(categoryID)
	if err != nil {
		return category, err
	}
	categoryDel, err := s.repository.Delete(category)

	if err != nil {
		return categoryDel, err
	}
	return categoryDel, nil
}

// func (s *service) GetUserByid(ID int) (User, error) {
// 	user, err := s.repository.FindById(ID)

// 	if err != nil {
// 		return user, err
// 	}

// 	if user.ID == 0 {
// 		return user, errors.New("User Not Found With That ID")
// 	}

// 	return user, nil

// }
