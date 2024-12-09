import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../services/api";

function Login() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  const handleLogin = async (e) => {
    e.preventDefault();
    try {
      console.log("Login credentials:", { user_mail: username, user_pass: password }); // Adjusted keys
      const response = await api.post("/login", {
        user_mail: username, // Correct key for backend
        user_pass: password, // Correct key for backend
      });
  
      console.log("Login successful:", response.data);
  
      if (response.data.token) {
        localStorage.setItem("token", response.data.token);
        console.log("Redirecting to dashboard...");
        navigate("/dashboard"); // Redirect to the dashboard
      } else {
        setError("Login response is missing a token.");
      }
    } catch (err) {
      console.error("Login error:", err.response?.data || err.message);
      setError(err.response?.data?.message || "Login failed");
    }
  };    
  
  
  return (
    <div className="container">
      <h2>Login</h2>
      {error && <p style={{ color: "red" }}>{error}</p>}
      <form onSubmit={handleLogin}>
        <div>
          <label>Email:</label>
          <input
            type="email"
            placeholder="Enter email address"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
        </div>
        <div>
          <label>Password:</label>
          <input
            type="password"
            placeholder="Enter password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>
        <button type="submit">Login</button>
      </form>
    </div>
  );
}

export default Login;
