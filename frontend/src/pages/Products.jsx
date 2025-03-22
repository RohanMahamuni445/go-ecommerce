import { useEffect, useState } from "react";

const Products = () => {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    fetch("http://localhost:5000/api/products") // Replace with your actual Go backend URL
      .then((res) => res.json())
      .then((data) => setProducts(data))
      .catch((err) => console.error("Error fetching products:", err));
  }, []);

  return (
    <div style={{ padding: "20px" }}>
      <h2>🛍 Products</h2>
      <div style={{ display: "grid", gridTemplateColumns: "repeat(3, 1fr)", gap: "20px" }}>
        {products.length > 0 ? (
          products.map((product) => (
            <div key={product.id} style={styles.card}>
              <img src={product.image} alt={product.name} style={styles.image} />
              <h3>{product.name}</h3>
              <p>${product.price}</p>
              <button style={styles.button}>Add to Cart</button>
            </div>
          ))
        ) : (
          <p>Loading products...</p>
        )}
      </div>
    </div>
  );
};

const styles = {
  card: { border: "1px solid #ddd", padding: "10px", borderRadius: "8px", textAlign: "center" },
  image: { width: "100%", height: "150px", objectFit: "cover" },
  button: { background: "#333", color: "white", padding: "5px 10px", border: "none", cursor: "pointer" },
};

export default Products;

