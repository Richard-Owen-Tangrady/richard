package main

import (
	"github.com/Richard-Owen-Tangrady/richard/controllers"
	"github.com/Richard-Owen-Tangrady/richard/initializers"
	"github.com/Richard-Owen-Tangrady/richard/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDb()
	initializers.SyncDataBase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.GET("/products", controllers.GetProducts)
	r.GET("/product/:product_id", controllers.GetProduct)
	r.POST("/product", controllers.CreateProduct)
	r.PUT("/product/:product_id", controllers.UpdateProduct)
	r.DELETE("/product/:product_id", controllers.DeleteProduct)

	r.GET("/cart/:cart_id", controllers.GetCart)
	r.POST("/cart", controllers.CreateCart)
	r.DELETE("/cart/:cart_id", controllers.DeleteProduct)

	r.Run()

}
