CREATE DATABASE IF NOT EXISTS ecommerce;
USE ecommerce;

CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS products (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    image VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS cart (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL
);

CREATE TABLE IF NOT EXISTS orders (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    total FLOAT NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'Pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS payments (
    id INT PRIMARY KEY AUTO_INCREMENT,
    order_id INT NOT NULL,
    user_id INT NOT NULL,
    amount FLOAT NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'Pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 🔽 Insert sample data into products (with image URLs)
INSERT INTO products (name, description, price, image) VALUES
('iPhone 14', 'Latest Apple phone', 999.99, 'https://m.media-amazon.com/images/I/31VjlrbE3bL._AC_SY350_QL15_.jpg'),
('Samsung Galaxy S23', 'Latest Samsung phone', 899.99, 'https://m.media-amazon.com/images/I/71OXmy3NMCL._AC_UF1000,1000_QL80_.jpg'),
('MacBook Pro', 'Apple laptop', 1999.99, 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRo9aPRBIi0vWKYuNaV9FIhV0dzaJH89bxEpA&s');

-- 🔽 Insert sample user
INSERT INTO users (name, email, password_hash) VALUES
('Rohan', 'rohan@example.com', 'hashed_password_123');

