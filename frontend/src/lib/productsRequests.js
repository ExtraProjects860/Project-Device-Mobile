import { requestPost, requestGet, requestPatch } from "./axios.js";

/**
 * @param {object} productData
 * @param {string} accessToken
 */
export async function createProductRequest(productData, accessToken) {
  const response = await requestPost("/products", productData, accessToken);
  return response.data;
}

/**
 * @param {string} productId
 * @param {object} updatedProductData
 * @param {string} accessToken
 */
export async function updateProductRequest(
  productId,
  updatedProductData,
  accessToken,
) {
  const response = await requestPatch(
    `/products/${productId}`,
    updatedProductData,
    accessToken,
  );
  return response.data;
}

/**
 * @param {number} itemsPerPage
 * @param {number} currentPage
 * @param {string} accessToken
 */
export async function getProductsRequest(
  itemsPerPage = 20,
  currentPage = 1,
  accessToken,
) {
  const response = await requestGet(
    `/products?itemsPerPage=${itemsPerPage}&currentPage=${currentPage}`,
    accessToken,
  );
  return response.data || [];
}
