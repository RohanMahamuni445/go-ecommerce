package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Product struct represents a product in the catalog
type Product struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

var db *gorm.DB

// Initialize Database
func initDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	db.AutoMigrate(&Product{})
	fmt.Println("Database connected and Product table migrated!")
}

// Add a new product
func addProduct(c *gin.Context) {
	var product Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	db.Create(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Product added successfully", "product": product})
}

// Get all products
func getProducts(c *gin.Context) {
	var products []Product
	db.Find(&products)
	c.JSON(http.StatusOK, products)
}

// Get product by ID
func getProductByID(c *gin.Context) {
	id := c.Param("id")
	var product Product
	if err := db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// Delete a product
func deleteProduct(c *gin.Context) {
	id := c.Param("id")
	db.Delete(&Product{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func main() {
	initDB()
	r := gin.Default()

	// Routes
	r.POST("/products", addProduct)        // Add a new product
	r.GET("/products", getProducts)        // Get all products
	r.GET("/products/:id", getProductByID) // Get a specific product
	r.DELETE("/products/:id", deleteProduct) // Delete a product

	// Start the service on port 8082
	r.Run(":8082")
}

