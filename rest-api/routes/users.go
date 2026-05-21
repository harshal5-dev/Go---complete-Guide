package routes

import (
	"example/com/rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse",
		})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not save user!",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User created!",
	})
}
