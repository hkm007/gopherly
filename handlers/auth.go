package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hkm007/gopherly/models"
	"github.com/hkm007/gopherly/utils"
)

func SignUp(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cannot parse request data!"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot save user!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created!"})
}

func Login(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cannot parse request data!"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Cannot authenticate user!"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Cannot authenticate user!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Logged in!", "token": token})
}
