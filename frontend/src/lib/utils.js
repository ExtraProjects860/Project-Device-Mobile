/**
 * @param {string} cpf
 * @returns {string}
 */
export function formatCPF(cpf) {
  if (!cpf) return "";
  return cpf.replace(/^(\d{3})(\d{3})(\d{3})(\d{2})$/, "$1.$2.$3-$4");
}

/**
 * @param {string} date
 * @returns {string}
 */
export function formatDate(date) {
  return;
}

/**
 * @param {string} datetime
 * @returns {string}
 */
export function formatDateTime(datetime) {
  return;
}

/**
 * @param {string} cnpj
 * @returns {string}
 */
export function formatCNPJ(cnpj) {
  return;
}
