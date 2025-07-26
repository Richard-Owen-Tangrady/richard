package controllers

import (
	"net/http"

	"github.com/Richard-Owen-Tangrady/richard/initializers"
	"github.com/Richard-Owen-Tangrady/richard/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	result := initializers.DB.Where("product_id = ?", productid).First(&product)

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

	newId := uuid.New().String()

	product := models.Product{
		ProductID:   newId,
		Name:        body.Name,
		Quantity:    body.Quantity,
		Price:       body.Price,
		Description: "",
	}

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

	var product models.Product
	result := initializers.DB.Where("product_id = ?", productid).First(&product)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to find product",
		})
		return
	}

	deleteResult := initializers.DB.Where("product_id = ?", productid).Delete(&product)

	if deleteResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete product",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product deleted",
	})

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
	result := initializers.DB.Where("product_id = ?", productid).First(&product)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to find product",
		})
		return
	}

	update := make(map[string]interface{})
	if body.Name != "" {
		update["Name"] = body.Name
	}
	if body.Quantity != 0 {
		update["Quantity"] = body.Quantity
	}
	if body.Price != 0.0 {
		update["Price"] = body.Price
	}

	updateResult := initializers.DB.Model(&product).Updates(update)

	if updateResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update product",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}
