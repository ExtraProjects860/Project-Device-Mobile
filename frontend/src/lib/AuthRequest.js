import { instanceMainApi } from "./axios.js";

export async function requestToken() {
  return;
}

export async function resetPasswordRequest() {
  return;
}

/**
 * @param {object} userData
 */
export async function loginRequest(userData) {
  const response = await instanceMainApi.post("/auth/login", userData);
  const accessToken = response.data?.access_token;

  return accessToken;
}

export async function refreshTokenRequest() {
  return;
}
