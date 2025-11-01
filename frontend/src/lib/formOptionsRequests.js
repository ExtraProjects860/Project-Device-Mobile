import { requestGet, configsToPagination } from "./axios.js";

/**
 * Busca a lista completa de Funções (Roles)
 * @param {string} accessToken
 * @param {number} itemsPerPage
 * @param {number} currentPage
 */
export async function getRolesRequest(
  accessToken,
  itemsPerPage = configsToPagination.itemsPerPage,
  currentPage = configsToPagination.currentPage,
) {
  const url = `/roles?itemsPerPage=${itemsPerPage}&currentPage=${currentPage}`;
  const response = await requestGet(url, accessToken);

  return response.data.data || [];
}

/**
 * Busca a lista completa de Empresas (Enterprises)
 * @param {string} accessToken
 * @param {number} itemsPerPage
 * @param {number} currentPage
 */
export async function getEnterprisesRequest(
  accessToken,
  itemsPerPage = configsToPagination.itemsPerPage,
  currentPage = configsToPagination.currentPage,
) {
  const url = `/enterprises?itemsPerPage=${itemsPerPage}&currentPage=${currentPage}`;
  const response = await requestGet(url, accessToken);

  return response.data.data || [];
}
