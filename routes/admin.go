package routes

import (
	"github.com/MawinaABC/finalGolang/controllers"
	"github.com/MawinaABC/finalGolang/middlewares"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine) {
	r.GET("/admin", middlewares.ReqAuth(), controllers.IndexProduct)
	r.POST("/admin/create", middlewares.ReqAuth(), controllers.CreateProduct)
	r.GET("/admin/:id", middlewares.ReqAuth(), controllers.GetProductForAdmin)
	r.PUT("/admin/:id", middlewares.ReqAuth(), controllers.UpdateProduct)
	r.DELETE("/admin/:id", middlewares.ReqAuth(), controllers.DeleteProduct)
}
