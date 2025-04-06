package main

import (
        "bytes"
        "io"
        "log"
        "net/http"
)

func enableCORS(w http.ResponseWriter) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func handlePreflight(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodOptions {
                enableCORS(w)
                w.WriteHeader(http.StatusOK)
                return
        }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
        enableCORS(w)
        w.Write([]byte("Welcome to the API Gateway"))
}

func productHandler(w http.ResponseWriter, r *http.Request) {
        enableCORS(w)
        handlePreflight(w, r)

        if r.Method == http.MethodGet {
                resp, err := http.Get("http://product-service:8081/products")
                if err != nil {
                        http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
                        return
                }
                defer resp.Body.Close()
                w.Header().Set("Content-Type", "application/json")
                io.Copy(w, resp.Body)
        } else if r.Method == http.MethodPost {
                body, _ := io.ReadAll(r.Body)
                resp, err := http.Post("http://product-service:8081/products/add", "application/json", bytes.NewReader(body))
                if err != nil {
                        http.Error(w, "Failed to add product", http.StatusInternalServerError)
                        return
                }
                defer resp.Body.Close()
                w.Header().Set("Content-Type", "application/json")
                io.Copy(w, resp.Body)
        } else {
                http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        }
}

func cartHandler(w http.ResponseWriter, r *http.Request) {
        enableCORS(w)
        handlePreflight(w, r)

        resp, err := http.Get("http://cart-service:8082/cart")
        if err != nil {
                http.Error(w, "Failed to fetch cart", http.StatusInternalServerError)
                return
        }
        defer resp.Body.Close()
        w.Header().Set("Content-Type", "application/json")
        io.Copy(w, resp.Body)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
        enableCORS(w)
        handlePreflight(w, r)

        resp, err := http.Post("http://auth-service:8083/register", "application/json", r.Body)
        if err != nil {
                http.Error(w, "Request failed", http.StatusInternalServerError)
                return
        }
        defer resp.Body.Close()
        w.Header().Set("Content-Type", "application/json")
        io.Copy(w, resp.Body)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
        enableCORS(w)
        handlePreflight(w, r)

        resp, err := http.Post("http://auth-service:8083/login", "application/json", r.Body)
        if err != nil {
                http.Error(w, "Request failed", http.StatusInternalServerError)
                return
        }
        defer resp.Body.Close()
        w.Header().Set("Content-Type", "application/json")
        io.Copy(w, resp.Body)
}

func main() {
        http.HandleFunc("/", homeHandler)
        http.HandleFunc("/products", productHandler)
        http.HandleFunc("/cart", cartHandler)
        http.HandleFunc("/register", registerHandler)
        http.HandleFunc("/login", loginHandler)

        log.Println("Starting API Gateway on :8080")
        err := http.ListenAndServe(":8080", nil)
        if err != nil {
                log.Fatal("Server failed to start: ", err)
        }
}
#fjfjfjfj
