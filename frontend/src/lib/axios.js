import axios from "axios";

export class ApiDefault {
  constructor(axiosInstance) {
    this.axiosInstance = axiosInstance;
  }
}

export const instanceMainApi = axios.create({
  baseURL: "http://localhost:5050",
});

instanceMainApi.defaults.headers.post["Content-Type"] = "application/json";
instanceMainApi.defaults.timeout = 15;
instanceMainApi.defaults.withCredentials = true;

export const instanceNoticesApi = axios.create({
  baseURL: "",
});
