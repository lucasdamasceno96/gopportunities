package router

import "github.com/gin-gonic/gin"

func Initialize(){
		r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Working! Hey this is a message from backend",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
