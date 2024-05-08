package controllers

import (
	"github.com/MawinaABC/finalGolang/initializers"
	"github.com/MawinaABC/finalGolang/models"
	"github.com/MawinaABC/finalGolang/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetProductForAdmin(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		c.JSON(401, err.Error())
		return
	}

	claims, err := utils.ParseToken(cookie)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	if claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "User is not admin"})
		return
	}

	id := c.Param("id")
	var product models.Product
	initializers.DB.First(&product, id)

	c.JSON(200, gin.H{
		"product": product,
	})
}

func CreateProduct(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		c.JSON(401, err.Error())
		return
	}

	claims, err := utils.ParseToken(cookie)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	if claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "User is not admin"})
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	initializers.DB.Create(&product)

	c.JSON(200, gin.H{
		"message":     "success",
		"new product": product,
	})
}

func DeleteProduct(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		c.JSON(401, err.Error())
		return
	}

	claims, err := utils.ParseToken(cookie)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	if claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "User is not admin"})
		return
	}

	id := c.Param("id")
	initializers.DB.Delete(&models.Product{}, id)

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func UpdateProduct(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		c.JSON(401, err.Error())
		return
	}

	claims, err := utils.ParseToken(cookie)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	if claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "User is not admin"})
		return
	}

	id := c.Param("id")
	var existingProduct, product models.Product
	initializers.DB.First(&existingProduct, id)
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Model(&existingProduct).Updates(models.Product{
		Name:     product.Name,
		Category: product.Category,
		Price:    product.Price,
	})
	c.JSON(200, gin.H{
		"message":         "success",
		"updated product": existingProduct,
	})
}

func IndexProduct(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		c.JSON(401, err.Error())
		return
	}

	claims, err := utils.ParseToken(cookie)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	if claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "User is not admin"})
		return
	}

	var products []models.Product
	initializers.DB.Find(&products)
	c.JSON(200, gin.H{
		"product list": products,
	})
}

func IndexOrder(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		c.JSON(401, err.Error())
		return
	}

	claims, err := utils.ParseToken(cookie)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	if claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "User is not admin"})
		return
	}

	var order []models.Order
	initializers.DB.Find(&order)

	if len(order) == 0 {
		c.JSON(401, gin.H{"error": "User doesn't have any order"})
		return
	}

	for i := 0; i < len(order); i++ {
		var cart []models.Cart
		initializers.DB.Find(&cart, "order_id = ?", order[i].ID)
		if len(cart) == 0 {
			continue
		}
		c.JSON(200, gin.H{
			"user email":                  order[i].UserEmail,
			"order #" + strconv.Itoa(i+1): cart,
			"total price":                 order[i].TotalPrice,
		})
	}
}
