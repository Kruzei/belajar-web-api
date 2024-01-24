package usecase

import (
	"belajar-api/app/user/repository"
	"belajar-api/domain"
	help "belajar-api/helper"
	"errors"
	"net/http"
)

type IUserUsecase interface {
	FindAllUsers() ([]domain.Users, any)
	FindUser(id int) (domain.Users, any)
	CreateUser(user domain.UsersRequests) (domain.Users, any)
	UpdateUser(id int, user domain.UsersRequests) (domain.Users, any)
	DeleteUser(id int) (domain.Users, any)
}

type UserUsecase struct {
	userRepository repository.IUserRepository
}

func NewUserUsecase(repository repository.IUserRepository) *UserUsecase {
	return &UserUsecase{repository}
}

func (u *UserUsecase) FindAllUsers() ([]domain.Users, any) {
	var users []domain.Users
	err := u.userRepository.FindAllUsers(&users)
	if err != nil {
		return []domain.Users{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "failed to get user by id",
			Err:     err,
		}
	}

	return users, err
}

func (u *UserUsecase) FindUser(id int) (domain.Users, any) {
	var user domain.Users
	err := u.userRepository.FindUser(&user, id)
	if err != nil {
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "failed to get user by id",
			Err:     errors.New("id not found"),
		}
	}

	return user, err
}

func (u *UserUsecase) CreateUser(userRequest domain.UsersRequests) (domain.Users, any) {
	isUserExist := u.userRepository.FindUserByCondition(&domain.Users{}, userRequest.Name)
	if isUserExist == nil{
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusConflict,
			Message: "failed to create user",
			Err:     errors.New("name already used"),
		}
	}
	age, _ := userRequest.Age.Int64()

	user := domain.Users{
		Name:     userRequest.Name,
		Age:      int(age),
		Gender:   userRequest.Gender,
		Password: userRequest.Password,
	}

	err := u.userRepository.CreateUser(&user)
	if err != nil {
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to create user",
			Err:     err,
		}
	}
	return user, err
}

func (u *UserUsecase) UpdateUser(id int, userRequest domain.UsersRequests) (domain.Users, any) {
	var user domain.Users
	err := u.userRepository.FindUser(&user, id)
	if err != nil {
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "failed to get user by id",
			Err:     errors.New("id not found"),
		}
	}

	user.Password = userRequest.Password

	err = u.userRepository.UpdateUser(&user)
	if err != nil {
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to update user",
			Err:     err,
		}
	}

	return user, err
}

func (u *UserUsecase) DeleteUser(id int) (domain.Users, any) {
	var user domain.Users
	err := u.userRepository.FindUser(&user, id)
	if err != nil {
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "failed to get user by id",
			Err:     errors.New("id not found"),
		}
	}

	err = u.userRepository.DeleteUser(&user)
	if err != nil {
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to delete user",
			Err:     err,
		}
	}

	return user, err
}
