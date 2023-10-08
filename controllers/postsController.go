package controllers

import (
	"github.com/AlmedinNasufi/go-crud/initializers"
	"github.com/AlmedinNasufi/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	// get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	// create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(40)
		return
	}

	// return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// get the post
	var posts []models.Post
	initializers.DB.Find(&posts)

	// return response
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostById(c *gin.Context) {

	//get id off url
	id := c.Param("id")
	var post models.Post

	// find the posts
	initializers.DB.First(&post, id)
	// return the response

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// get the id off the url
	id := c.Param("id")
	// get the data of request body
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)
	// Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	// update it

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// response with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	// get the id off the url
	id := c.Param("id")

	// find the post that we are deleting

	initializers.DB.Delete(&models.Post{}, id)

	// response with it
	c.JSON(200, gin.H{
		"message": "deleted successfully",
	})

}
