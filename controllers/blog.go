package controllers

import (
	"github.com/gin-gonic/gin"
	"inktrail/config"
	"inktrail/repositories"

	"inktrail/models"
	"inktrail/utils"
	"net/http"
)

func AddBlog(context *gin.Context) {
	var input models.Blog
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := utils.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserID = user.ID

	savedEntry, err := repositories.SaveBlog(config.DB, &input)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func GetAllBlogs(context *gin.Context) {
	user, err := utils.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user.Blogs})
}
