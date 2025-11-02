import { requestPost, requestGet, requestPatch } from "./axios.js";

/**
 * @param {object} productData 
 * @param {string | null} photoUri 
 * @param {string} accessToken
 */
export async function createProductRequest(productData, photoUri, accessToken) {
  const formData = new FormData();
  formData.append("data", JSON.stringify(productData)); 

  if (photoUri) {
    const filename = photoUri.split("/").pop();
    const fileType = filename.endsWith(".png") ? "image/png" : "image/jpeg";
    formData.append("image", { 
      uri: photoUri,
      name: filename,
      type: fileType,
    });
  }
  const response = await requestPost("/product", formData, accessToken);
  return response.data;
}

/**
 * @param {string} productId
 * @param {object} updatedProductData
 * @param {string | null} photoUri 
 * @param {string} accessToken
 */
export async function updateProductRequest(productId, updatedProductData, photoUri, accessToken) {
  const formData = new FormData();
  formData.append("data", JSON.stringify(updatedProductData));

  if (photoUri) {
    const filename = photoUri.split("/").pop();
    const fileType = filename.endsWith(".png") ? "image/png" : "image/jpeg";
    formData.append("image", {
      uri: photoUri,
      name: filename,
      type: fileType,
    });
  }

  const response = await requestPatch(
    `/product?id=${productId}`,
    formData,
    accessToken
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
  searchFilter = "",
  itemsOrder = "DESC",
) {
  let url = `/products?itemsPerPage=${itemsPerPage}&currentPage=${currentPage}&itemsOrder=${itemsOrder}`;

  if (searchFilter && searchFilter.trim() !== "") {
    url += `&searchFilter=${encodeURIComponent(searchFilter)}`;
  }
  const response = await requestGet(url, accessToken);
  return response.data || [];
}