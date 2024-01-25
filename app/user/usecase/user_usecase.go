package usecase

import (
	"belajar-api/app/user/repository"
	"belajar-api/domain"
	help "belajar-api/helper"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	FindAllUsers() ([]domain.Users, any)
	FindUser(id int) (domain.Users, any)
	SignUp(user domain.UsersRequests) (domain.Users, any)
	UpdateUser(id int, user domain.UsersRequests) (domain.Users, any)
	DeleteUser(id int) (domain.Users, any)
	LoginUser(UserLogin domain.UserLogin, email string) (domain.Users, string, any)
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
	isUserExist := u.userRepository.FindUserByCondition(&domain.Users{}, userRequest.Email)
	if isUserExist == nil {
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusConflict,
			Message: "failed to create user",
			Err:     errors.New("name already used"),
		}
	}
	age, _ := userRequest.Age.Int64()

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 10)
	if err != nil {
		return domain.Users{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to hash password",
			Err:     err,
		}
	}
	user := domain.Users{
		Email:    userRequest.Email,
		Name:     userRequest.Name,
		Age:      int(age),
		Gender:   userRequest.Gender,
		Password: string(hashPassword),
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

func (u *UserUsecase) LoginUser(UserLogin domain.UserLogin, email string) (domain.Users, string, any) {
	var user domain.Users
	err := u.userRepository.FindUserByCondition(&user, email)
	if err != nil {
		return domain.Users{}, "", help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "invalid user or password",
			Err:     err,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserLogin.Password))
	if err != nil {
		return domain.Users{}, "" ,help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "invalid user or password",
			Err:     err,
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Minute * 3).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return domain.Users{}, "" , help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "Failed to make token",
			Err:     err,
		}
	}

	return user, tokenString, err
}
