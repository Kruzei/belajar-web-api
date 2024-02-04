package usecase

import (
	"belajar-api/app/user/repository"
	"belajar-api/domain"
	help "belajar-api/helper"
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	FindAllUsers() ([]domain.Users, any)
	FindUser(id int) (domain.Users, any)
	SignUp(user domain.UsersRequests) (domain.Users, any)
	UpdateUser(id int, user domain.UsersRequests) (domain.Users, any)
	DeleteUser(id int) (domain.Users, any)
	LoginUser(UserLogin domain.UserLogin, email string) (domain.Users, interface{}, any)
	// GetUser(param domain.UserParam) (domain.Users, error)
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

func (u *UserUsecase) SignUp(userRequest domain.UsersRequests) (domain.Users, any) {
	isUserExist := u.userRepository.FindUserByEmail(&domain.Users{}, userRequest.Email)
	if isUserExist == nil {
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusConflict,
			Message: "failed to create user",
			Err:     errors.New("phone number already used"),
		}
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 10)
	if err != nil {
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to hash password",
			Err:     err,
		}
	}
	user := domain.Users{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: string(hashPassword),
		Role:     userRequest.Role,
	}

	err = u.userRepository.SignUp(&user)
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

	user.Name = userRequest.Name
	user.Email = userRequest.Email

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

func (u *UserUsecase) LoginUser(UserLogin domain.UserLogin, email string) (domain.Users, interface{}, any) {
	var user domain.Users
	err := u.userRepository.FindUserByEmail(&user, email)
	if err != nil {
		return domain.Users{}, "", help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "invalid user or password",
			Err:     errors.New("invalid user or password"),
		}
	}
	// user, err := u.userRepository.GetUser(domain.UserParam{Email: email})
	// if err != nil {
	// 	return domain.Users{}, "", help.ErrorObject{
	// 		Code:    http.StatusNotFound,
	// 		Message: "invalid user or password",
	// 		Err:     errors.New("invalid user or password"),
	// 	}
	// }

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserLogin.Password))
	if err != nil {
		return domain.Users{}, "", help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "invalid user or password",
			Err:     errors.New("invalid user or password"),
		}
	}

	tokenString, err := help.GenerateToken(user)
	if err != nil {
		return domain.Users{}, "", help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "Failed to make token",
			Err:     err,
		}
	}

	apiResponse := struct {
		Token string `json:"token"`
	}{
		tokenString,
	}

	return user, apiResponse, err
}

// func (u *UserUsecase) GetUser(param domain.UserParam) (domain.Users, error) {
// 	user, err := u.userRepository.GetUser(param)
// 	if err != nil {
// 		return user, nil
// 	}

// 	return user, nil
// }
