package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define struct to hold data
type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Mock data
var products = []Product{
	{ID: "1", Name: "Laptop", Price: 1000},
	{ID: "2", Name: "Phone", Price: 500},
}

func main() {
	r := gin.Default()

	// Get the list of products
	r.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, products)
	})

	// Get a product by ID
	r.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, p := range products {
			if p.ID == id {
				c.JSON(http.StatusOK, p)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
	})

	// Add a new product
	r.POST("/products", func(c *gin.Context) {
		var newProduct Product
		if err := c.ShouldBindJSON(&newProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		products = append(products, newProduct)
		c.JSON(http.StatusCreated, newProduct)
	})

	// Update a product
	r.PUT("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedProduct Product

		if err := c.ShouldBindJSON(&updatedProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i, p := range products {
			if p.ID == id {
				products[i] = updatedProduct
				c.JSON(http.StatusOK, updatedProduct)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
	})

	// Delete a product
	r.DELETE("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, p := range products {
			if p.ID == id {
				products = append(products[:i], products[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
	})

	// Run the server on port 8080
	r.Run(":8080")
}
