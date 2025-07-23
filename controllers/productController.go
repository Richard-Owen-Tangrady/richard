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

	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})

}

func GetProduct(c *gin.Context) {
	productid := c.Param("product_id")

	var product models.Product
	result := initializers.DB.First(&product, productid)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to find product",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})

}

func CreateProduct(c *gin.Context) {
	var body models.ProductCreateRequest

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	product := models.Product{ProductID: body.ProductID, Name: body.Name, Quantity: body.Quantity, Price: body.Price}
	result := initializers.DB.Create(&product)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create product",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})

}

func DeleteProduct(c *gin.Context) {
	productid := c.Param("product_id")

	result := initializers.DB.Delete(&models.Product{}, productid)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete product",
		})
		return
	}

	c.Status(200)

}

func UpdateProduct(c *gin.Context) {
	productid := c.Param("product_id")
	var body models.ProductCreateRequest

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var product models.Product
	initializers.DB.First(&product, productid)

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
