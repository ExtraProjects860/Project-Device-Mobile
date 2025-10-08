import axios from "axios";

// TODO não esquecer de mudar sempre a URL do axios quando for utilizar
export const instanceMainApi = axios.create({
  baseURL: "https://thermogenous-jimmie-manly.ngrok-free.dev/api/v1",
});

instanceMainApi.defaults.headers.post["Content-Type"] = "application/json";
instanceMainApi.defaults.timeout = 15000;
instanceMainApi.defaults.withCredentials = true;

export const instanceNoticesApi = axios.create({
  baseURL: "",
});
