package routes

import (
	"github.com/MawinaABC/finalGolang/controllers"
	"github.com/MawinaABC/finalGolang/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.GET("/home", middlewares.ReqAuth(), controllers.CategoryList)
	r.GET("/home/:id", middlewares.ReqAuth(), controllers.ProductList)
	r.GET("/home/:id/:idd", middlewares.ReqAuth(), controllers.GetProduct)
	r.POST("/home/:id/:idd", middlewares.ReqAuth(), controllers.AddToCart)
	r.POST("/home/order", middlewares.ReqAuth(), controllers.Ordering)
	r.GET("/home/cart", middlewares.ReqAuth(), controllers.IndexCart)
	r.GET("/home/cart/:id", middlewares.ReqAuth(), controllers.GetCart)
	r.POST("/home/cart/:id", middlewares.ReqAuth(), controllers.DeleteCart)
}
