import axios from "axios";
import { env } from "./env.js"

export const instanceMainApi = axios.create({
  baseURL: env.axiosURL,
});

instanceMainApi.defaults.headers.post["Content-Type"] = "application/json";
instanceMainApi.defaults.timeout = 15000;
instanceMainApi.defaults.withCredentials = true;
