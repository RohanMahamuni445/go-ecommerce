import React, { useEffect, useState } from "react";
import axios from "axios";

function Products() {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    axios.get("http://localhost:8080/products")
      .then((response) => {
        setProducts(response.data);
      })
      .catch((error) => {
        console.error("Failed to fetch products:", error);
      });
  }, []);

  const addToCart = (product) => {
    axios.post("http://localhost:8080/cart", { 
        product_id: product.id,   // ✅ Fixed key name
        quantity: 1               // ✅ Added quantity
      })
      .then(() => alert(`${product.name} added to cart!`))
      .catch((error) => {
        console.error("Failed to add to cart:", error);
        alert("Failed to add to cart");
      });
  };

  return (
    <div>
      <h2>Products</h2>
      <ul>
        {products.map((product) => (
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

