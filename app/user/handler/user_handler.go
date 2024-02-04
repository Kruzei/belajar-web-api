package handler

import (
	"belajar-api/app/user/usecase"
	"belajar-api/domain"
	help "belajar-api/helper"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecase.IUserUsecase
}

func NewUserHandler(userUsecase usecase.IUserUsecase) *UserHandler {
	return &UserHandler{userUsecase}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, errorObject := h.userUsecase.GetAllUsers()
	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, http.StatusBadRequest, "Failed get all users", errorObject.Err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "Success get all users", users)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	user, errorObject := h.userUsecase.GetUser(id)
	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, http.StatusBadRequest, "Failed to get user", errorObject.Err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "Success get user", user)
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var userRequest domain.UsersRequests

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		help.FailedResponse(c, http.StatusBadRequest, "Failed bind user", err)
		return
	}

	user, errorObject := h.userUsecase.SignUp(userRequest)
	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, http.StatusBadRequest, "Failed to create user", errorObject.Err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "Success create user", user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var userRequest domain.UsersRequests

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		help.FailedResponse(c, http.StatusBadRequest, "Failed bind book", err)
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	user, errorObject := h.userUsecase.UpdateUser(id, userRequest)
	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, http.StatusBadRequest, "Failed to update user", errorObject.Err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "Success update user", user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	userDelete, errorObject := h.userUsecase.DeleteUser(id)

	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, http.StatusBadRequest, "Failed to delete user", errorObject.Err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "Success delete user", userDelete)
}

func (h *UserHandler) Login(c *gin.Context) {
	var userRequest domain.UserLogin

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		help.FailedResponse(c, http.StatusBadRequest, "Failed bind book", err)
		return
	}

	user, apiResponse, errorObject := h.userUsecase.LoginUser(userRequest, userRequest.Email)
	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, http.StatusBadRequest, "Failed to login", errorObject.Err)
		return
	}

	fmt.Println(apiResponse)

	help.SuccessResponse(c, http.StatusOK, "Welcome "+user.Name, apiResponse)
}
