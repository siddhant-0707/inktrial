package controllers

import (
	"github.com/gin-gonic/gin"
	"inktrail/config"
	"inktrail/models"
	"inktrail/repositories"
	"inktrail/utils"
	"net/http"
	"strconv"
)

func AddComment(c *gin.Context) {
	var input models.Comment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blogIDStr := c.Param("id")
	id, err := strconv.ParseUint(blogIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	db := config.DB
	// Check if the blog with the given ID exists
	blog, err := repositories.FindBlogByID(db, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	input.BlogID = blog.ID
	input.UserID = user.ID

	savedEntry, err := repositories.SaveComment(config.DB, &input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

// GetCommentsByBlogID gets all comments for a blog by its ID
func GetCommentsByBlogID(c *gin.Context) {
	blogIDStr := c.Param("id")
	blogID, err := strconv.ParseUint(blogIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}
	db := config.DB

	// Check if the blog with the given ID exists
	blog, err := repositories.FindBlogByID(db, uint(blogID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	// Get all comments for the blog from the database
	comments, err := repositories.FindCommentsByBlogID(db, blog.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comments})
}
