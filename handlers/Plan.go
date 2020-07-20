package handlers //handlers/Products.go

import (
	"api/controllers"
	"api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetProductss ... Get all productss
// Upload example
// @Summary Gato mortal pra trás
// @Description Upload file
// @ID file.upload
// @Accept  multipart/form-data
// @Produce  json
// @Param   file formData file true  "this is a test file"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "We need ID!!"
// @Failure 404 {string} string "Can not find ID"
// @Router /file/upload [post]
func GetProducts(c *gin.Context) {
	var products []models.Products
	err := controllers.GetAllProducts(&products)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, products)
	}
}

//CreateProducts ... Create Products and save it on DB
func CreateProducts(c *gin.Context) {
	var products models.Products
	c.BindJSON(&products)
	err := controllers.CreateProducts(&products)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, products)
	}
}

//GetProductsByID ... Get the products by id
func GetProductsByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var products models.Products
	err := controllers.GetProductsByID(&products, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, products)
	}
}

//UpdateProducts ... Update the products information
func UpdateProducts(c *gin.Context) {
	var products models.Products
	id := c.Params.ByName("id")
	err := controllers.GetProductsByID(&products, id)
	if err != nil {
		c.JSON(http.StatusNotFound, products)
	}
	c.BindJSON(&products)
	err = controllers.UpdateProducts(&products, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, products)
	}
}

//DeleteProducts ... Delete the products
func DeleteProducts(c *gin.Context) {
	var products models.Products
	id := c.Params.ByName("id")
	err := controllers.DeleteProducts(&products, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
