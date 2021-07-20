package main

import (
	"github.com/TesyarRAz/sa-api-lks-jabar-2021/config"
	"github.com/TesyarRAz/sa-api-lks-jabar-2021/controller"
	"github.com/TesyarRAz/sa-api-lks-jabar-2021/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.NewDatabase()

	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(db.Middleware)

	r.POST("/login", controller.LoginUser)
	r.POST("/register", controller.RegisterUser)

	auth := r.Group("/", middleware.AuthorizedUser)

	{
		auth.GET("/user", controller.InfoUser)

		auth.GET("/menu", controller.IndexMenu)
		auth.GET("/menu/:id", controller.ShowMenu)
		auth.POST("/menu", controller.StoreMenu)
		auth.PUT("/menu/:id", controller.UpdateMenu)
		auth.DELETE("/menu/:id", controller.DestroyMenu)
	}

	addr := ":3000"

	r.Run(addr)

	// in, err := net.Listen("tcp4", addr)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// r.RunListener(in)
}
