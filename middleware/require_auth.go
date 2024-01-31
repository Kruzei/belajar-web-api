package middleware

import (
	"belajar-api/domain"
	help "belajar-api/helper"
	"belajar-api/infrastructure/database"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {
	bearerToken := c.Request.Header.Get("Authorization")
	fmt.Println(bearerToken)
	token := strings.Split(bearerToken, " ")[1]

	if token == "" {
		help.UnathorizedResponse(c, "failed to validate token", nil)
		return
	}

	userId, expTime, err := help.ValidateToken(token)

	if err != nil {
		help.UnathorizedResponse(c, "failed to validate token", err)
		return
	}

	var user domain.Users
	err = database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		help.UnathorizedResponse(c, "failed to get user id", err)
		return
	}

	if float64(time.Now().Unix()) > expTime {
		help.UnathorizedResponse(c, "token expired", err)
	}

	c.Set("user", user)
	c.Next()
}

func OnlyAdmin(c *gin.Context){
	user, ok := c.Get("user")
	if !ok {
		help.FailedResponse(c, http.StatusNotFound, "user not found", errors.New(""))
		return
	}

	if user.(domain.Users).Role != "ADMIN"{
		help.UnathorizedResponse(c, "admin only", errors.New("acces denied"))
		return
	}

	c.Next()
}
