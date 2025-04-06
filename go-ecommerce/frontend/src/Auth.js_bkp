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

  const handleRegister = () => {
    axios.post("http://localhost:8080/register", { username, password })
      .then(response => alert(response.data))
      .catch(error => alert("Error registering"));
  };

  const handleLogin = () => {
    axios.post("http://localhost:8080/login", { username, password })
      .then(response => {
        localStorage.setItem("user", username);
        setLoggedInUser(username);
      })
      .catch(error => alert("Invalid login"));
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
        <>
          <input type="text" placeholder="Username" onChange={(e) => setUsername(e.target.value)} />
          <input type="password" placeholder="Password" onChange={(e) => setPassword(e.target.value)} />
          <button onClick={handleRegister}>Register</button>
          <button onClick={handleLogin}>Login</button>
        </>
      )}
    </div>
  );
}

export default Auth;

