package controllers

import (
	"net/http"

	"github.com/Richard-Owen-Tangrady/richard/initializers"
	"github.com/Richard-Owen-Tangrady/richard/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateCheckout(c *gin.Context) {
	var body models.CheckoutCreateRequest

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
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

	checkout := models.Checkout{
		CheckoutID:   newId,
		ProductRefer: body.ProductRefer,
		Product:      product,
		Quantity:     body.Quantity,
	}

	result := initializers.DB.Create(&checkout)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to create checkout",
			"details": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"checkout": checkout,
	})

}

func GetCheckout(c *gin.Context) {
	checkoutid := c.Param("checkout_id")

	var checkout models.Checkout
	result := initializers.DB.Find(&checkout, checkoutid)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to find checkout",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"checkout": checkout,
	})

}

func DeleteCheckout(c *gin.Context) {
	checkoutid := c.Param("checkout_id")

	var checkout models.Checkout
	result := initializers.DB.Find(&checkout, checkoutid)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to find checkout",
		})
		return
	}

	deleteresult := initializers.DB.Delete(&checkout, checkoutid)

	if deleteresult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete checkout",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "checkout deleted",
	})

}
