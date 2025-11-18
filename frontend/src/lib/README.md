# Visão Geral "lib"

Estrutura de Requisições e Utilitários
Este repositório contém funções para interagir com a API principal, manipulando dados do usuário, produtos, listas de desejos (wishlist), autenticação e fornecendo utilitários de formatação e armazenamento.

# 1. Visão Geral "axios.js"

Configura o cliente HTTP principal e funções genéricas de requisição.

- >Cliente Principal: Exporta instanceMainApi, uma instância configurada do Axios com baseURL de env.axiosURL, timeout de 15 segundos, e withCredentials: true.

- >Paginação Padrão: Exporta configsToPagination, com itemsPerPage: 10 e currentPage: 1.

- >Interceptors: A função setupAxiosInterceptors(logoutCallback) configura um interceptor de resposta para tratar erros HTTP. Especificamente, se a resposta retornar o status 401 (Não Autorizado), ele executa o logoutCallback fornecido, tratando tokens inválidos ou expirados.

- >Requisições Genéricas: Exporta funções de wrapper para métodos HTTP: requestGet, requestPost, requestPatch, requestPut, e requestDelete, todas baseadas em baseRequest.

# 2. Visão Geral "authRequests.js"

- >Contém funções dedicadas à autenticação e recuperação de senha.
- >loginRequest(userData): Realiza o login e retorna o access_token.
- >requestToken(email): Solicita um token de recuperação de senha via e-mail.
- >esetPasswordRequest(email, token, newPassword): Redefine a senha usando o token externo (enviado por e-mail).
- >resetPasswordInternalRequest(newPassword, accessToken): Redefine a senha do usuário que está logado.



# 3. Visão Geral "userRequests.js"

Funções para gerenciar dados de usuários.

- >createUserRequest(userData, image, accessToken): Cria um novo usuário, com suporte ao envio de dados e uma imagem via FormData.
- >getInfoUserRequest(accessToken): Busca as informações do usuário logado.
- >updateUserRequest(userId, updatedUserData, image, accessToken): Atualiza os dados de um usuário específico, também com suporte a upload de imagem via FormData.
- > getUsersRequest(...): Busca uma lista paginada de usuários, permitindo filtros de busca, paginação e ordenação.


# 4. Visão Geral "productsRequests.js"

Funções para gerenciar produtos.

- >createProductRequest(productData, photoUri, accessToken): Cria um novo produto, aceitando dados do produto e uma URI de foto.

- >updateProductRequest(productId, updatedProductData, photoUri, accessToken): Atualiza um produto existente por productId, com suporte para a atualização da foto.

- >getProductsRequest(...): Obtém uma lista paginada de produtos, com opções para filtros, paginação e ordenação.

# 5. Visão Geral "wishListRequests.js"
Funções para gerenciar a Lista de Desejos "Wish List".

- >addInWishListRequest(productId, accessToken): Adiciona um produto à lista de desejos usando o ID do produto.

- >deleteWishListRequest "productId", "accessToken": Remove um produto da lista de desejos.

- >getItemsWishListRequest: Busca a lista paginada de itens na wishlist do usuário.

# 6. Visão Geral "formOptionsRequests.js"
Funções para obter listas de opções usadas em formulários.

### getRolesRequest: 
Busca a lista completa de Funções (Roles), utilizando paginação.
```js
export async function getRolesRequest(
  accessToken,
  itemsPerPage = configsToPagination.itemsPerPage,
  currentPage = configsToPagination.currentPage,
) 
```

### getEnterprisesRequest: 
Busca a lista completa de Empresas "Enterprises", utilizando paginação.
```js
export async function getEnterprisesRequest(
  accessToken,
  itemsPerPage = configsToPagination.itemsPerPage,
  currentPage = configsToPagination.currentPage,
)
```

# 7. Visão Geral "utils.js"

Contém funções utilitárias para formatação de dados.

### FormatCPF(cpf): 
Formata uma string de CPF no padrão 999.999.999-99.
```js
export function formatCPF(cpf) {
  if (!cpf) return "";
  return cpf.replace(/^(\d{3})(\d{3})(\d{3})(\d{2})$/, "$1.$2.$3-$4");
}
```

### formatDate(date): 
Converte uma string de data para o formato de data local brasileiro (pt-BR).
```js
export function formatDate(date) {
  const dateConvert = new Date(date);
  return dateConvert.toLocaleDateString("pt-BR");
}
```

### formatCNPJ(cnpj): 
Função de formatação para CNPJ, atualmente não implementada.
```js
export function formatCNPJ(cnpj) {
  return;
}
```

### formatPromotion(value, percentage): 
Calcula e retorna o valor final de um item após aplicar uma porcentagem de desconto, fixado em 2 casas decimais.
```js
export function formatPromotion(value, percentage) {
  const each = (value * percentage);
  const result = value - each;
  return result.toFixed(2);
  ```

# 8. Visão Geral "storage.js"

Gerenciamento de armazenamento persistente usando AsyncStorage.

- >Storage.setItem(key, value): Salva um valor, serializando-o para JSON.
- >Storage.getItem(key): Lê um valor pela chave, fazendo o parse do JSON.
- >Storage.removeItem(key): Remove um item do armazenamento.

* Tratamento de Erros: Todos os métodos tratam erros de leitura/escrita com logs no console.