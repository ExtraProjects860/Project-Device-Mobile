import {
  configsToPagination,
  requestGet,
  requestPatch,
  requestPost,
} from "./axios.js";

/**
 * @param {object} userData
 * @param {object} image
 * @param {string} accessToken
 */
export async function createUserRequest(userData, image = null, accessToken) {
  const formData = new FormData();

  formData.append("data", JSON.stringify(userData));

  if (image) {
    formData.append("image", {
      uri: image.uri,
      name: image.fileName || image.uri.split("/").pop(),
      type: image.mimeType || "image/jpeg",
    });
  }

  const response = await requestPost("/user", formData, accessToken);
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
  const formData = new FormData();
  let photoAssetObject = null;

  if (updatedUserData.hasOwnProperty("photo_asset")) {
    photoAssetObject = updatedUserData.photo_asset;
    delete updatedUserData.photo_asset;
  }

  formData.append("data", JSON.stringify(updatedUserData));

  if (photoAssetObject) {
    formData.append("image", {
      uri: photoAssetObject.uri,
      name: photoAssetObject.fileName || photoAssetObject.uri.split("/").pop(),
      type: photoAssetObject.mimeType || "image/jpeg",
    });
  }

  const response = await requestPatch("/user", formData, accessToken, {
    params: { id: userId },
  });
  return response.data;
}

/**
 * @param {number} itemsPerPage
 * @param {number} currentPage
 * @param {string} accessToken
 * @param {string} searchFilter
 */
export async function getUsersRequest(
  itemsPerPage = configsToPagination.itemsPerPage,
  currentPage = configsToPagination.currentPage,
  accessToken,
  searchFilter = "",
  itemsOrder = "DESC",
) {
  let url = `/users?itemsPerPage=${itemsPerPage}&currentPage=${currentPage}&itemsOrder=${itemsOrder}`;

  if (searchFilter && searchFilter.trim() !== "") {
    url += `&searchFilter=${encodeURIComponent(searchFilter)}`;
  }
  const response = await requestGet(url, accessToken);
  return response.data || [];
}
