package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var db *gorm.DB

// Initialize Database
func initDB() {
	dsn := "rohan:Rohan@445@tcp(127.0.0.1:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	db.AutoMigrate(&User{})
	fmt.Println("Database connected and migrated!")
}

// Register User
func registerUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// Login User (Mock Authentication)
func loginUser(c *gin.Context) {
	var user User
	var input User

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := db.Where("username = ? AND password = ?", input.Username, input.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}

func main() {
	initDB()
	r := gin.Default()

	r.POST("/register", registerUser)
	r.POST("/login", loginUser)

	r.Run(":8081")
}

