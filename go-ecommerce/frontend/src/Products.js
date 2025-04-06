import React, { useEffect, useState } from "react";
import axios from "axios";

function Products() {
  const [products, setProducts] = useState([]);
  const [error, setError] = useState("");

  useEffect(() => {
    axios
      .get("http://localhost:8080/products")
      .then((response) => {
        console.log("✅ Product API response:", response.data);

        // Handle if data is wrapped in an object like { products: [...] }
        if (Array.isArray(response.data)) {
          setProducts(response.data);
        } else if (Array.isArray(response.data.products)) {
          setProducts(response.data.products);
        } else {
          console.error("❌ Unexpected response format:", response.data);
          setError("Invalid product data format.");
        }
      })
      .catch((error) => {
        console.error("❌ Failed to fetch products:", error);
        setError("Failed to load products.");
      });
  }, []);

  const addToCart = (product) => {
    axios
      .post("http://localhost:8080/cart", {
        product_id: product.id,
        quantity: 1,
      })
      .then(() => alert(`${product.name} added to cart!`))
      .catch((error) => {
        console.error("❌ Failed to add to cart:", error);
        alert("Failed to add to cart");
      });
  };

  return (
    <div>
      <h2>🛍️ Products</h2>
      {error && <p style={{ color: "red" }}>{error}</p>}
      {products.length === 0 && !error && <p>Loading products...</p>}
      <ul>
        {Array.isArray(products) &&
          products.map((product) => (
            <li key={product.id}>
              <img src={product.image} alt={product.name} width="100" />
              <div>
                <strong>{product.name}</strong> - ${product.price}
              </div>
              <button onClick={() => addToCart(product)}>Add to Cart</button>
            </li>
          ))}
      </ul>
    </div>
  );
}

export default Products;

