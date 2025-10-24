import { instanceMainApi } from "./axios.js";

export async function addInWishListRequest() {
  return;
}

export async function deleteWishListRequest() {
  return;
}

/**
 * @param {number} itemsPerPage
 * @param {number} currentPage
 * @param {string} accessToken
 */
export async function getItemsWishListRequest(
  itemsPerPage = 20,
  currentPage = 1,
  accessToken,
) {
  const response = await instanceMainApi.get(
    `/wishlist?itemsPerPage=${itemsPerPage}&currentPage=${currentPage}`,
    {
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    },
  );

  return response.data || [];
}
