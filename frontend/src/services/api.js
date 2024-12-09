import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:1000", 
  headers: {
    "Content-Type": "application/json",
  },
});

console.log("API base URL:", api.defaults.baseURL);

// Add a request interceptor to include the token if it exists
api.interceptors.request.use((config) => {
  const token = localStorage.getItem("token");
  if (token) {
    console.log("Token found, adding to request header:", token);
    config.headers.Authorization = `Bearer ${token}`;
  } else {
    console.log("No token found in localStorage.");
  }
  console.log("Outgoing request config:", config);
  return config;
}, (error) => {
  console.error("Error in request interceptor:", error);
  return Promise.reject(error);
});

// Add a response interceptor to debug responses
api.interceptors.response.use(
  (response) => {
    console.log("API response:", response);
    return response;
  },
  (error) => {
    console.error("API error response:", error.response);
    return Promise.reject(error);
  }
);

export default api;
