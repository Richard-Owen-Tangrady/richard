package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Richard-Owen-Tangrady/richard/controllers"
	"github.com/Richard-Owen-Tangrady/richard/initializers"
	"github.com/Richard-Owen-Tangrady/richard/middleware"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDb()
	initializers.SyncDataBase()
}

func main() {
	r := gin.Default()
	muxr := mux.NewRouter()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	muxr.HandleFunc("/product", controllers.GetProducts).Methods("GET")
	muxr.HandleFunc("/product/{productid}", controllers.GetProduct).Methods("GET")
	muxr.HandleFunc("/product", controllers.CreateProduct).Methods("POST")
	muxr.HandleFunc("/product/{productid}", controllers.UpdateProduct).Methods("PUT")
	muxr.HandleFunc("/product/{productid}", controllers.DeleteProduct).Methods("DELETE")

	fmt.Printf("Starting server at port 3030\n")
	log.Fatal(http.ListenAndServe(":3030", muxr))
	r.Run()

}
