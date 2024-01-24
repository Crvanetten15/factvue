// src/services/ApiService.js
import axios from "axios";

const API_URL = "http://localhost:8080"; // Make sure this is set in your .env file

export default {
  getRootMessage() {
    return axios.get(`${API_URL}/`); // Assuming the root path gives the desired response
  },

  // Add other API methods as needed
};
