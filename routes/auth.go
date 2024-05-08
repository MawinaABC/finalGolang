package routes

import (
	"github.com/MawinaABC/finalGolang/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)
}
