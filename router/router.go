package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Initialize(){
	// Initialize router 
		r := gin.Default()
		r.Use(cors.Default())

	//Initialize routes 
		InitializeRoutes(r)
		
	// Running the server 
	r.Run() // listen and serve on 0.0.0.0:8080

	
}
