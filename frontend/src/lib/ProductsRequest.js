import { instanceMainApi } from "./axios.js";

/**
 * @param {object} productData
 */
export async function createProductRequest(productData) {
  const response = await instanceMainApi.post("/products", productData);
  return response.data;
}

/**
 * @param {string} productId
 * @param {object} updatedProductData
 */
export async function updateProductRequest(productId, updatedProductData) {
  const response = await instanceMainApi.patch("/product", updatedProductData, {
    params: { id: productId },
  });
  return response.data;
}

/**
 * @param {number} itemsPerPage
 * @param {number} currentPage
 */
export async function getProductsRequest(itemsPerPage = 20, currentPage = 1) {
  const response = await instanceMainApi.get(
    `/products?itemsPerPage=${itemsPerPage}&currentPage=${currentPage}`,
  );

  return response.data || [];
}
