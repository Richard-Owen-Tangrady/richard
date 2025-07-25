package controllers

import (
	"net/http"
	"time"

	"github.com/Richard-Owen-Tangrady/richard/initializers"
	"github.com/Richard-Owen-Tangrady/richard/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetCart(c *gin.Context) {
	cartid := c.Param("cart_id")

	var cart models.Cart
	result := initializers.DB.Where("cart_id = ?", cartid).First(&cart)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to find cart",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cart": cart,
	})

}

func CreateCart(c *gin.Context) {
	var body models.CartCreateRequest

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var user models.User
	if initializers.DB.First(&user, "user_id=?", body.UserRefer).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to find user",
		})
		return
	}

	var product models.Product
	if initializers.DB.First(&product, "product_id=?", body.ProductRefer).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to find product",
		})
		return
	}

	newId := uuid.New().String()

	cart := models.Cart{
		CartID:       newId,
		UserRefer:    body.UserRefer,
		User:         user,
		ProductRefer: body.ProductRefer,
		Product:      product,
		Quantity:     body.Quantity,
		CreatedAt:    time.Now(),
	}

	result := initializers.DB.Create(&cart)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to create cart",
			"details": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cart": cart,
	})

}

func DeleteCart(c *gin.Context) {
	cartid := c.Param("cart_id")

	var cart models.Cart
	result := initializers.DB.Where("cart_id = ?", cartid).First(&cart)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to find cart",
		})
		return
	}

	deleteResult := initializers.DB.Where("cart_id =?", cartid).Delete(&cart)

	if deleteResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete cart",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "checkout deleted",
	})

}
