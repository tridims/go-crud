package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tridims/go-crud/initializers"
	"github.com/tridims/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// sent response
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// get id off the url
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// get the id off the url
	id := c.Param("id")

	// get the data off the request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	// updated it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// respond with it
	c.JSON(200, gin.H{
		"updated_post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.Status(200)
}
