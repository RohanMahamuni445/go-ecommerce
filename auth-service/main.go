package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

// User struct represents a user in the database
type User struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    Email    string `json:"email" gorm:"unique"`
    Password string `json:"password"`
}

var db *gorm.DB
var jwtKey = []byte("secret_key") // Change this in production

// JWT claims structure
type Claims struct {
    Email string `json:"email"`
    jwt.RegisteredClaims
}

// Initialize Database
func initDB() {
    dsn := "rohan:Rohan@445@tcp(mysql-db:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }
    db.AutoMigrate(&User{})
    fmt.Println("Database connected and User table migrated!")
}

// Hash password
func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

// Compare password with hashed password
func checkPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

// Register a new user
func registerUser(c *gin.Context) {
    var user User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Ensure email is unique
    if err := db.Where("email = ?", user.Email).First(&User{}).Error; err == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
        return
    }

    // Hash the password before storing
    hashedPassword, err := hashPassword(user.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
        return
    }
    user.Password = hashedPassword

    db.Create(&user)
    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// Login a user
func loginUser(c *gin.Context) {
    var input struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    var user User

    if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Check if user exists
    if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Verify password
    if !checkPasswordHash(input.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Generate JWT token
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Email: user.Email,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func main() {
    initDB()
    r := gin.Default()

    // Routes
    r.POST("/register", registerUser) // Register user
    r.POST("/login", loginUser)       // User login

    // Start the service on port 8085
    r.Run(":8085")
}

