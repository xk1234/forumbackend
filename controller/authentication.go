package controller

import (
	"cvwoapi/helper"
    "cvwoapi/model"
    "github.com/gin-gonic/gin"
    "net/http"
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

    context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func Login(context *gin.Context) {
    var input model.AuthenticationInput

    if err := context.ShouldBindJSON(&input); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := model.FindUserByUsername(input.Username)

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    jwt, err := helper.GenerateJWT(user)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusOK, gin.H{"jwt": jwt})
} 

