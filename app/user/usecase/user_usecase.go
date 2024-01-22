package usecase

import (
	"belajar-api/app/user/repository"
	"belajar-api/domain"
)

type IUserUsecase interface {
	FindAllUsers() ([]domain.Users, error)
	FindUser(id int) (domain.Users, error)
	CreateUser(user domain.UsersRequests) (domain.Users, error)
	UpdateUser(id int, user domain.UsersRequests) (domain.Users, error)
	DeleteUser(id int) (domain.Users, error)
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

func (u *UserUsecase) FindUser(id int)(domain.Users, error){
	return u.userRepository.FindUser(id)
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

func (u *UserUsecase) UpdateUser(id int, userRequest domain.UsersRequests)(domain.Users, error){
	user, _ := u.userRepository.FindUser(id)

	user.Password = userRequest.Password

	updatedUser, err := u.userRepository.UpdateUser(user)
	return updatedUser, err
}

func (u *UserUsecase) DeleteUser(id int)(domain.Users, error){
	user, _ := u.userRepository.FindUser(id)
	deletedUser, err := u.userRepository.DeleteUser(user)
	return deletedUser, err
}



