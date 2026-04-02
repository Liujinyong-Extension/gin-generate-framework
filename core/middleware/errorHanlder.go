package middleware

import (
	"fmt"
	"gin-generate-framework/app/controllers"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		fmt.Println("Panic recovered:", recovered)
		c.JSON(controllers.SuccessCode, gin.H{"code": controllers.ServerErrorCode, "message": recovered})
	})
}
