package controller

import (
	"cvwoapi/database"
	"cvwoapi/helper"
	"cvwoapi/model"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
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
	postID := context.Param("pid")
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
	postID := context.Param("pid")
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

func UpdateComment(context *gin.Context) {
	commentID := context.Param("id")
	pid, err := strconv.ParseUint(commentID, 10, 64)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	
	user, err := helper.CurrentUser(context)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	
	origComment := model.Comment{}
	err = database.Database.Model(model.Comment{}).Where("id = ?", pid).Take(&origComment).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	uid := user.ID
	if uid != origComment.UserID {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Wrong User"})
        return
	}
	
	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	comment := model.Comment{}
	err = json.Unmarshal(body, &comment)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	comment.ID = origComment.ID 
	comment.UserID = origComment.UserID
	comment.PostRefer = origComment.PostRefer

	commentUpdated, err := comment.UpdateComment()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": commentUpdated,
	})
	
}

func DeleteComment(context *gin.Context) {
    commentID := context.Param("id")
    cid, err := strconv.ParseUint(commentID, 10, 64)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := helper.CurrentUser(context)
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get current user"})
        return
    }
    var comment model.Comment
    err = database.Database.Where("id = ? AND user_id = ?", cid, user.ID).First(&comment).Error
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Comment not found or not authorized"})
        return
    }
    _, err = comment.DeleteComment();
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "response": "Deleted Comment"})
}