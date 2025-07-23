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

	r.GET("/getproducts", controllers.GetProducts)
	r.GET("/getproduct/:product_id", controllers.GetProduct)
	r.POST("/createproduct", controllers.CreateProduct)
	r.PUT("/updateproduct/:product_id", controllers.UpdateProduct)
	r.DELETE("/deleteproduct/:product_id", controllers.DeleteProduct)

	r.GET("/cart/:cart_id", controllers.GetCart)
	r.POST("/cart", controllers.CreateCart)

	r.Run()

}
