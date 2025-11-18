## 1. Visão Geral "useFormOpitions.js"

* Ele é um custom hook para:
- >Fazer fetch de duas coleções simultaneamente:
- > Funções de usuário (roles)
- > Empresas (enterprises)

* Gerenciar os estados de carregamento, erro e dados;
* Integrar com o contexto global (AppContext e ErrorContext);
* Oferecer uma função refetch para recarregar os dados sob demanda.

### Estrutura principal

```js
export function useFormOptions() {
  const { accessToken } = useAppContext();
  const { showErrorModal } = useError();

  const [roles, setRoles] = useState([]);
  const [enterprises, setEnterprises] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);

  const fetchOptions = useCallback(async () => {
    if (!accessToken) return;
    
  }, [accessToken, showErrorModal]);

  useEffect(() => {
    fetchOptions();
  }, [fetchOptions]);

  return { roles, enterprises, isLoading, error, refetch: fetchOptions };
}
```

### Passo a passo do fluxo

- Acessa o token de autenticação
```js
const { accessToken } = useAppContext();
```

- Sem ele, as requisições não são feitas.
#
### Busca os dados
```js
const [rolesData, enterprisesData] = await Promise.all([
  getRolesRequest(accessToken, allItemsLimit, firstPage),
  getEnterprisesRequest(accessToken, allItemsLimit, firstPage),
]);
```

* O uso de Promise.all é eficiente, pois executa ambas as requisições em paralelo.
* allItemsLimit e firstPage definem um limite alto para capturar tudo de uma vez (500 registros, página 1).
#
### Atualiza os estados locais
```js
setRoles(rolesData);
setEnterprises(enterprisesData);
```
#
### Em caso de erro

* Mostra o modal de erro global via ErrorContext:
```js
showErrorModal("Não foi possível carregar as opções de Funções e Empresas.");
```
- E também guarda o erro em setError(err) para controle local.
#
### Gerencia o estado de carregamento
```js
setIsLoading(true); ... finally { setIsLoading(false); }
```

### Executa automaticamente ao montar
```js
useEffect(() => { fetchOptions(); }, [fetchOptions]);
```
#
### Retorno do hook
```js
return {
  roles,         // lista de funções
  enterprises,   // lista de empresas
  isLoading,     // indica se está carregando
  error,         // guarda o erro (se houver)
  refetch: fetchOptions, // função para recarregar manualmente
};
```
- Assim, qualquer componente de formulário pode fazer:
```js
const { roles, enterprises, isLoading, refetch } = useFormOptions();
```
- E ter dados sempre atualizados e reusáveis.
#
### Integração no sistema

- Esse hook deve ser usado em páginas de cadastro ou edição de usuários, produtos, ou qualquer entidade que precise preencher selects de “Função” ou “Empresa”.

- >Exemplo:
```js
function UserForm() {
  const { roles, enterprises, isLoading } = useFormOptions();

  if (isLoading) return <Loading />;

  return (
    <>
      <Select label="Função" options={roles} />
      <Select label="Empresa" options={enterprises} />
    </>
  );
}
```
#
### Em resumo
- >Função | Descrição
- >useFormOptions() | Hook para carregar e gerenciar opções de formulários
- >roles | Array de funções (roles) retornadas da API
- >enterprises | Array de empresas
- >isLoading | Flag de carregamento
- >error | Objeto de erro se algo falhar
- >refetch() | Recarrega manualmente as opções
#

## 2. Visão geral "useHandleLogoutConfirm.js"

Custom hook que cria e retorna uma função handleLogout.

- Essa função, quando chamada, faz duas coisas:
  - Remove os dados locais de autenticação (token e usuário);
  - Redireciona o usuário para a tela de login (/login).

### explicação passo a passo:
```js
import { useAppContext } from "../context/AppContext";
import { useNavigateTo } from "./useNavigateTo";

export function useHandleLogoutConfirm() {
  const goTo = useNavigateTo(); // Hook de navegação

  const { manuallyLogout } = useAppContext(); // Função global de logout

  const handleLogout = async () => {
    await manuallyLogout(); // Limpa token e dados do usuário do storage
    goTo("/login"); // Redireciona para tela de login
  };
  
  return handleLogout; // Retorna a função pronta para uso
}
```
# 
### Como ele se encaixa no sistema

Ele é usado apartir do trecho:

```js
const handleLogoutConfirm = useHandleLogoutConfirm();
```

Depois, é passado para o modal de confirmação:
```js
<ModalWarning
  visible={isLogoutModalVisible}
  message={"Você tem certeza que deseja sair da sua conta?"}
  onClose={() => setLogoutModalVisible(false)}
  onConfirm={handleConfirmLogout}
/>
```
E dentro de handleConfirmLogout:
```js
const handleConfirmLogout = () => {
  handleLogoutConfirm();
  setLogoutModalVisible(false);
};
```
Ou seja:
- Usuário clica em Logout no menu.
- Abre o ModalWarning pedindo confirmação.
- Ao confirmar, o hook faz o logout seguro e volta para o login.

