package controllers

import (
	"net/http"

	"github.com/Richard-Owen-Tangrady/richard/initializers"
	"github.com/Richard-Owen-Tangrady/richard/models"
	"github.com/gin-gonic/gin"
)

func GetCart(c *gin.Context) {
	cartid := c.Param("cart_id")

	var cart models.Cart
	result := initializers.DB.First(&cart, cartid)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to find products",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cart": cart,
	})

}

func CreateCart(c *gin.Context) {
	var cart models.Cart
	var pro models.Product
	var user models.User

	if c.Bind(&cart) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	cart = models.Cart{CartID: cart.CartID, UserID: user.UserID, ProductId: pro.ProductID, Quantity: cart.Quantity}
	result := initializers.DB.Create(&cart)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create product",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cart": cart,
	})

}
