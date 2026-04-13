package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Could not bind JSON": err.Error()})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Could not save user": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Could not bind JSON": err.Error()})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
