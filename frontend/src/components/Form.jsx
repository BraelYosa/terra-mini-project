import React, { useState } from "react";
import api from "../services/api";

function Form() {
  const [name, setName] = useState("");
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(null);
    setSuccess(false);

    try {
      await api.post("/data", { name }); // Post data to backend
      setSuccess(true); // Display success message
      setName(""); // Clear the form
    } catch (err) {
      setError("Failed to submit data");
    }
  };

  return (
    <form className="form" onSubmit={handleSubmit}>
      {error && <p style={{ color: "red" }}>{error}</p>}
      {success && <p style={{ color: "green" }}>Data submitted successfully!</p>}
      <div>
        <label>Name:</label>
        <input
          type="text"
          placeholder="Enter name"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
      </div>
      <button type="submit">Submit</button>
    </form>
  );
}

export default Form;
