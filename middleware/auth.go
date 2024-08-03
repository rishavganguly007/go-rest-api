package middleware

import (
	"net/http"

	"example.com/go-rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context){
	token := context.Request.Header.Get("Authorization")
	 
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not auhtorised"})
	}

	userId, err := utils.VerifyToken(token) 
	
	if err != nil{
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not auhtorised"})
	}
	context.Set("userId", userId)
	context.Next()
}