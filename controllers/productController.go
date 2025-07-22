package controllers

import (
	"net/http"

	"github.com/Richard-Owen-Tangrady/richard/initializers"
	"github.com/Richard-Owen-Tangrady/richard/models"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	result := initializers.DB.Find(&products)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to find products",
		})
		return
	}

	if result != nil {
		c.JSON(http.StatusOK, gin.H{
			"product": products,
		})
		return
	}
}

func GetProduct(c *gin.Context) {
	var product models.Product
	result := initializers.DB.First(&product, "product=?", product.ProductID)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to find product",
		})
		return
	}

	if result != nil {
		c.JSON(http.StatusOK, gin.H{
			"product": product,
		})
		return
	}
}

func CreateProduct(c *gin.Context) {
	var body models.BodyP

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	product := models.Product{ProductID: body.ProductID, Name: body.Name}
	result := initializers.DB.Create(&product)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create product",
		})
		return
	}

	if result != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "product created",
		})
		return
	}

}

func DeleteProduct(c *gin.Context) {
	var product models.Product
	result := initializers.DB.Delete(&models.Product{}, "product=?", product.ProductID)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete product",
		})
		return
	}

	if result != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "product deleted",
		})
		return
	}
}

func UpdateProduct(c *gin.Context) {
	var body models.BodyP

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var product models.Product
	initializers.DB.First(&product, "product=?", product.ProductID)

	initializers.DB.Model(&product).Updates(models.Product{
		ProductID: body.ProductID,
		Name:      body.Name,
		Quantity:  body.Quantity,
		Price:     body.Price,
	})

	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}
