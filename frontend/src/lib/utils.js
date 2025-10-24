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
  const dateConvert = new Date(date);
  return dateConvert.toLocaleDateString("pt-BR");
}

/**
 * @param {string} cnpj
 * @returns {string}
 */
export function formatCNPJ(cnpj) {
  return;
}

/**
 * @param {number} value 
 * @param {number} percentage 
 * @returns {number}
 */
export function formatPromotion(value, percentage) {
  const each = (value * percentage);
  const result = value - each;
  return result.toFixed(2);
}
