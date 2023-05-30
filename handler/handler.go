package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
					"msg": "GET it's working",
				})
}

func ShowOpeningsHandler(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
					"msg": "Openings working",
				})
}

func CreateOpeningsHandler(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
					"msg": "Create is working",
				})
}
func UpdateOpeningsHandler(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
					"msg": "Update working",
				})
}

func DeleteOpeningsHandler(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
					"msg": "Deleted succesfully",
				})
}
