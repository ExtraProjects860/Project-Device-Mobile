import { requestPost, requestGet, requestPatch } from "./axios.js";

/**
 * @param {object} productData
 * @param {object | null} photoAsset
 * @param {string} accessToken
 */
export async function createProductRequest(
  productData,
  photoAsset = null,
  accessToken,
) {
  const formData = new FormData();
  formData.append("data", JSON.stringify(productData));

  if (photoAsset) {
    const filename = photoAsset.fileName || photoAsset.uri.split("/").pop();
    const fileType =
      photoAsset.mimeType ||
      (filename.endsWith(".png") ? "image/png" : "image/jpeg");
    formData.append("image", {
      uri: photoAsset.uri,
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
 * @param {object | null} photoAsset
 * @param {string} accessToken
 */
export async function updateProductRequest(
  productId,
  updatedProductData,
  photoAsset = null,
  accessToken,
) {
  const formData = new FormData();
  formData.append("data", JSON.stringify(updatedProductData));

  if (photoAsset) {
    const filename = photoAsset.fileName || photoAsset.uri.split("/").pop();
    const fileType =
      photoAsset.mimeType ||
      (filename.endsWith(".png") ? "image/png" : "image/jpeg");
    formData.append("image", {
      uri: photoAsset.uri,
      name: filename,
      type: fileType,
    });
  }

  const response = await requestPatch(
    `/product?id=${productId}`,
    formData,
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
