package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type CartItem struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(mysql:3306)/ecommerce")
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
}

func addToCartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var item CartItem
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO cart (product_id, quantity) VALUES (?, ?)", item.ProductID, item.Quantity)
	if err != nil {
		http.Error(w, "Failed to add to cart", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Item added to cart"})
}

func getCartHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, product_id, quantity FROM cart")
	if err != nil {
		http.Error(w, "Failed to fetch cart", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cart []CartItem
	for rows.Next() {
		var item CartItem
		if err := rows.Scan(&item.ID, &item.ProductID, &item.Quantity); err != nil {
			http.Error(w, "Error reading cart", http.StatusInternalServerError)
			return
		}
		cart = append(cart, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

func main() {
	initDB()
	defer db.Close()

	http.HandleFunc("/cart", getCartHandler)
	http.HandleFunc("/cart/add", addToCartHandler)

	fmt.Println("Cart Service running on port 8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