### Em resumo
- > Função | Descrição
- useHandleLogoutConfirm() | Hook que gera a função de logout com redirecionamento
- handleLogout() | Executa manuallyLogout() e redireciona para /login
- Dependências | useAppContext, useNavigateTo
- Uso principal | Menu.jsx no botão de “Logout”
#
- > Esse hook é o elo final entre o contexto global de autenticação e a interface de usuário (menus e modais). Ele garante que o fluxo de logout sempre siga o mesmo padrão e evita duplicação de código.

#
## 3. Visão Geral "useHandlerRefresh.js"

Hook personalizado criado para forçar a recarga de componentes especialmente listas (FlatList, ScrollView, etc).

### Funcionamento

```js
import { useState } from "react";

export function useHandleRefresh() {
  const [listKey, setListKey] = useState(0);

  const handleRefresh = () => {
    setListKey((prevKey) => prevKey + 1);
  };

  return {
    listKey,
    handleRefresh,
  };
}
```

- > "listKey": estado numérico que serve como identificador de renderização. Ele começa em 0 e é incrementado sempre que o usuário quer “recarregar” o conteúdo.
- > "handleRefresh": função que incrementa listKey. Como o valor da chave muda, o React força uma nova renderização do componente que a usa.
- > "Retorno" -> o hook retorna ambos:
  - >"listKey": a chave que muda a cada atualização;
  - >handleRefresh: a função que dispara a atualização.

#
## 4. Visão Geral "useLoading.js"

Hook personalizado responsável por controlar estados de carregamento "isLoading" e tratar erros automaticamente, exibindo uma mensagem amigável via ErrorContext.

Ele é especialmente útil quando executar funções assíncronas com indicador de carregamento e tratamento de erros centralizado, sem duplicar lógica em cada componente.

###  Funcionamento
```js
import { useCallback, useState } from "react";
import { useError } from "../context/ErrorContext";

export function useLoading(initialState = false) {
  const [isLoading, setIsLoading] = useState(initialState);
  const { showErrorModal } = useError();

  const withLoading = useCallback(
    async (asyncFunction) => {
      try {
        setIsLoading(true);
        await asyncFunction();
      } catch (error) {
        console.error("Erro capturado pelo withLoading:", error);
        showErrorModal("Erro ao carregar componente.");
      } finally {
        setIsLoading(false);
      }
    },
    [showErrorModal],
  );

  return { isLoading, withLoading };
}
```

## 
- > isLoading: 
  - > Estado booleano que indica se uma operação está em andamento.
  - > Inicia com false, mas pode receber um estado inicial via initialState.

- > setIsLoading:
  - >Atualiza o estado de carregamento (true para iniciar, false para encerrar).

- > useError:
  - > Importa o showErrorModal() do contexto de erro global para exibir mensagens padronizadas caso algo falhe.

- > withLoading(asyncFunction):
  - > Função wrapper que executa uma operação assíncrona com controle automático de carregamento e erros.

- > Exemplo de comportamento:
  - > Antes: ativa o isLoading = true;
  - > Tenta executar a função passada;
  - > Se ocorrer erro → exibe modal e registra no console;
  - >Por fim: isLoading = false sempre é chamado (mesmo em erro).

## 5. Visão Geral "useMenu.js"

Hook de controle de estado que gerencia a visibilidade de menus ou pop ups interativos. Ele fornece uma interface padronizada para abrir e fechar menus, centralizando a lógica e evitando repetição de código em componentes.

### Funcionamento
```js
import { useState } from "react";

export function useMenu() {
  const [isVisible, setIsVisible] = useState(false);

  const openMenu = () => setIsVisible(true);
  const closeMenu = () => setIsVisible(false);

  return {
    isVisible,
    openMenu,
    closeMenu,
  };
}
```
- > isVisible:
  - > Estado booleano que representa se o menu está visível ou não.
  - > Inicia como false (menu fechado por padrão).

- > setIsVisible:
  - > Atualiza o estado de visibilidade.

- > openMenu():
  - > Define isVisible = true > torna o menu visível.

- > closeMenu():
  - > Define isVisible = false > oculta o menu.

- > Retorno final:
  - > O hook retorna um objeto com as funções e o estado atual

## 6. Visão Geral "useNavigateTo.js"

Custom hook criado para simplificar e centralizar a navegação entre telas no aplicativo React Native que utiliza React Router Native como sistema de rotas. Ele abstrai o uso direto de useNavigate, permitindo uma chamada mais limpa e consistente "goTo(/rota, params)".

### Funcionamento
```js
import { useNavigate } from "react-router-native";

export function useNavigateTo() {
  const navigate = useNavigate();

  const goTo = (screen, params = {}) => {
    navigate(screen, { state: params });
  };

  return goTo;
}
```
- >Importação do hook nativo useNavigate:
  - >Vem da biblioteca react-router-native, usada para navegação no React Native.
  - >Retorna a função navigate, que é responsável por mudar de rota.

