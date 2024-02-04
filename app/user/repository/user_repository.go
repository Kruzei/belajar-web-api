package repository

import (
	"belajar-api/domain"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetAllUsers(users *[]domain.Users) (error)
	GetUser(user *domain.Users, id int) (error)
	GetUserByEmail(user *domain.Users, email string) (error)
	SignUp(user *domain.Users) (error)
	UpdateUser(user *domain.Users) (error)
	DeleteUser(user *domain.Users) (error)
	// GetUser(param domain.UserParam) (domain.Users, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetAllUsers(users *[]domain.Users) (error) {
	err := r.db.Find(users).Error
	return err
}

func (r *UserRepository) GetUser(user *domain.Users, id int) (error) {
	err := r.db.Where("id = ?", id).First(user).Error
	return err
}

func (r *UserRepository) GetUserByEmail(user *domain.Users, email string)(error){
	err := r.db.Where("email = ?", email).First(&user).Error
	return err
}

func (r *UserRepository) SignUp(user *domain.Users)(error){
	err := r.db.Create(user).Error
	return err
}

func (r *UserRepository) UpdateUser(user *domain.Users) error {
	err := r.db.Save(user).Error
	return err
}

func (r *UserRepository) DeleteUser(user *domain.Users) error {
	err := r.db.Delete(user).Error
	return err
}

// func (r *UserRepository) GetUser(param domain.UserParam) (domain.Users, error) {
// 	var user domain.Users
// 	err := r.db.Model(&domain.Users{}).Where(&param).First(&user).Error

// 	if err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }

// func (r *UserRepository) CreateUser(inputParam domain.UserInput) (domain.Users, error) {
// 	var user domain.Users
// 	err := r.db.Model(&user).Create(&inputParam).Error
// 	if err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }
