import React, { useEffect, useState } from "react";
import axios from "axios";

function Cart() {
  const [cart, setCart] = useState([]); // Default to empty array to prevent errors
  const [total, setTotal] = useState(0);

  useEffect(() => {
    axios.get("http://localhost:8080/cart")
      .then((response) => {
        if (response.data && Array.isArray(response.data)) {
          setCart(response.data);
          let totalPrice = response.data.reduce((sum, item) => sum + (item.price * item.quantity || 0), 0);
          setTotal(totalPrice);
        } else {
          setCart([]); // Set an empty cart if the response is invalid
          setTotal(0);
        }
      })
      .catch((error) => {
        console.error("Error fetching cart:", error);
        setCart([]); // Handle API failure gracefully
        setTotal(0);
      });
  }, []);

  return (
    <div>
      <h2>Cart</h2>

      {cart.length === 0 ? (
        <p>Your cart is empty</p>
      ) : (
        <ul>
          {cart.map((item) => (
            <li key={item.id}>
              {item.name} - {item.quantity} pcs - ${item.price * item.quantity || 0}
            </li>
          ))}
        </ul>
      )}

      <h3>Total: ${total}</h3>
    </div>
  );
}

export default Cart;

