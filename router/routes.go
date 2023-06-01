package router

import (
	"gopportunities/handler"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine){

		handler.InitializeHandler()
		
		v1 := router.Group("/api/v1")
		{
			v1.GET("/opening", handler.ShowOpeningHandler)

			v1.GET("/openings", handler.ShowOpeningsHandler)

			v1.POST("/opening", handler.CreateOpeningsHandler)

			v1.PUT("/opening", handler.UpdateOpeningsHandler)

			v1.DELETE("/opening", handler.DeleteOpeningsHandler)
		}
}
