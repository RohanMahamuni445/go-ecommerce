package main

import (
	"io"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// forwardRequest forwards the incoming request to the appropriate microservice
func forwardRequest(c *gin.Context, serviceURL string) {
	req, err := http.NewRequest(c.Request.Method, serviceURL, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}
	req.Header = c.Request.Header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reach service"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func main() {
	r := gin.Default()

	// ✅ Enable CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Allow frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// ✅ User Authentication Service (Port 8085)
	r.POST("/register", func(c *gin.Context) {
		forwardRequest(c, "http://localhost:8085/register")
	})
	r.POST("/login", func(c *gin.Context) {
		forwardRequest(c, "http://localhost:8085/login")
	})

	// ✅ Product Catalog Service (Port 8082)
	r.GET("/products", func(c *gin.Context) {
		forwardRequest(c, "http://localhost:8082/products")
	})
	r.GET("/products/:id", func(c *gin.Context) {
		forwardRequest(c, "http://localhost:8082/products/" + c.Param("id"))
	})

	// ✅ Shopping Cart Service (Port 8083)
	r.GET("/cart/:user_id", func(c *gin.Context) {
		forwardRequest(c, "http://localhost:8083/cart/" + c.Param("user_id"))
	})
	r.POST("/cart", func(c *gin.Context) {
		forwardRequest(c, "http://localhost:8083/cart")
	})
	r.DELETE("/cart/:id", func(c *gin.Context) {
		forwardRequest(c, "http://localhost:8083/cart/" + c.Param("id"))
	})

	// ✅ Order Management Service (Port 8084)
	r.POST("/orders", func(c *gin.Context) {
		forwardRequest(c, "http://localhost:8084/orders")
	})
	r.GET("/orders", func(c *gin.Context) {
		forwardRequest(c, "http://localhost:8084/orders")
	})

	// ✅ Payment Service (Port 8084)
	r.POST("/payments", func(c *gin.Context) {
		forwardRequest(c, "http://localhost:8084/payments")
	})

	// Start API Gateway on port 8086
	r.Run(":8086")
}

