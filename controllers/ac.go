package controllers

import (
	"github.com/MawinaABC/finalGolang/initializers"
	"github.com/MawinaABC/finalGolang/models"
	"github.com/MawinaABC/finalGolang/utils"
	"github.com/gin-gonic/gin"
)

func CategoryList(c *gin.Context) {
	var products []models.Product
	initializers.DB.Find(&products)
	for i := 0; i < len(products); i++ {
		c.JSON(200, gin.H{
			"ID":                  products[i].ID,
			"Category of product": products[i].Category,
		})
	}
	c.JSON(200, gin.H{
		"info only for api": "to place an order, go to the link localhost:3000/home/order",
	})
}

func Ordering(c *gin.Context) {
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

	cart := getAllProductFromCart(claims.UserId)
	totalPrice := countTotalPrice(cart)
	var user models.User
	initializers.DB.First(&user, claims.UserId)
	if totalPrice > user.Amount {
		c.JSON(200, gin.H{"error": "The user does not have enough credits"})
		return
	} else {
		order := models.Order{
			OrderItem:  cart,
			UserName:   user.Name,
			TotalPrice: totalPrice,
			Status:     1,
		}
		initializers.DB.Create(&order)
		initializers.DB.Model(&user).Updates(models.User{Amount: user.Amount - totalPrice})
		c.JSON(200, gin.H{
			"message":    "success",
			"order info": order,
		})
		updateAllCart(cart)
	}
}

func ProductList(c *gin.Context) {
	id := c.Param("id")
	var products []models.Product
	initializers.DB.Find(&products, id)
	for i := 0; i < len(products); i++ {
		c.JSON(200, gin.H{
			"ID of product":   products[i].ID,
			"Name of product": products[i].Name,
			"Price":           products[i].Price,
		})
	}
}

func GetProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("idd")
	initializers.DB.First(&product, id)
	c.JSON(200, gin.H{
		"Name of product": product.Name,
		"Description":     product.Description,
		"Price":           product.Price,
	})
}

func AddToCart(c *gin.Context) {
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

	id := c.Param("idd")
	var product models.Product
	initializers.DB.First(&product, id)

	cart := models.Cart{
		UserId:       claims.UserId,
		ProductName:  product.Name,
		ProductPrice: product.Price,
		Status:       1,
	}
	initializers.DB.Create(&cart)
	userCart := getAllProductFromCart(claims.UserId)
	c.JSON(200, gin.H{
		"message":     "success",
		"user cart":   userCart,
		"total price": countTotalPrice(userCart),
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
		"message":              "success",
		"updated product list": product,
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
		Name:        product.Name,
		Category:    product.Category,
		Description: product.Description,
		Price:       product.Price,
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

func getAllProductFromCart(id uint) []models.Cart {
	var cart []models.Cart
	initializers.DB.Where("user_id = ?, status = ?", id, 1).Find(&cart)
	return cart
}

func countTotalPrice(cart []models.Cart) float64 {
	var count float64
	for i := 0; i < len(cart); i++ {
		count += cart[i].ProductPrice
	}
	return count
}

func updateAllCart(cart []models.Cart) {
	for i := 0; i < len(cart); i++ {
		initializers.DB.Model(&cart[i]).Updates(models.Cart{
			Status: 0,
		})
	}
}
