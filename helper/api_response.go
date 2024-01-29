package help

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FailedResponse(c *gin.Context, code int, message string, err error) {
	c.JSON(code, gin.H{
		"Message": message,
		"error":   err.Error(),
	})
}

func SuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, gin.H{

		"message": message,
		"data":    data,
	})
}

func UnathorizedResponse(c *gin.Context, message string, err error){
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": message,
		"err": err.Error(),
	})
}

type ErrorObject struct {
	Code    int
	Message string
	Err     error
}
