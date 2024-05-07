package main

import (
	"github.com/MawinaABC/finalGolang/controllers"
	"github.com/MawinaABC/finalGolang/initializers"
	"github.com/MawinaABC/finalGolang/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVar()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/home", middlewares.ReqAuth(), controllers.CategoryList)
	r.GET("/home/:id", middlewares.ReqAuth(), controllers.ProductList)
	r.GET("/home/:id/:idd", middlewares.ReqAuth(), controllers.GetProduct)
	r.POST("/home/:id/:idd", middlewares.ReqAuth(), controllers.AddToCart)
	r.POST("/home/order", middlewares.ReqAuth(), controllers.Ordering)
	r.GET("/admin", middlewares.ReqAuth(), controllers.IndexProduct)
	r.POST("/admin/create", middlewares.ReqAuth(), controllers.CreateProduct)
	r.GET("/admin/:id", middlewares.ReqAuth(), controllers.GetProductForAdmin)
	r.PUT("/admin/:id", middlewares.ReqAuth(), controllers.UpdateProduct)
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)
	r.Run()
}
