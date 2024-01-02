package controller

import (
    "cvwoapi/helper"
    "cvwoapi/model"
    "github.com/gin-gonic/gin"
    "net/http"
	"strconv"
)

func AddComment(context *gin.Context) {
    var input model.Comment
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
	postID := context.Param("id")
	pid64, err := strconv.ParseUint(postID, 10, 64)
	if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
        return
    }
    if pid64 == 0 {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Post ID cannot be zero"})
        return
    }
	pid := uint(pid64)
	input.PostRefer = pid
    savedComment, err := input.Save()
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    context.JSON(http.StatusCreated, gin.H{"data": savedComment})
}

func GetComments(context *gin.Context) {
	var comment model.Comment
	postID := context.Param("id")
	pid64, err := strconv.ParseUint(postID, 10, 64)
	if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
        return
    }
    if pid64 == 0 {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Post ID cannot be zero"})
        return
    }
	pid := uint(pid64)
	comments, err := comment.GetComments(pid)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  "No comments found",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": comments,
	})
}