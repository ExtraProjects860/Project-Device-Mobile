## Visão geral "AppContext.js"

- Núcleo de gerenciamento global de estado do seu app React Native — responsável por:
  - autenticação (token e dados do usuário),
  - tema (claro/escuro),
  - interceptação de requisições HTTP,
  - e verificação de conexão com a internet.

#
-  O AppContext é criado via:

```js
const AppContext = createContext();
```

-  E o provedor principal:

```js
export function AppProvider({ children }) { ... }
```

- > Envolve o app inteiro, fornecendo estado global e funções utilitárias.

### Persistência com Storage

Durante o useEffect inicial, ele carrega:


```js
const storedToken = await Storage.getItem("token");
const storedUserData = await Storage.getItem("user");
const storedTheme = await Storage.getItem("theme");
```

- E depois:
- >Atualiza os estados locais (setAccessToken, setUserData, etc.),
- >Define o tema inicial "setColorScheme(storedTheme", e começa a monitorar a conexão com NetInfo.

* Conclusão: toda a sessão (token, user e tema) sobrevive mesmo após o app ser fechado.

### Gerenciamento de tema
```js
const { colorScheme, setColorScheme } = useColorScheme();
```

- E funções auxiliares:
```js
const toggleTheme = async () => {
  const newTheme = colorScheme === "dark" ? "light" : "dark";
  await updateTheme(newTheme);
};
```

- Isso permite alternar instantaneamente entre modo claro e escuro, salvando a preferência no storage para uso futuro.
#
### Monitoramento de conexão

```js
const unsubscribe = NetInfo.addEventListener((state) => {
  setIsConnected(state.isConnected);
});
```

* Faz com que o app reaja automaticamente a mudanças de conectividade.
A prop "checkInternetConection", é fornecida aos componentes para verificar isso.

## Descrição Geral "ErrorContext.js"

* Cria um contextReact que expõe duas funções globais:
  * showErrorModal(message, onRetry) -> mostra o modal de erro.
  * hideErrorModal() -> fecha o modal.

Ele também injeta automaticamente o componente "<.ModalErrors />" no final da árvore do React. Ou seja, o modal sempre estará disponível em qualquer parte da aplicação, mesmo que você o chame de um hook ou função de serviço.

### Estrutura de estado

```js
const [errorState, setErrorState] = useState({
  visible: false,
  message: "",
  onRetry: null,
});
```

- Isso define três coisas:
  - visible: controla se o modal está visível.
  - message: a mensagem exibida no modal.
  - onRetry: callback opcional para tentar novamente uma ação.

### Funções principais
  - >"showErrorModal"

```js
const showErrorModal = useCallback((message, onRetryCallback) => {
  setErrorState({
    visible: true,
    message: message || "Ocorreu um erro inesperado.",
    onRetry: onRetryCallback,
  });
}, 
```

- Exibe o modal com a mensagem fornecida.
- Se nenhuma mensagem for passada, mostra uma genérica.
- Pode receber uma função onRetry (por exemplo, tentar novamente uma requisição).
##
- "hideErrorModal"

```js
const hideErrorModal = useCallback(() => {
  setErrorState((prev) => ({ ...prev, visible: false }));
}, 
```
* Fecha o modal, mantendo os outros dados.
#

### Retorno do Provider
```jsx
<ErrorContext.Provider value={value}>
  {children}
  <ModalErrors
    visible={errorState.visible}
    message={errorState.message}
    onClose={hideErrorModal}
    onRetry={errorState.onRetry ? () => {
      hideErrorModal();
      errorState.onRetry();
    } : null}
  />
</ErrorContext.Provider>
```

* children: toda a aplicação (ou parte dela).
* "<ModalErrors.>": o componente que renderiza o modal (geralmente um <Modal> do React Native).
* Se onRetry existir, ele:
#
* 1 -> Fecha o modal.

* 2 -> Executa a função de tentativa novamente.
#
### Hook auxiliar
```js
export function useError() {
  const context = useContext(ErrorContext);
  if (!context) {
    throw new Error("useError deve ser usado dentro de um ErrorProvider");
  }
  return context;
}
```

- Isso permite usar o contexto em qualquer componente, assim:

```js
const { showErrorModal } = useError();

try {
  await api.get('/data');
} catch (err) {
  showErrorModal("Falha ao carregar dados", () => fetchData());
}
```
#

## Visão Geral "PageLoader.js"

- Esse componente recebe dois props:
  - fetchData -> uma função de busca de dados.
  - children -> o conteúdo da tela (que será exibido somente após o carregamento).
#
- Assim, ele:
  - Usa o hook "useLoading()".
  - Chama "withLoading(fetchData)" dentro de um "useEffect()"; ou seja, executa fetchData assim que o componente monta, e automaticamente alterna o estado de carregamento.
  - Se "isLoading" for true, exibe o componente "<Loading. />". Caso contrário, renderiza o conteúdo (children).

### Ciclo completo de execução

- Montagem da página -> PageLoader monta.
- Executa withLoading(fetchData) → chama a função passada e ativa o estado de "carregando".
- Durante o carregamento -> isLoading === true -> renderiza "<.Loading />".
- Quando a requisição termina -> isLoading === false -> exibe os children.

### Função principal
```js
export default function PageLoader({ fetchData, children }) {
  const { isLoading, withLoading } = useLoading();

  useEffect(() => {
    withLoading(fetchData);
  }, [fetchData, withLoading]);

  if (isLoading) return <Loading />;

  return <>{children}</>;
}
```

### Exemplo de uso

Por exemplo:
#
```js
return (
  <PageLoader fetchData={initialLoad}>
    {/* conteúdo da lista vem aqui */}
  </PageLoader>
);
```
#

* Isso garante que:
  - A tela mostre um "<.Loading />" no início.
  - Só depois que "initialLoad()" terminar, os itens e componentes de lista aparecem.

### Resumo

- O PageLoader é um wrapper reutilizável de carregamento de página. Ele garante que toda tela que dependa de um "fetchData()" só seja exibida após a conclusão da requisição, mostrando o "<.Loading />" automaticamente no meio do processo.

- Em conjunto com useLoading e Loading.jsx, ele compõe uma arquitetura elegante de controle visual de estado assíncrono.