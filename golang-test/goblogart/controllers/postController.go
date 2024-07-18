package controllers

import (
	"fmt"
	"goblogart/inits"
	"goblogart/models"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var body struct {
		Title  string
		Body   string
		Likes  int
		Draft  bool
		Author string
		UserID uint `json:"user_id"`
	}

	user, exists := c.Get("user")

	if !exists {
		c.JSON(500, gin.H{
			"error": "user not found",
		})
		return
	}
	body.UserID = user.(models.User).ID

	c.BindJSON(&body)

	// Should be in DAL
	post := models.Post{
		Title:  body.Title,
		Body:   body.Body,
		Likes:  body.Likes,
		Draft:  body.Draft,
		Author: body.Author,
		UserID: body.UserID,
	}

	// in routes/services?
	fmt.Println(post)
	result := inits.DB.Create(&post)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error,
		})
		return
	}
	c.JSON(200, gin.H{
		"data": post,
	})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post

	result := inits.DB.Find(&posts)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error,
		})
		return
	}
	c.JSON(200, gin.H{
		"data": posts,
	})
}

func GetPost(c *gin.Context) {
	var post models.Post

	result := inits.DB.First(&post, c.Param("id"))
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error,
		})
		return
	}
	c.JSON(200, gin.H{
		"data": post,
	})
}

func UpdatePost(c *gin.Context) {
	var body struct {
		Title  string
		Body   string
		Likes  int
		Draft  bool
		Author string
	}

	c.BindJSON(&body)

	var post models.Post

	result := inits.DB.First(&post, c.Param("id"))
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error,
		})
		return
	}

	inits.DB.Model(&post).Updates(models.Post{
		Title:  body.Title,
		Body:   body.Body,
		Likes:  body.Likes,
		Draft:  body.Draft,
		Author: body.Author,
	})

	c.JSON(200, gin.H{
		"data": post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	inits.DB.Delete(&models.Post{}, id)

	c.JSON(200, gin.H{
		"data": "post has been deleted successfully",
	})
}
