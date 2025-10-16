import { instanceMainApi } from "./axios.js";

export async function getUsersRequest(itemsPerPage = 20, currentPage = 1) {
  const response = await instanceMainApi.get(
    `/users?itemsPerPage=${itemsPerPage}&currentPage=${currentPage}`,
  );

  return response.data || [];
}

export async function createUserRequest(userData) {
  const response = await instanceMainApi.post("/user", userData);
  return response.data;
}

export async function getInfoUserRequest(userId) {
  const response = await instanceMainApi.get(`/user/${userId}`);
  return response.data;
}

/**
 * @param {string} userId
 * @param {object} updatedUserData
 */
export async function updateUserRequest(userId, updatedUserData) {
  const response = await instanceMainai.patch(
    `/user/${userId}`,
    updatedUserData,
  );
  return response.data;
}
