package middleware

import (
	"gin-generate-framework/app/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": controllers.AuthorErrorCode, "message": "未认证", "data": nil})
			return
		}

		jwtKey := []byte(viper.GetString("jwt.secret"))

		claims := &jwt.MapClaims{}
		parsedToken, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !parsedToken.Valid {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": controllers.AuthorErrorCode, "message": "Invalid or expired token", "data": nil})
			return
		}

		// 将完整用户信息注入 context，后续通过 BaseController.GetLoginedUser 取出
		if user, ok := (*claims)["user"]; ok {
			c.Set("user", user)
		}

		c.Next()
	}
}
