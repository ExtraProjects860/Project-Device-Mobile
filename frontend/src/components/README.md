## 1. Descrição Geral "Header.jsx" 

O componente Header é responsável por exibir o cabeçalho principal do aplicativo, contendo a logo da aplicação e o ícone de menu que pode ser utilizado para abrir uma barra lateral, navegação ou opções adicionais.

Ele também gerencia dinamicamente a visibilidade do ícone de menu conforme o status da conexão com a internet, garantindo melhor controle de interface.

### Comportamento Visual

- O cabeçalho possui fundo temático, alternando entre cores claras e escuras de acordo com o tema ativo (light ou dark).
- A logo é exibida centralizada na horizontal, com o ícone de menu à direita.
- O StatusBar é configurado com estilo light-content e modo translúcido.
- Quando não há conexão com a internet, o botão de menu é ocultado "checkInternetConection === false".

### Exemplo de Uso

```jsx
import Header from "../components/Header";

export default function HomeScreen({ navigation }) {
  const handleMenu = () => {
    navigation.toggleDrawer(); // exemplo com React Navigation
  };

  return (
    <>
      <Header onMenuPress={handleMenu} />
      {/* Conteúdo da tela */}
    </>
  );
}
```

### Dependências Utilizadas

- react-native-vector-icons / MaterialCommunityIcons → Ícone de menu.
- useThemeColors → Hook personalizado que fornece as cores do tema atual.
- useAppContext → Hook de contexto global para acessar o status da conexão.

### Resumo

O componente Header fornece um cabeçalho funcional e adaptativo,
com integração direta ao tema e controle de visibilidade dinâmica
do ícone de menu conforme o estado da conexão do aplicativo.

## 2. Descrição Geral "ListItems.jsx"

O componente ListItems é um componente genérico responsável por renderizar listas paginadas de itens no aplicativo.

Ele integra recursos de:

* Paginação automática (scroll infinito)
* Atualização por gesto de pull-to-refresh
* Mensagem quando não há resultados
* Botão de "Voltar ao topo" ao final da lista
* Indicadores visuais de carregamento

Esse componente serve como estrutura base para diferentes telas que exibem listas de elementos, recebendo dinamicamente o componente de renderização de cada item e a função de busca de dados.

### Comportamento Visual

- Enquanto carrega os primeiros dados, exibe a tela do componente PageLoader.
- Durante carregamentos adicionais (scroll), exibe o componente Loading.
- Se todos os itens já foram carregados, exibe um botão "Voltar ao Topo".
- Caso não haja resultados, mostra o componente WarningNotFound.
- Exibe o total de itens renderizados em relação ao total geral retornado pela API.

### Lógica Interna e Hooks

- "usePagination(callbackFetch, searchFilter)" controla:
  - Estado dos itens (listItems)
  - Paginação e carregamento incremental (loadMore)
  - Atualização por gesto (handleRefresh)
  - Scroll até o topo (scrollToTop)
  - Indicadores de carregamento (isRefreshing, isLoadingMore)

- "PageLoader"
  - Wrapper que executa o carregamento inicial e gerencia erros globais de fetch.

- "Loading" e "WarningNotFound"
  - Elementos visuais auxiliares para estados de lista vazia ou carregamento.

### Exemplo de Uso

```jsx
import ListItems from "../components/ListItems";
import CardUserList from "../components/ui/CardUserList";
import { fetchUsers } from "../api/users";

export default function UsersScreen() {
  return (
    <ListItems
      callbackFetch={fetchUsers}
      CardListRender={CardUserList}
      searchFilter=""
    />
  );
}
```
### Dependências Utilizadas

- React Native FlatList → renderização otimizada de listas grandes.
- usePagination → hook customizado para paginação e controle de scroll.
- PageLoader, Loading, WarningNotFound → componentes auxiliares de estado.

### Resumo

O componente ListItems oferece uma estrutura reutilizável, robusta e responsiva para exibir dados paginados,
integrando feedback visual de carregamento, controle de atualização e gerenciamento automático de fim de lista.

