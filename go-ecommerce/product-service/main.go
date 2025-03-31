package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
}

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/go_ecommerce") // Update password if needed
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
}

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, price, image FROM products")
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Image); err != nil {
			http.Error(w, "Error reading products", http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func addProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var p Product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO products (name, price, image) VALUES (?, ?, ?)", p.Name, p.Price, p.Image)
	if err != nil {
		http.Error(w, "Failed to add product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product added"})
}

func main() {
	initDB()
	defer db.Close()

	http.HandleFunc("/products", getProductsHandler)
	http.HandleFunc("/products/add", addProductHandler)

	fmt.Println("Product Service running on port 8081:")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

