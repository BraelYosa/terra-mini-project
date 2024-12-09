import React, { useEffect, useState } from "react";
import api from "../services/api";

function DataList() {
  const [data, setData] = useState([]);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await api.get("/data"); // Fetch data from backend
        console.log("Fetched data:", response.data); // Debugging log
  
        // Check if response data has a nested "data" field
        if (response.data && Array.isArray(response.data.data)) {
          setData(response.data.data); // Use the array inside "data"
        } else {
          console.error("Unexpected response format:", response.data);
          setError("Unexpected response format. Please contact support.");
        }
      } catch (err) {
        console.error("Error fetching data:", err);
        setError("Failed to fetch data");
      }
    };
  
    fetchData();
  }, []);
  
  

  if (error) return <p style={{ color: "red" }}>{error}</p>;

  return (
    <div>
      <h2>Data List</h2>
      {data.length === 0 ? (
        <p>No data available.</p> 
      ) : (
        <ul>
          {data.map((item, index) => (
            <li key={index}>{item.name}</li> 
          ))}
        </ul>
      )}
    </div>
  );
}

export default DataList;
