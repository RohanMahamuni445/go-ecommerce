import React from "react";
import "./styles.css";
import { BrowserRouter as Router, Route, Routes, Link } from "react-router-dom";
import Products from "./Products";
import Cart from "./Cart";
import Auth from "./Auth";

function App() {
  return (
    <Router>
      <div>
        <nav>
          <ul>
            <li><Link to="/products">Products</Link></li>
            <li><Link to="/cart">Cart</Link></li>
            <li><Link to="/auth">Login / Register</Link></li>
          </ul>
        </nav>

        <Routes>
          <Route path="/products" element={<Products />} />
          <Route path="/cart" element={<Cart />} />
          <Route path="/auth" element={<Auth />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;

