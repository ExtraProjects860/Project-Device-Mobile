import {
  configsToPagination,
  requestDelete,
  requestGet,
  requestPost,
} from "./axios";

/**
 * @param {number} productId 
 * @param {string} accessToken 
 */
export async function addInWishListRequest(productId, accessToken) {
  const response = await requestPost(
    `wishlist?product_id=${productId}`,
    {},
    accessToken,
  );

  return response.data;
}

/**
 * @param {number} productId 
 * @param {string} accessToken 
 */
export async function deleteWishListRequest(productId, accessToken) {
  const response = await requestDelete(
    `wishlist?product_id=${productId}`,
    accessToken,
  );

  return response.data;
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
