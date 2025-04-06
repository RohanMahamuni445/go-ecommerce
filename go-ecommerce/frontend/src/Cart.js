import React, { useEffect, useState } from "react";
import axios from "axios";

function Cart() {
  const [cart, setCart] = useState([]);
  const [total, setTotal] = useState(0);

  useEffect(() => {
    axios
      .get("http://localhost:8080/cart")
      .then((response) => {
        if (response.data && Array.isArray(response.data)) {
          setCart(response.data);
          const totalPrice = response.data.reduce(
            (sum, item) => sum + (item.price * item.quantity || 0),
            0
          );
          setTotal(totalPrice);
        } else {
          setCart([]);
          setTotal(0);
        }
      })
      .catch((error) => {
        console.error("Error fetching cart:", error);
        setCart([]);
        setTotal(0);
      });
  }, []);

  return (
    <div>
      <h2>Your Cart</h2>

      {cart.length === 0 ? (
        <p>Your cart is empty</p>
      ) : (
        <ul>
          {cart.map((item) => (
            <li key={item.id}>
              <strong>{item.name}</strong> - {item.quantity} pcs - $
              {(item.price * item.quantity || 0).toFixed(2)}
            </li>
          ))}
        </ul>
      )}

      <h3>Total: ${total.toFixed(2)}</h3>
    </div>
  );
}

export default Cart;

