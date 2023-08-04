package controllers

import (
	"github.com/gin-gonic/gin"
	"inktrail/config"
	"inktrail/repositories"
	"strconv"

	"inktrail/models"
	"inktrail/utils"
	"net/http"
)

func AddBlog(c *gin.Context) {
	var input models.Blog
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := utils.CurrentUser(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserID = user.ID

	savedEntry, err := repositories.SaveBlog(config.DB, &input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func GetAllBlogs(c *gin.Context) {
	user, err := utils.CurrentUser(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user.Blogs})
}

// GetBlogByID returns the details of a specific blog post by its ID
func GetBlogByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	db := config.DB
	blog, err := repositories.FindBlogByID(db, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": blog})
}
