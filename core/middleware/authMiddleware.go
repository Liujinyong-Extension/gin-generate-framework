package middleware

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO 实现token验证
		c.Next()
	}
}