- >Função "goTo":
  - >Cria uma camada personalizada sobre navigate.
  - >Recebe dois parâmetros:
    - >screen: a rota/tela de destino (ex: "/login");
    - >params: um objeto com dados que serão enviados como estado para a próxima tela.

- >Retorno:
  - >O hook retorna diretamente a função goTo, tornando seu uso extremamente simples.

## 7. Visão Geral "usePagination.js"

Custom hook responsável por gerenciar o estado e o comportamento da paginação de listas no app.

Ele abstrai completamente a lógica dos...
- Carregar a lista inicial de dados;
- Fazer scroll infinito para buscar mais páginas;
- Atualizar a lista via pull-to-refresh;
- Controlar erros e estados de carregamento.

O hook integra-se com o AppContext (para acessar o token de autenticação) e com o ErrorContext (para exibir erros em modais).

## Funcionamento
```js
import { useState, useCallback, useRef, useEffect } from "react";
import { useError } from "../context/ErrorContext";
import { useAppContext } from "../context/AppContext";

const initialData = {
  data: [],
  refreshing: false,
  loadingMore: false,
  currentPage: 1,
  totalPages: 1,
  totalResult: 0,
};

export function usePagination(callbackFetch, searchFilter) {
  const { accessToken } = useAppContext();
  const { showErrorModal } = useError();

  const [listItems, setListItems] = useState(initialData.data);
  const [currentPage, setCurrentPage] = useState(initialData.currentPage);
  const [totalPages, setTotalPages] = useState(initialData.totalPages);
  const [totalResult, setTotalResult] = useState(initialData.totalResult);
  const [isLoadingMore, setIsLoadingMore] = useState(initialData.loadingMore);
  const [isRefreshing, setRefreshing] = useState(initialData.refreshing);

  const isFetchingRef = useRef(false);
  const flatListRef = useRef(null);
  const allItemsLoaded = currentPage >= totalPages && listItems.length > 0;
  const itemsPerPage = 20;
  ...
}
```

## Funções Principais

### "initialLoad"

Carrega os primeiros dados da lista (página 1).
Chamado automaticamente na inicialização e também usado em refresh.
```js
const initialLoad = useCallback(async () => {
  const result = await callbackFetch(20, 1, accessToken, searchFilter);
  setListItems(result.data);
  ...
}, [...]);
```
Exibe um modal de erro e tenta recarregar automaticamente caso a requisição falhe.

### "loadMore"

Carrega a próxima página da lista quando o usuário chega ao final da rolagem.
```js
const loadMore = useCallback(async () => {
  if (isLoadingMore || currentPage >= totalPages) return;

  const result = await callbackFetch(20, currentPage + 1, accessToken, searchFilter);
  setListItems(prev => [...prev, ...result.data]);
  setCurrentPage(result.current_page);
}, [...]);
```
Evita chamadas duplicadas utilizando isFetchingRef e isLoadingMore.

### "handleRefresh"

Recarrega os dados do início (equivalente ao pull-to-refresh).
```js
const handleRefresh = useCallback(async () => {
  setRefreshing(true);
  await initialLoad();
  setRefreshing(false);
}, [initialLoad]);
```

### "scrollToTop"

Rola a lista de volta para o topo de forma animada.
```js
const scrollToTop = () => {
  flatListRef.current?.scrollToOffset({ offset: 0, animated: true });
};
```

## 8. Visão Geral "UseThemeColors.js"

Custom hook responsável por fornecer as cores apropriadas do tema (claro ou escuro) para o restante dos componentes do aplicativo.

Ele faz a leitura da configuração de cores diretamente do tailwind.config.js, garantindo consistência visual entre os estilos definidos no Tailwind e os componentes React Native.

### Dependências
```js
import resolveConfig from "tailwindcss/resolveConfig";
import tailwindConfig from "../../tailwind.config.js";
import { useAppContext } from "../context/AppContext.js";
```

* resolveConfig: converte o arquivo tailwind.config.js em um objeto completo de configuração.
* useAppContext: acessa o estado global do tema (isThemeDark) armazenado no contexto principal da aplicação.

### Funcionamento
```js
const fullConfig = resolveConfig(tailwindConfig);
const colors = fullConfig.theme.colors;

export function useThemeColors() {
  const { isThemeDark } = useAppContext();

  return {
    primary: isThemeDark ? colors.dark["text-primary"] : colors.light.primary,

    header: isThemeDark
      ? colors.dark["text-primary"]
      : colors.light["text-inverted"],

    switch: {
      track: {
        true: colors.dark["text-secondary"],
        false: colors.dark["text-secondary"],
      },
      thumb: isThemeDark ? colors.dark["white"] : colors.light.primary,
    },
  };
}
```
#
- >Importa e processa as configurações do Tailwind para ter acesso às cores (colors).
- >Obtém o estado do tema (claro ou escuro) via useAppContext.
- >Retorna um conjunto de cores dinâmicas, alternando conforme o valor de isThemeDark.

##