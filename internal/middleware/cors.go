package middleware

import "github.com/gin-gonic/gin"

//cara cors manual
func CORSMiddleware(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Credintial", "*")
	ctx.Header("Acces-Control-Allow-Headers", "Authorization, Origin, Content-Type")
	ctx.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}

	ctx.Next()
}
