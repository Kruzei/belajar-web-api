package help

import "github.com/gin-gonic/gin"

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

type ErrorObject struct {
	Code    int
	Message string
	Err     error
}
