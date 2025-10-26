import { requestPost } from "./axios.js";

export async function requestToken() {
  return;
}

export async function resetPasswordRequest() {
  return;
}

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
