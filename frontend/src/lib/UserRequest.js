import { instanceMainApi } from "./axios.js";

export async function getUsersRequest(itemsPerPage = 20, currentPage = 1) {
  const response = await instanceMainApi.get(
    `/users?itemsPerPage=${itemsPerPage}&currentPage=${currentPage}`,
  );

  return response.data || [];
}

/**
 * @param {object} userData
 * @param {string} accessToken
 */
export async function createUserRequest(userData, accessToken) {
  const response = await instanceMainApi.post("/user", userData, {
    Authorization: `Bearer ${accessToken}`,
  });
  return response.data;
}

/**
 * @param {string} accessToken
 */
export async function getInfoUserRequest(accessToken) {
  const response = await instanceMainApi.get(`/user`, {
    headers: {
      Authorization: `Bearer ${accessToken}`,
    },
  });
  return response.data;
}

/**
 * @param {string} userId
 * @param {object} updatedUserData
 * @param {string} accessToken
 */
export async function updateUserRequest(userId, updatedUserData, accessToken) {
  const response = await instanceMainApi.patch("/user", updatedUserData, {
    params: { id: userId },
    headers: {
      Authorization: `Bearer ${accessToken}`,
    },
  });
  return response.data;
}
