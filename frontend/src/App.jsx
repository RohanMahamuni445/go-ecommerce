import { useEffect, useState } from "react";
import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/js/bootstrap.bundle.min";
import Navbar from "./components/Navbar";
import "./App.css";
import Cart from "./pages/Cart";  
import Login from "./pages/Login"; 

// ✅ Set API base URL dynamically
const API_BASE_URL = process.env.REACT_APP_API_URL || "http://localhost:8086";

function ProductList() {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    fetch(`${API_BASE_URL}/products`)
      .then((response) => response.json())
      .then((data) => setProducts(data))
      .catch((error) => console.error("Error fetching products:", error));
  }, []);

  return (
    <div className="container mt-4">
      <h1 className="text-center">Product List</h1>
      <div className="row justify-content-center">
        {products.map((product) => (
          <div key={product.id} className="col-md-4 mb-4">
            <div className="card shadow-sm text-center">
              <img
                src={product.image_url}
                alt={product.name}
                className="card-img-top"
                style={{ height: "200px", objectFit: "cover" }}
              />
              <div className="card-body">
                <h3 className="card-title">{product.name}</h3>
                <p className="card-text">{product.description}</p>
                <p><strong>Price:</strong> ${product.price}</p>
                <p><strong>Stock:</strong> {product.stock}</p>
                <button className="btn btn-primary">Add to Cart</button>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

// ✅ Protected Route Component
const ProtectedRoute = ({ children }) => {
  const token = localStorage.getItem("token");
  return token ? children : <Navigate to="/login" />;
};

function App() {
  return (
    <Router>
      <Navbar />
      <Routes>
        <Route path="/" element={<ProductList />} />
        <Route path="/cart" element={<ProtectedRoute><Cart /></ProtectedRoute>} />
        <Route path="/login" element={<Login />} />
      </Routes>
    </Router>
  );
}

export default App;

