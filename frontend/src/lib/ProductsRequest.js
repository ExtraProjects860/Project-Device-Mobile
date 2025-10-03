import { instanceMainApi } from "./axios.js";

export async function createProductRequest() {
  return;
}

export async function updateProductRequest() {
  return;
}

export async function getProductsRequest(itemsPerPage = 20, currentPage = 1) {
  const response = await instanceMainApi.get(
    `/products?itemsPerPage=${itemsPerPage}&currentPage=${currentPage}`,
  );

  return response.data || [];
}
