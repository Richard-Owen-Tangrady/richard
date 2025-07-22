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
	r.GET("/getproduct/:productid", controllers.GetProduct)
	r.POST("/createproduct", controllers.CreateProduct)
	r.PUT("/updateproduct", controllers.UpdateProduct)
	r.DELETE("/deleteproduct", controllers.DeleteProduct)

	r.Run()

}
