import axios from "axios";

const instance = axios.create({
  baseURL: "http://localhost:5050",
});

instance.defaults.headers.post["Content-Type"] = "application/json";
instance.defaults.timeout = 15
instance.defaults.withCredentials = true;
