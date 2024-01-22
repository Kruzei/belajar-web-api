package usecase

import (
	"belajar-api/app/user/repository"
	"belajar-api/domain"
)

type IUserUsecase interface {
	FindAllUsers() ([]domain.Users, error)
	FindUser(name string) (domain.Users, error)
	CreateUser(user domain.UsersRequests) (domain.Users, error)
	UpdateUser(name string, user domain.UsersRequests) (domain.Users, error)
	DeleteUser(name string) (domain.Users, error)
}

type UserUsecase struct{
	userRepository repository.IUserRepository
}

func NewUserUsecase(repository repository.IUserRepository) *UserUsecase{
	return &UserUsecase{repository}
}

func (u *UserUsecase) FindAllUsers()([]domain.Users, error){
	return u.userRepository.FindAllUsers()
}

func (u *UserUsecase) FindUser(name string)(domain.Users, error){
	return u.userRepository.FindUser(name)
}

func (u *UserUsecase) CreateUser(userRequest domain.UsersRequests)(domain.Users, error){
	age, _ := userRequest.Age.Int64()

	user := domain.Users{
		Name: userRequest.Name,
		Age: int(age),
		Gender: userRequest.Gender,
		Password: userRequest.Password,
	}

	newUser, err := u.userRepository.CreateUser(user)
	return newUser, err
}

func (u *UserUsecase) UpdateUser(name string, userRequest domain.UsersRequests)(domain.Users, error){
	user, _ := u.userRepository.FindUser(name)

	user.Password = userRequest.Password

	updatedUser, err := u.userRepository.UpdateUser(user)
	return updatedUser, err
}

func (u *UserUsecase) DeleteUser(name string)(domain.Users, error){
	user, _ := u.userRepository.FindUser(name)
	deletedUser, err := u.userRepository.DeleteUser(user)
	return deletedUser, err
}



