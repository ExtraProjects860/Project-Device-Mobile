import { configsToPagination, requestGet } from "./axios";

export async function addInWishListRequest(userId, productId, accessToken) {
  return;
}

export async function deleteWishListRequest(userId, productId, accessToken) {
  return;
}

/**
 * @param {number} itemsPerPage
 * @param {number} currentPage
 * @param {string} accessToken
 */
export async function getItemsWishListRequest(
  itemsPerPage = configsToPagination.itemsPerPage,
  currentPage = configsToPagination.currentPage,
  accessToken,
) {
  const response = await requestGet(
    `/wishlist?itemsPerPage=${itemsPerPage}&currentPage=${currentPage}`,
    accessToken,
  );
  return response.data || [];
}
