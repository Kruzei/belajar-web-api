package handler

import (
	"belajar-api/app/user/usecase"
	"belajar-api/domain"
	"belajar-api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecase.IUserUsecase
}

func NewUserHandler(userUsecase usecase.IUserUsecase) *UserHandler {
	return &UserHandler{userUsecase}
}

func (h *UserHandler) FindAllUsers(c *gin.Context) {
	users, err := h.userUsecase.FindAllUsers()
	if err != nil {
		help.FailedResponse(c, http.StatusBadRequest, "Failed get all users", err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "Success get all users", users)
}

func (h *UserHandler) FindUser(c *gin.Context){
	nameString := c.Param("name")
	user, err := h.userUsecase.FindUser(nameString)
	if err != nil {
		help.FailedResponse(c, http.StatusBadRequest, "Failed to find user", err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "Success find user", user)
}

func (h *UserHandler) CreateUser(c *gin.Context){
	var userRequest domain.UsersRequests

	err := c.ShouldBindJSON(&userRequest)
	if err != nil{
		help.FailedResponse(c, http.StatusBadRequest, "Failed bind book", err)
		return
	}

	user, err := h.userUsecase.CreateUser(userRequest)
	if err != nil {
		help.FailedResponse(c, http.StatusBadRequest, "Failed to create user", err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "Success create user", user)
}

func (h *UserHandler) UpdateUser(c *gin.Context){
	var userRequest domain.UsersRequests

	err := c.ShouldBindJSON(&userRequest)
	if err != nil{
		help.FailedResponse(c, http.StatusBadRequest, "Failed bind book", err)
		return
	}

	nameString := c.Param("name")
	user, err := h.userUsecase.UpdateUser(nameString, userRequest)

	if err != nil {
		help.FailedResponse(c, http.StatusBadRequest, "Failed to update user", err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "Success update user", user)
}

func (h *UserHandler) DeleteUser(c *gin.Context){
	nameString := c.Param("name")
	userDelete, err := h.userUsecase.DeleteUser(nameString)

	if err != nil {
		help.FailedResponse(c, http.StatusBadRequest, "Failed to delete user", err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "Success delete user", userDelete)
}