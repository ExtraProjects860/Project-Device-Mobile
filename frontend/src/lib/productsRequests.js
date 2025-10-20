import { instanceMainApi } from "./axios.js";

/**
 * @param {object} productData
 * @param {string} accessToken
 */
export async function createProductRequest(productData, accessToken) {
  const response = await instanceMainApi.post("/products", productData, {
    headers: {
      Authorization: `Bearer ${accessToken}`,
    },
  });
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
  const response = await instanceMainApi.patch(
    `/products/${productId}`,
    updatedProductData,
    {
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    },
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
  const response = await instanceMainApi.get(
    `/products?itemsPerPage=${itemsPerPage}&currentPage=${currentPage}`,
    {
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    },
  );

  return response.data || [];
}
