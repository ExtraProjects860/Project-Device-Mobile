import {
  configsToPagination,
  requestGet,
  requestPatch,
  requestPost,
} from "./axios.js";

/**
 * @param {object} userData
 * @param {string} accessToken
 */
export async function createUserRequest(userData, accessToken) {
  const response = await requestPost("/user", userData, accessToken);
  return response.data;
}

/**
 * @param {string} accessToken
 */
export async function getInfoUserRequest(accessToken) {
  const response = await requestGet(`/user`, accessToken);
  return response.data;
}

/**
 * @param {string} userId
 * @param {object} updatedUserData
 * @param {string} accessToken
 */
export async function updateUserRequest(userId, updatedUserData, accessToken) {
  const response = await requestPatch("/user", updatedUserData, accessToken, {
    params: { id: userId },
  });
  return response.data;
}

/**
 * @param {number} itemsPerPage
 * @param {number} currentPage
 * @param {string} accessToken
 */
export async function getUsersRequest(
  itemsPerPage = configsToPagination.itemsPerPage,
  currentPage = configsToPagination.currentPage,
  accessToken,
) {
  const response = await requestGet(
    `/users?itemsPerPage=${itemsPerPage}&currentPage=${currentPage}`,
    accessToken,
  );
  return response.data || [];
}
