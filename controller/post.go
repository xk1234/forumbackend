package controller

import (
    "cvwoapi/helper"
    "cvwoapi/model"
    "github.com/gin-gonic/gin"
    "net/http"
)

func AddPost(context *gin.Context) {
    var input model.Post
    if err := context.ShouldBindJSON(&input); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    user, err := helper.CurrentUser(context)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    input.UserID = user.ID
    savedPost, err := input.Save()
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    context.JSON(http.StatusCreated, gin.H{"data": savedPost})
}

func GetAllPosts(context *gin.Context) {
	var post model.Post
	posts, err := post.All()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  "No Post Found",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": posts,
	})
}

func GetUserPosts(context *gin.Context) {
    user, err := helper.CurrentUser(context)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    context.JSON(http.StatusOK, gin.H{"data": user.Posts})
}


