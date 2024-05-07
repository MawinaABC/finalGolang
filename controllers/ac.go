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

}

func getAllProductFromCart(id uint) []models.Cart {
	var cart []models.Cart
	initializers.DB.Where("user_id = ?, status = ?", id, 1).Find(&cart)
	
}
