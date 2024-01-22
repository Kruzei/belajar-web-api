package repository

import (
	"belajar-api/domain"

	"gorm.io/gorm"
)

type IUserRepository interface {
	FindAllUsers() ([]domain.Users, error)
	FindUser(name string) (domain.Users, error)
	CreateUser(user domain.Users) (domain.Users, error)
	UpdateUser(user domain.Users) (domain.Users, error)
	DeleteUser(user domain.Users) (domain.Users, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindAllUsers() ([]domain.Users, error) {
	var users []domain.Users
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepository) FindUser(name string) (domain.Users, error) {
	var user domain.Users
	err := r.db.Where(name).Find(&user).Error
	return user, err
}

func (r *UserRepository) CreateUser(user domain.Users)(domain.Users, error){
	err := r.db.Create(&user).Error
	return user, err
}

func (r *UserRepository) UpdateUser(user domain.Users)(domain.Users, error){
	err := r.db.Save(&user).Error
	return user, err
}

func (r *UserRepository) DeleteUser(user domain.Users)(domain.Users, error){
	err := r.db.Delete(&user).Error
	return user, err
}
