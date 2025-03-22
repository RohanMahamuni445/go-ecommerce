package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// CartItem represents an item in the shopping cart
type CartItem struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	UserID    uint    `json:"user_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
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
	db.AutoMigrate(&CartItem{})
	fmt.Println("Database connected and Cart table migrated!")
}

// Add item to cart
func addToCart(c *gin.Context) {
	var item CartItem
	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	db.Create(&item)
	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart", "cartItem": item})
}

// Get cart items by user ID
func getCartByUserID(c *gin.Context) {
	var cart []CartItem
	userID := c.Param("user_id")
	db.Where("user_id = ?", userID).Find(&cart)
	c.JSON(http.StatusOK, cart)
}

// Remove item from cart
func removeFromCart(c *gin.Context) {
	id := c.Param("id")
	db.Delete(&CartItem{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}

// Clear the cart for a user
func clearCart(c *gin.Context) {
	userID := c.Param("user_id")
	db.Where("user_id = ?", userID).Delete(&CartItem{})
	c.JSON(http.StatusOK, gin.H{"message": "Cart cleared"})
}

func main() {
	initDB()
	r := gin.Default()

	// Routes
	r.POST("/cart", addToCart)                // Add item to cart
	r.GET("/cart/:user_id", getCartByUserID)  // Get user's cart items
	r.DELETE("/cart/:id", removeFromCart)     // Remove specific item
	r.DELETE("/cart/clear/:user_id", clearCart) // Clear all cart items for user

	// Start the service on port 8083
	r.Run(":8083")
}