## 3. Descrição Geral "Menu.jsx"

O componente Menu é o menu lateral interativo (sidebar) do aplicativo, responsável por exibir opções de navegação, configuração, ações administrativas e logout.

Ele é totalmente animado, com transições suaves de abertura e fechamento via React Native Animated, e adaptado para modo claro e escuro.

- Além disso, contém:
  - Controle de visibilidade com animação de fade + slide
  - Modais de alteração de senha e confirmação de logout
  - Alternância de tema escuro/claro
  - Seções distintas para usuário comum e administrador
  - Navegação integrada via hook customizado useNavigateTo

### Comportamento Visual e Funcional

- Animação: 
  - Abertura e fechamento com transições simultâneas (fade + slideX).

- Tema: Permite 
  - alternar entre modo claro e modo escuro com Switch.

- Modais:
  - PasswordChange → altera senha do usuário.
  - ModalWarning → confirma ação de logout.

- Navegação:
  - /home — tela inicial
  - /products — listagem de produtos
  - /wishlist — lista de desejos
  - /users — (admin) gerenciar usuários
  - /products — (admin) gerenciar produtos

- Permissões:
  - Opções “Admin” visíveis apenas se role for "ADMIN" ou "SUPERADMIN".

- Fechamento do Menu:
  - Ao clicar fora (Pressable) ou no botão de voltar.

### Hooks e Dependências Internas
* Hook / Módulo	Função

- >useNavigateTo   |	Hook customizado de navegação interna do app
- >useAppContext    |	Gerencia tema, dados do usuário e estado global
- >useThemeColors   |	Retorna as cores do tema atual (dark/light)
- >useHandleLogoutConfirm   |	Gerencia confirmação e execução de logout
- >Animated | Controla animações de slide e fade do menu

### Exemplo de Uso

```jsx
import Menu from "../components/Menu";
import { useState } from "react";

export default function HomeScreen() {
  const [isMenuVisible, setMenuVisible] = useState(false);

  return (
    <>
      <Header onMenuPress={() => setMenuVisible(true)} />
      <Menu visible={isMenuVisible} closeMenu={() => setMenuVisible(false)} />
    </>
  );
}
```

### Resumo

O componente Menu combina design dinâmico, controle de tema, navegação e segurança.
Ele oferece uma experiência fluida e moderna, sendo um dos principais elementos de interação do app.


## 4. Descrição Geral "Navbar.jsx"

Wrapper de navegação principal. Ele serve como container que junta o cabeçalho "Header" e o menu lateral "Menu" em um único lugar.

### Estrutura e Função

- Import:
  - Header -> o topo do app com o logo e ícone de menu.
  - Menu -> o menu lateral deslizante (drawer) com configurações e navegação.
  - useMenu -> hook customizado que controla a visibilidade do menu (aberto/fechado).

### Lógica principal

```jsx
const { isVisible, openMenu, closeMenu } = useMenu();
```

- "isVisible" -> controla se o menu está sendo exibido.
- "openMenu()" -> abre o menu (passado para o Header, que chama isso ao clicar no ícone).
- "closeMenu()" -> fecha o menu (usado internamente no próprio Menu).


## 5. Descrição Geral "SearchBar.jsx"

Componente de busca com suporte a ordenação e botão adicional.

### Função e Estrutura
- > O componente renderiza uma barra de pesquisa personalizável, com:
- Ícone de ordenação ascendente/descendente,
- Campo de texto para busca,
- Ícone de lupa (estético),
- E opcionalmente, um botão adicional à direita (ex: “adicionar item”).

### Lógica de ordenação

```jsx
const orderIcon = itemsOrder === "ASC" ? "sort-ascending" : "sort-descending";
```

* Escolhe dinamicamente o ícone certo conforme o estado atual de ordenação.

### Estrutura visual (Tailwind)

* Container principal com flex-row, espaçamento (gap-x-2) e padding lateral.

* Campo de busca com:
  * fundo branco (bg-white),
  * bordas arredondadas (rounded-full),
  * padding horizontal interno (px-4),
  * ícones à esquerda e à direita.