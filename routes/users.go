package routes

import (
	"net/http"

	"example.com/go-rest-api/models"
	"example.com/go-rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBind(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse request data"})
		return
	}

	user.Id = 1

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created", "event": user})


}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBind(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate user"})
	}

	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not authenticate user"})
		
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "login successful", "token": token})
	                                                                              
}