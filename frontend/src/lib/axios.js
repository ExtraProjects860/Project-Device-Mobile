import axios from "axios";
import { env } from "./env.js";

export const instanceMainApi = axios.create({
  baseURL: env.axiosURL,
});

instanceMainApi.defaults.headers.post["Content-Type"] = "application/json";
instanceMainApi.defaults.timeout = 15000;
instanceMainApi.defaults.withCredentials = true;

/**
 *
 * @param {Function} logoutCallback
 */
export const setupAxiosInterceptors = (logoutCallback) => {
  instanceMainApi.interceptors.response.use(
    (response) => {
      return response;
    },

    (error) => {
      if (error.response && error.response.status === 401) {
        console.warn(
          "Axios Interceptor: Token inválido ou expirado (401). Executando logout.",
        );
        logoutCallback();
        return new Promise(() => {});
      }
      return Promise.reject(error);
    },
  );
};
