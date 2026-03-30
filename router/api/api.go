package api

import "github.com/gin-gonic/gin"

func TestApi(r *gin.Engine) *gin.Engine {
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	return r
}
