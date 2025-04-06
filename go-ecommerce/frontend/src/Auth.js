import React, { useState, useEffect } from "react";
import axios from "axios";

function Auth() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [loggedInUser, setLoggedInUser] = useState(null);

  useEffect(() => {
    const user = localStorage.getItem("user");
    if (user) {
      setLoggedInUser(user);
    }
  }, []);

  const handleRegister = async () => {
    try {
      const response = await axios.post("http://localhost:8080/register", {
        username,
        password,
      });
      alert(response.data);
    } catch (error) {
      console.error("Registration failed:", error);
      alert("Error registering");
    }
  };

  const handleLogin = async () => {
    try {
      const response = await axios.post("http://localhost:8080/login", {
        username,
        password,
      });
      localStorage.setItem("user", username);
      setLoggedInUser(username);
      alert("Login successful!");
    } catch (error) {
      console.error("Login failed:", error);
      alert("Invalid login");
    }
  };

  const handleLogout = () => {
    localStorage.removeItem("user");
    setLoggedInUser(null);
  };

  return (
    <div>
      <h2>Login / Register</h2>

      {loggedInUser ? (
        <div>
          <p>Welcome, {loggedInUser}!</p>
          <button onClick={handleLogout}>Logout</button>
        </div>
      ) : (
        <div>
          <input
            type="text"
            placeholder="Username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            style={{ marginRight: "10px" }}
          />
          <input
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            style={{ marginRight: "10px" }}
          />
          <button onClick={handleRegister} style={{ marginRight: "10px" }}>
            Register
          </button>
          <button onClick={handleLogin}>Login</button>
        </div>
      )}
    </div>
  );
}

export default Auth;

