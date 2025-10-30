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
 * @param {string} searchFilter
 */
export async function getItemsWishListRequest(
  itemsPerPage = configsToPagination.itemsPerPage,
  currentPage = configsToPagination.currentPage,
  accessToken,
  searchFilter,
  itemsOrder = "ASC",
) {
  let url = `/wishlist?itemsPerPage=${itemsPerPage}&currentPage=${currentPage}&itemsOrder=${itemsOrder}`;

  if (searchFilter) {
    url += `&searchFilter=${searchFilter}`;
  }

  const response = await requestGet(url, accessToken);
  return response.data || [];
}
