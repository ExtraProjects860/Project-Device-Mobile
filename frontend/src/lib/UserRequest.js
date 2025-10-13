import { instanceMainApi } from "./axios.js";

export async function getUsersRequest(itemsPerPage = 20, currentPage = 1) {
  const response = await instanceMainApi.get(
    `/users?itemsPerPage=${itemsPerPage}&currentPage=${currentPage}`,
  );

  return response.data || [];
}

export async function createUserRequest(userData) {
  const response = await instanceMainApi.post("/users", userData);
  return response.data;
}

export async function getInfoUserRequest(userId) {
  const response = await instanceMainApi.get(`/users/${userId}`);
  return response.data;
}

/**
 * @param {string} userId
 * @param {object} updatedUserData
 * @param {string} [updatedUserData.name]
 * @param {string} [updatedUserData.register_number]
 * @param {string} [updatedUserData.email]
 */
export async function updateUserRequest(userId, updatedUserData) {
  const response = await instanceMainai.patch(
    `/users/${userId}`,
    updatedUserData,
  );
  return response.data;
}
