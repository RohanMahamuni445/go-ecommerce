import { useEffect, useState } from "react";

function Cart() {
  const [cartItems, setCartItems] = useState([]);
  const [error, setError] = useState(null);
  const userId = 1; // Change this dynamically if needed

  useEffect(() => {
    const fetchCart = async () => {
      try {
        const response = await fetch(`http://localhost:8086/cart/${userId}`);
        
        if (!response.ok) {
          throw new Error(`HTTP error! Status: ${response.status}`);
        }

        const text = await response.text(); // Read raw response first
        console.log("Raw API Response:", text); // Debugging output

        const jsonData = JSON.parse(text.trim()); // Ensure valid JSON
        setCartItems(jsonData);
      } catch (err) {
        console.error("Error fetching cart:", err);
        setError("Failed to fetch cart items. Please try again.");
      }
    };

    fetchCart();
  }, [userId]);

  return (
    <div className="container mt-4">
      <h1 className="text-center">Shopping Cart</h1>
      
      {error ? (
        <p className="text-danger text-center">{error}</p>
      ) : cartItems.length > 0 ? (
        cartItems.map((item) => (
          <div key={item.id} className="card mb-3 p-3">
            <h3>Product ID: {item.product_id}</h3>
            <p>Price: ${item.price}</p>
            <p>Quantity: {item.quantity}</p>
          </div>
        ))
      ) : (
        <p className="text-center">Your cart is empty</p>
      )}
    </div>
  );
}

export default Cart;

