package help

import (
	"belajar-api/domain"
	"errors"

	"github.com/gin-gonic/gin"
)

func GetLoginUser(c *gin.Context) (domain.Users, error) {
	user, exists := c.Get("user")
	if !exists {
		return domain.Users{}, errors.New("user not exist")
	}


	return user.(domain.Users), nil
}
