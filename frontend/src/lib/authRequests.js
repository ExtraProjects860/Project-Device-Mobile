import { requestPost } from "./axios.js";

/**
 * @param {string} email 
 */
export async function requestToken(email) {
  const response = await requestPost(`/auth/request-token?email=${email}`);
  return response.data;
}

/**
 * @param {string} email 
 * @param {string} token 
 * @param {string} newPassword 
 */
export async function resetPasswordRequest(email, token, newPassword) {
  const response = await requestPost(
    `/auth/reset-password?email=${email}&token=${token}`,
    {
      new_password: newPassword,
    },
  );
  return response.data;
}

/**
 * @param {string} newPassword 
 * @param {string} accessToken 
 */
export async function resetPasswordInternalRequest(newPassword, accessToken) {
  const response = await requestPost(
    "/auth/reset-pass-log-in",
    { new_password: newPassword },
    accessToken,
  );
  return response.data;
}

/**
 * @param {object} userData
 */
export async function loginRequest(userData) {
  const response = await requestPost("/auth/login", userData);
  return response.data?.access_token;
}
