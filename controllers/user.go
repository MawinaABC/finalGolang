package controllers

import (
	"github.com/MawinaABC/finalGolang/initializers"
	"github.com/MawinaABC/finalGolang/models"
	"github.com/MawinaABC/finalGolang/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CategoryList(c *gin.Context) {
	var products []models.Product
	initializers.DB.Find(&products)
	if products == nil {
		c.Status(401)
		return
	}
	set := make(map[string]bool)
	for i := 0; i < len(products); i++ {
		if set[products[i].Category] {
			continue
		}
		set[products[i].Category] = true
		c.JSON(200, gin.H{
			"ID":                  products[i].ID,
			"Category of product": products[i].Category,
		})
	}
	c.JSON(200, gin.H{
		"info only for api": "to place an order, go to the link localhost:3000/home/order",
	})
}

func OrderHistory(c *gin.Context) {
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

	if claims.Role == "admin" {
		c.JSON(401, gin.H{"error": "This function only for users"})
		return
	}

	var user models.User
	initializers.DB.First(&user, claims.UserId)

	var order []models.Order
	initializers.DB.Find(&order, "user_email = ?", user.Email)
	if len(order) == 0 {
		c.JSON(401, gin.H{"error": "User doesn't have any order"})
		return
	}

	c.JSON(200, gin.H{
		"history of user orders": "User has " + strconv.Itoa(len(order)) + " orders",
	})

	for i := 0; i < len(order); i++ {
		var cart []models.Cart
		initializers.DB.Find(&cart, "order_id = ?", order[i].ID)
		if len(cart) == 0 {
			continue
		}
		c.JSON(200, gin.H{
			"order #" + strconv.Itoa(i+1): cart,
			"total price":                 order[i].TotalPrice,
		})
	}
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

	if claims.Role == "admin" {
		c.JSON(401, gin.H{"error": "This function only for users"})
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
		if len(cart) == 0 {
			c.JSON(400, gin.H{"error": "User doesn't have items"})
			return
		}
		order := models.Order{
			OrderItem:  cart,
			UserName:   user.Name,
			UserEmail:  user.Email,
			TotalPrice: totalPrice,
		}
		initializers.DB.Create(&order)
		initializers.DB.Model(&user).Updates(models.User{Amount: user.Amount - totalPrice})
		c.JSON(200, gin.H{
			"message":    "success",
			"order info": order,
		})
		for i := 0; i < len(cart); i++ {
			initializers.DB.Model(&cart[i]).Updates(models.Cart{
				Status: "false",
			})
		}
	}
}

func ProductList(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	initializers.DB.First(&product, id)
	if product.ID == 0 {
		c.Status(401)
		return
	}

	products := getAllProductByCategory(product.Category)
	for i := 0; i < len(products); i++ {
		c.JSON(200, gin.H{
			"ID of product":   products[i].ID,
			"Name of product": products[i].Name,
			"Price":           products[i].Price,
		})
	}
}

func IndexCart(c *gin.Context) {
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

	if claims.Role == "admin" {
		c.JSON(401, gin.H{"error": "This function only for users"})
		return
	}

	cart := getAllProductFromCart(claims.UserId)
	totalPrice := countTotalPrice(cart)
	c.JSON(200, gin.H{
		"cart items":  cart,
		"total price": totalPrice,
	})
}

func GetCart(c *gin.Context) {
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

	if claims.Role == "admin" {
		c.JSON(401, gin.H{"error": "This function only for users"})
		return
	}

	id := c.Param("id")
	var cart models.Cart
	initializers.DB.First(&cart, id)

	if cart.UserId != claims.UserId || cart.Status == "false" {
		c.Status(401)
		return
	}

	c.JSON(200, gin.H{
		"cart": cart,
	})
}

func DeleteCart(c *gin.Context) {
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

	if claims.Role == "admin" {
		c.JSON(401, gin.H{"error": "This function only for users"})
		return
	}

	id := c.Param("id")
	var cart models.Cart
	initializers.DB.First(&cart, id)
	if cart.UserId != claims.UserId || cart.Status == "false" {
		c.Status(401)
		return
	}

	initializers.DB.Delete(&models.Cart{}, id)
	c.JSON(200, gin.H{
		"message": "deleted successfully",
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

	if claims.Role == "admin" {
		c.JSON(401, gin.H{"error": "This function only for users"})
		return
	}

	id := c.Param("idd")
	var product models.Product
	initializers.DB.First(&product, id)

	if product.ID == 0 || product.Name == "" {
		c.Status(401)
		return
	}

	cart := models.Cart{
		UserId:       claims.UserId,
		ProductName:  product.Name,
		ProductPrice: product.Price,
		Status:       "true",
	}
	initializers.DB.Create(&cart)
	userCart := getAllProductFromCart(claims.UserId)
	c.JSON(200, gin.H{
		"message":     "success",
		"user cart":   userCart,
		"total price": countTotalPrice(userCart),
	})
}

func GetProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("idd")
	initializers.DB.First(&product, id)

	if product.ID == 0 || product.Name == "" {
		c.Status(401)
		return
	}

	c.JSON(200, gin.H{
		"Name of product": product.Name,
		"Price":           product.Price,
	})
}

func getAllProductFromCart(id uint) []models.Cart {
	var cart []models.Cart
	initializers.DB.Find(&cart, "user_id = ? AND status = ?", id, "true")
	return cart
}

func countTotalPrice(cart []models.Cart) float64 {
	var count float64
	for i := 0; i < len(cart); i++ {
		count += cart[i].ProductPrice
	}
	return count
}

func getAllProductByCategory(category string) []models.Product {
	var products []models.Product
	initializers.DB.Find(&products, "category = ?", category)
	return products
}

func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	initializers.DB.Create(models.Comment{
		Text: comment.Text,
	})
	c.JSON(200, gin.H{
		"comment": comment.Text,
	})
}
