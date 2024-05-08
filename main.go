package main

import (
	"github.com/MawinaABC/finalGolang/initializers"
	"github.com/MawinaABC/finalGolang/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVar()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	routes.AuthRoutes(r)
	routes.AdminRoutes(r)
	routes.UserRoutes(r)
	r.Run()
}
