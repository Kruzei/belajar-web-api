package repository

import (
	"belajar-api/domain"

	"gorm.io/gorm"
)

type IUserRepository interface {
	FindAllUsers(users *[]domain.Users) (error)
	FindUser(user *domain.Users, id int) (error)
	SignUp(user *domain.Users) (error)
	UpdateUser(user *domain.Users) (error)
	DeleteUser(user *domain.Users) (error)
	FindUserByCondition(user *domain.Users, name string) (error)
	FindByEmail(user *domain.Users, email string)(error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindAllUsers(users *[]domain.Users) (error) {
	err := r.db.Find(users).Error
	return err
}

func (r *UserRepository) FindUser(user *domain.Users, id int) (error) {
	err := r.db.Where("id = ?", id).First(user).Error
	return err
}

func (r *UserRepository) SignUp(user *domain.Users)(error){
	err := r.db.Create(user).Error
	return err
}

func (r *UserRepository) UpdateUser(user *domain.Users)(error){
	err := r.db.Save(user).Error
	return err
}

func (r *UserRepository) DeleteUser(user *domain.Users)(error){
	err := r.db.Delete(user).Error
	return err
}

func (r *UserRepository) FindUserByCondition(user *domain.Users, email string)(error){
	err := r.db.Where("email = ?", email).First(&user).Error
	return err
}

func(r *UserRepository) FindByEmail(user *domain.Users, email string)(error){
	err := r.db.Where("email = ?", email).First(user).Error
	return err
}
