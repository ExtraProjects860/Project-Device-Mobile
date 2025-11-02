import { requestPost } from "./axios.js";

export async function requestToken(email) {
  const response = await requestPost(`/auth/request-token?email=${email}` );
  return response.data;
}
{/*AINDA NAO SEI SE FUNCIONA*/}
export async function resetPasswordRequest(email, token, newPassword) {
  const response = await requestPost(`/auth/reset-password?email=${email}&token=${token}`,
      {
        new_password: newPassword,
      } );
  return response.data;
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
