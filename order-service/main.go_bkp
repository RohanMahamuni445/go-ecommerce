package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Order struct represents an order placed by a user
type Order struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	UserID    uint    `json:"user_id"`
	Total     float64 `json:"total"`
	Status    string  `json:"status"` // Pending, Completed, Failed
}

// Payment struct represents a payment for an order
type Payment struct {
	ID      uint    `json:"id" gorm:"primaryKey"`
	OrderID uint    `json:"order_id"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"` // Success, Failed
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
	db.AutoMigrate(&Order{}, &Payment{})
	fmt.Println("Database connected and Order/Payment tables migrated!")
}

// Place a new order
func placeOrder(c *gin.Context) {
	var order Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	order.Status = "Pending"
	db.Create(&order)
	c.JSON(http.StatusOK, gin.H{"message": "Order placed successfully", "order": order})
}

// Get order by ID
func getOrderByID(c *gin.Context) {
	id := c.Param("id")
	var order Order
	if err := db.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

// Make a payment
func makePayment(c *gin.Context) {
	var payment Payment
	if err := c.BindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	// Simulate payment success
	payment.Status = "Success"
	db.Create(&payment)

	// Update order status to Completed
	db.Model(&Order{}).Where("id = ?", payment.OrderID).Update("status", "Completed")

	c.JSON(http.StatusOK, gin.H{"message": "Payment successful", "payment": payment})
}

// Get payment details by order ID
func getPaymentByOrderID(c *gin.Context) {
	orderID := c.Param("order_id")
	var payment Payment
	if err := db.Where("order_id = ?", orderID).First(&payment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}
	c.JSON(http.StatusOK, payment)
}

func main() {
	initDB()
	r := gin.Default()

	// Routes
	r.POST("/orders", placeOrder)         // Place an order
	r.GET("/orders/:id", getOrderByID)    // Get order details
	r.POST("/payments", makePayment)      // Make a payment
	r.GET("/payments/:order_id", getPaymentByOrderID) // Get payment details by order ID

	// Start the service on port 8084
	r.Run(":8084")
}

