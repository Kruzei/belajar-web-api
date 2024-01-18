package help

import "github.com/gin-gonic/gin"

func FailedResponse(c *gin.Context, code int, message string, err error){
	c.JSON(code, gin.H{
		"status" : "Error occured",
		"Message": message,
		"error": err.Error(),
	})
}

func SuccessResponse(c *gin.Context, code int, message string, data interface{}){
	c.JSON(code, gin.H{
		"status": "succes",
		"message": message,
		"data": data,
	})
}