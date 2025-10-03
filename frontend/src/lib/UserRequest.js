import { instanceMainApi } from "./axios.js";

export async function getUsersRequest(itemsPerPage = 20, currentPage = 1) {
  const response = await instanceMainApi.get(
    `/users?itemsPerPage=${itemsPerPage}&currentPage=${currentPage}`,
  );

  return response.data || [];
}

export async function createUserRequest() {
  return;
}

export async function getInfoUserRequest() {
  return;
}

export async function updateUserRequest() {
  return;
}
