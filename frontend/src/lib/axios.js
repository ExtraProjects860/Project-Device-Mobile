import axios from "axios";
import { env } from "./env.js";

export const instanceMainApi = axios.create({
  baseURL: env.axiosURL,
  timeout: 15000,
  withCredentials: true,
  headers: {
    "Content-Type": "application/json",
  },
});

export const configsToPagination = Object.freeze({
  itemsPerPage: 10,
  currentPage: 1,
});

const authorizedHeader = (accessToken) => ({
  Authorization: `Bearer ${accessToken}`,
});

/**
 * Configura interceptors globais (ex: tratamento de 401)
 * @param {Function} logoutCallback
 */
export const setupAxiosInterceptors = (logoutCallback) => {
  instanceMainApi.interceptors.response.use(
    (response) => {
      return response;
    },

    (error) => {
      if (!error.response) return Promise.reject(error);

      if (error.response.status === 401) {
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

/**
 * Função genérica para requisições
 * @param {'get'|'post'|'patch'|'put'|'delete'} method
 * @param {string} url
 * @param {object} [data]
 * @param {string} [accessToken]
 * @param {object} [config]
 */
const baseRequest = async (
  method,
  url,
  data = {},
  accessToken,
  config = {},
) => {
  const headers = accessToken
    ? { ...authorizedHeader(accessToken), ...(config.headers || {}) }
    : { ...(config.headers || {}) };

  if (data instanceof FormData) {
    headers["Content-Type"] = "multipart/form-data";
  } else {
    if (!headers["Content-Type"]) {
      headers["Content-Type"] = "application/json";
    }
  }

  const finalConfig = { ...config, headers };

  switch (method.toLowerCase()) {
    case "get":
      return instanceMainApi.get(url, finalConfig);
    case "post":
      return instanceMainApi.post(url, data, finalConfig);
    case "patch":
      return instanceMainApi.patch(url, data, finalConfig);
    case "put":
      return instanceMainApi.put(url, data, finalConfig);
    case "delete":
      return instanceMainApi.delete(url, finalConfig);
    default:
      throw new Error(`Método HTTP inválido: ${method}`);
  }
};

export const requestGet = async (url, accessToken, config = {}) =>
  await baseRequest("get", url, null, accessToken, config);

export const requestPost = async (url, data, accessToken, config = {}) =>
  await baseRequest("post", url, data, accessToken, config);

export const requestPatch = async (url, data, accessToken, config = {}) =>
  await baseRequest("patch", url, data, accessToken, config);

export const requestPut = async (url, data, accessToken, config = {}) =>
  await baseRequest("put", url, data, accessToken, config);

export const requestDelete = async (url, accessToken, config = {}) =>
  await baseRequest("delete", url, null, accessToken, config);
