package controller

import (
	"cvwoapi/helper"
	"cvwoapi/model"
	"net/http"

	"github.com/gin-gonic/gin"
	// "fmt"
)

func SignUp(context *gin.Context) {
    var input model.AuthenticationInput

    if err := context.ShouldBindJSON(&input); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := model.User{
        Username: input.Username,
    }

    savedUser, err := user.Save()

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusCreated, gin.H{"user": savedUser, "ID": savedUser.ID, "username": user.Username})
}

func Login(context *gin.Context) {
    var input model.AuthenticationInput

    if err := context.ShouldBindJSON(&input); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := model.FindUserByUsername(input.Username)

    if (err != nil) {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	if (user.ID == 0) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User"})
        return
	}
	

    jwt, err := helper.GenerateJWT(user)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusOK, gin.H{"jwt": jwt, "ID": user.ID, "username": user.Username})
} 

