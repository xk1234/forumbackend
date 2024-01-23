package controller

import (
	"cvwoapi/database"
	"cvwoapi/helper"
	"cvwoapi/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPost(context *gin.Context) {
	postID := context.Param("id")
	pid, err := strconv.ParseUint(postID, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := model.Post{}
	postReceived, err := post.GetPost(pid)
	context.JSON(http.StatusCreated, gin.H{"data": postReceived})
}

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
	sort := context.DefaultQuery("sort", "newest")
	posts, err := post.All(sort)
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

func UpdatePost(context *gin.Context) {
	postID := context.Param("id")
	pid, err := strconv.ParseUint(postID, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := helper.CurrentUser(context)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	currPost := model.Post{}
	err = database.Database.Model(model.Post{}).Where("id = ?", pid).Take(&currPost).Error
	if user.ID != currPost.UserID {
		context.JSON(http.StatusBadRequest, gin.H{"error": "wrong user"})
        return
    }
	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	post := model.Post{}
	err = json.Unmarshal(body, &post)
	// fmt.Println(currPost)
	if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	post.ID = currPost.ID
	post.UserID = currPost.UserID
	post.User = currPost.User
	postUpdated, err := post.UpdatePost()
	if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	fmt.Println(postUpdated)
	context.JSON(http.StatusOK, gin.H{"data": postUpdated})
}

func DeletePost(context *gin.Context) {
	postID := context.Param("id")
	pid, err := strconv.ParseUint(postID, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := model.Post{}
	err = database.Database.Model(model.Post{}).Where("id = ?", pid).Take(&post).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := helper.CurrentUser(context)
    if user.ID != post.UserID {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	_, err = post.DeletePost()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": "Deleted",
	})
}