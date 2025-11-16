## 1. Descrição Geral "Background.jsx"

O componente Background é responsável por renderizar a View principal de fundo da aplicação. Ele define a cor base do tema (claro/escuro) e exibe uma imagem decorativa no rodapé (footer), mantendo o restante do conteúdo sobreposto através do uso de children.

### Implementação

 ```js
- import React from "react";
- import { View, Image } from "react-native";
- import footer from "../../assets/images/footer.png";

  * Componente responsável pelo background do aplicativo.
  Renderiza a cor de fundo de acordo com o tema ativo e adiciona uma imagem decorativa fixa no rodapé.
 
 

export default function Background({ children } )
  return (
    <View className="flex-1 bg-light-primary dark:bg-dark-primary -z-10">
      {children}

      <View className="absolute bottom-0 w-full h-[400px] -z-10">
        <Image source={footer} className="w-full h-full" resizeMode="stretch" />
      </View>
    </View>
  );
}
```

###  Propriedades (Props)
Propriedade	Tipo Obrigatório Descrição children	ReactNode Conteúdo que será renderizado dentro do fundo. className	string (não usada atualmente). Permite passar estilos adicionais ao componente pode ser implementada futuramente.

### Comportamento Visual

- Cor de fundo:
  - Alterna automaticamente entre os temas light e dark (bg-light-primary / bg-dark-primary).

- Imagem decorativa:
  - Uma imagem (footer.png) é posicionada no rodapé, esticada horizontalmente (resizeMode="stretch").

- Z-index negativo:
  - Mantém o background atrás de todo o conteúdo renderizado em children.

### Exemplo de Uso
- import Background from "../components/ui/background";
- import { Text } from "react-native";

### Resumo

O componente Background provê uma estrutura visual consistente para todas as telas da aplicação, garantindo o uso de tema dinâmico, hierarquia de camadas correta e estilo unificado com a identidade visual do app.

### ----------------------------------------------------------------------------------
## 2. Descrição Geral "ButtonAdd.jsx"

O componente ButtonAdd é um botão de ação reutilizável que segue o padrão visual do aplicativo.
Ele é comumente utilizado para ações de adição (por exemplo, adicionar usuários, produtos ou itens), exibindo um ícone configurável e um rótulo padrão “Add”.

Este botão é personalizável via props, integrando-se facilmente com a paleta de cores do tema (claro/escuro) e o sistema de ícones MaterialCommunityIcons.

### Implementação

 ```js
import React from "react";
import { Text, TouchableOpacity } from "react-native";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import { useThemeColors } from "../../hooks/useThemeColors.js";


 * Componente responsável pelo botão de adicionar.
 * Exemplo de uso: botão para adicionar usuário na tela UsersScreen.
 
 * - Props:
 * - onPress: função executada ao clicar.
 * - name: nome do ícone (MaterialCommunityIcons).

  export default function ButtonAdd({ onPress, name }) {
  const themeColors = useThemeColors();
  return (
    <TouchableOpacity
      className="flex-row items-center bg-light-secondary dark:bg-dark-sencondary rounded-full p-2 shadow-md"
      onPress={onPress}
    >
      <Icon name={name} size={24} color={themeColors.header} />
      <Text className="ml-2 text-white font-semibold">Add</Text>
    </TouchableOpacity>
  );
}
```

### Comportamento Visual

- Formato: 
  - Botão arredondado (rounded-full) com leve sombra (shadow-md).

- Cor: 
  - Adapta-se automaticamente ao tema (claro/escuro) via classes bg-light-secondary e dark:bg-dark-sencondary.

- Ícone: 
  - Exibido à esquerda, definido dinamicamente pela prop name.

- Texto: 
  - “Add” em branco, com espaçamento lateral (ml-2) e fonte em negrito.

### Exemplo de Uso
```js
import ButtonAdd from "../components/ui/ButtonAdd";

export default function UsersScreen() {
  const handleAddUser = () => {
    console.log("Usuário adicionado!");
  };

  return (
    <ButtonAdd onPress={handleAddUser} name="account-plus" />
  );
}
```
###  Práticas de Uso

- Mantenha a semântica do botão: Use o ButtonAdd apenas para ações de adição (ex: criar, incluir, cadastrar). Para outras ações (como editar ou excluir), crie variantes dedicadas.

- Prefira nomes de ícones padronizados: Utilize nomes claros do pacote MaterialCommunityIcons. Ex: "plus", "account-plus", "plus-box-outline".

- Integração com tema:
O componente usa useThemeColors() para adaptar-se automaticamente ao tema ativo.
Certifique-se de que o hook esteja devidamente configurado no projeto.

- Evite sobrecarregar o botão com textos longos: O layout foi pensado para ícone + rótulo curto (“Add”). Se precisar de botões maiores ou com textos personalizados, crie uma variação específica.


### Resumo

O componente ButtonAdd fornece um botão estilizado e adaptável para ações de adição, com suporte a ícones dinâmicos, temas claros/escuros e integração visual com o design do app.
### -------------------------------------------------------------------------------------------

## 3. Descrição Geral "CardProductList.jsx"

O componente CardProductList é responsável por exibir as informações de um produto em formato de card dentro de listas ou catálogos.
Ele apresenta nome, preço, imagem, descrição, e — quando aplicável — um selo de promoção com o valor descontado.

Além disso, caso o usuário logado tenha perfil administrativo (ADMIN ou SUPERADMIN), o componente exibe um botão “Editar”, permitindo ações de gerenciamento de produto.

### Comportamento Visual

- Layout: 
  - Card arredondado com sombra e margens laterais.
- Tema: 
  - Alterna automaticamente entre claro e escuro.
- Promoção: 
  - Exibe um banner superior com a label % Promoção.
- Imagem: 
  - Mostra a foto do produto ou uma imagem padrão (shopping-bag.png).
- Admin: 
  - Exibe o botão “Editar” apenas se o usuário for administrador.

### Exemplo de Uso

```js
import CardProductList from "../components/ui/CardProductList";

export default function ProductScreen() {
  const handleEdit = (item) => {
    console.log("Editando produto:", item.name);
  };

  const sampleItem = {
    name: "Camiseta React Native",
    value: 79.9,
    discount: 10,
    description: "Camiseta estilosa com logo do React Native",
    photo_url: "",
    is_promotion_avaible: true,
  };

  return <CardProductList item={sampleItem} onEdit={handleEdit} />;
}
```

### Práticas de Uso

- Validação de dados:
  - Sempre verifique se o objeto item contém as chaves esperadas (name, value, description, etc).

- Atenção a permissões:
  - O botão “Editar” é exibido apenas para usuários com cargo ADMIN ou SUPERADMIN.
Garanta que userData esteja corretamente configurado no contexto global.

- Descontos e promoções:
  - Utilize a função formatPromotion(value, discount) para garantir o cálculo correto dos valores promocionais.

- Design responsivo:
  - O layout foi projetado com flexbox para se adaptar bem em diferentes tamanhos de tela.

- Fallbacks visuais:
  - Caso photo_url esteja ausente, o componente exibe automaticamente uma imagem padrão (DefaultProduct).

### Resumo

O componente CardProductList exibe informações completas de um produto com suporte a temas, descontos, imagem padrão e edição restrita a administradores.
Ideal para listas de produtos em catálogos ou painéis administrativos.

## 4. Descrição Geral "CardUserList.jsx"

O componente CardUserList é responsável por exibir as informações de um usuário em formato de card visual. Ele apresenta nome, cargo (role), empresa, registro, CPF e data de criação da conta, além da foto de perfil (ou um avatar padrão, se não houver imagem).

Também inclui um botão de edição “Editar” que executa uma função passada via prop, permitindo que o componente seja reutilizado em listas administrativas.

### Comportamento Visual

- Layout: 
  - Card horizontal com imagem à esquerda e dados à direita.
- Tema: 
  - Adapta-se automaticamente aos temas claro e escuro.
- Avatar: 
  - Exibe a imagem de perfil (photo_url) ou uma imagem padrão (avatar.png).
- Informações formatadas:
  - CPF formatado via formatCPF
  - Data formatada via formatDate
- Botão “Editar”: 
  - Sempre visível, com ícone e fundo colorido conforme o tema.

### Exemplo de Uso

```js
import CardUserList from "../components/ui/CardUserList";

export default function UserScreen() {
  const handleEdit = (user) => {
    console.log("Editando usuário:", user.name);
  };

  const exampleUser = {
    name: "Carlos Souza",
    role: "ADMIN",
    enterprise: "TechCorp",
    cpf: "12345678900",
    register_number: 42,
    created_at: "2025-10-01T14:32:00Z",
    photo_url: "",
  };

  return <CardUserList item={exampleUser} onEdit={handleEdit} />;
} 
```

### Práticas de Uso

- Padronize o formato de dados:
  - Garanta que os campos cpf, created_at e role sigam o padrão esperado para correta exibição e formatação.

- Evite dependência de contexto:
  - O componente não depende diretamente de contexto global — ele é totalmente reutilizável em qualquer tela.

- Respeite o design system:
  - A cor e o estilo do botão “Editar” são baseados nas variáveis de tema (useThemeColors), então mantenha a consistência.

- Fallback de imagem:
  - Se o campo photo_url for nulo ou inválido, o componente exibirá automaticamente um avatar padrão.

### Resumo

O componente CardUserList exibe informações detalhadas de um usuário
de forma elegante e responsiva, com suporte a temas e formatação de dados.
Ideal para telas de administração, controle de acesso ou listagens de usuários.

## 5. Descrição Geral "CardWishListItem.jsx"

O componente CardWishListItem representa um item dentro da lista de desejos "wishlist" do usuário.
Ele exibe informações básicas de um produto, como imagem, nome, valor, promoção e descrição, e inclui um botão para remover o item da lista.

O design é adaptado para tema claro e escuro, mantendo consistência com o restante da interface do app.

### Comportamento Visual

- Layout: Card horizontal com imagem do produto à esquerda e informações à direita.
- Tema: Adapta-se dinamicamente ao modo claro e escuro.

- Promoções:
  - Exibe o selo "% Promoção" quando "is_promotion_avaible" é verdadeiro.
  - Mostra o preço original riscado e o preço promocional formatado via "formatPromotion()".

- Botão de Remoção:
  - Usa ícone de lixeira "trash-can-outline".
  - Ação definida por "onRemove".
  - Fundo colorido conforme tema atual.

### Exemplo de Uso
```js
import CardWishListItem from "../components/ui/CardWishListItem";

export default function WishListScreen() {
  const handleRemove = () => {
    console.log("Produto removido da lista de desejos!");
  };

  const exampleItem = {
    name: "Tênis Esportivo XYZ",
    value: 299.99,
    discount: 20,
    description: "Confortável e leve, ideal para corridas.",
    photo_url: "",
    is_promotion_avaible: true,
  };

  return <CardWishListItem item={exampleItem} onRemove={handleRemove} />;
}
```

### Boas Práticas de Uso

- Função de remoção obrigatória:
  - Sempre defina a prop "onRemove" para garantir interatividade.

- Promoção condicional:
  - Defina "is_promotion_avaible" e "discount" apenas quando aplicável. O componente já lida automaticamente com a ausência de promoção.

- Evite textos longos:
  - O  campo description é truncado automaticamente para uma linha.

- Imagem padrão:
  - Se photo_url for nulo, o componente exibe o ícone de sacola (shopping-bag.png) como fallback.

### Resumo

O componente "CardWishListItem" fornece uma forma elegante e funcional
de exibir itens da lista de desejos do usuário, com destaque para promoções
e uma interação simples de remoção.


## 6. "Loading.jsx"

O componente Loading representa a tela de carregamento global do aplicativo.
Ele é exibido durante processos assíncronos. Como requisições, inicialização de dados ou autenticação, informando ao usuário que o sistema está processando algo.

Apresenta um indicador de atividade animado "ActivityIndicator" e uma mensagem textual dinâmica "Carregando..." que adiciona e remove pontos em sequência para transmitir movimento.

 ### Comportamento Visual

- Layout:
  - Fundo ocupa toda a tela (flex-1).

- Elementos centralizados tanto vertical quanto horizontalmente.

- Tema:
  - Fundo e texto se ajustam automaticamente conforme o modo claro ou escuro.

- Animação de texto:
  - A cada 300ms, o texto adiciona um ponto ao final, até três, e depois reinicia (Carregando, Carregando., Carregando.., Carregando...).

- Indicador de carregamento:
  - ActivityIndicator nativo do React Native com tamanho large e cor fixa (#E91D62 - Rosa avermelhado).

### Exemplo de Uso

```js
import Loading from "../components/ui/Loading";

export default function ExampleScreen() {
  const [isLoading, setIsLoading] = React.useState(true);

  React.useEffect(() => {
    setTimeout(() => setIsLoading(false), 3000);
  }, []);

  if (isLoading) {
    return <Loading />;
  }

  return (
    <Text>Tela carregada com sucesso!</Text>
  );
}
```

### Boas Práticas de Uso

- Renderização Condicional:
  - Exiba o componente Loading apenas quando o estado de carregamento estiver ativo. Ex: "isLoading === true".

- Evite sobreposição:
  - Como o componente cobre toda a tela, não o use em conjunto com outros elementos renderizados no mesmo nível.

- Integração com temas:
  - O componente já se adapta aos temas definidos globalmente, não há necessidade de ajustes manuais.

### Resumo

O componente Loading oferece uma experiência visual leve e informativa
durante processos de carregamento, com uma animação simples e responsiva,
mantendo o design consistente entre os modos claro e escuro.

## 7. Descrição Geral "Logo.jsx"

O componente Logo é responsável por exibir a logo oficial do aplicativo de forma simples e reutilizável.
Ele foi projetado para ser facilmente incorporado em diferentes partes da interface, como: telas de login, splash screens, headers e rodapés; permitindo personalização de tamanho e estilo através de propriedades.

### Comportamento Visual

- Imagem carregada:
  - A logo é importada do diretório local: src/assets/images/logo.png


- Estilização:
  - Através do className, é possível ajustar livremente dimensões, bordas, margens ou posicionamento. Exemplo:
```go
<Logo className="w-32 h-32 mx-auto" />
```

- Redimensionamento (resizedMode):
  - "contain" padrão: preserva as proporções da logo sem cortá-la.
  - "cover": preenche completamente a área especificada, podendo cortar partes.
  - "stretch": distorce para ocupar todo o espaço.

### Exemplo de Uso

```js
import Logo from "../components/ui/Logo";

export default function LoginScreen() {
  return (
    <View className="flex-1 items-center justify-center bg-light-primary dark:bg-dark-primary">
      <Logo className="w-40 h-40 mb-5" resizedMode="contain" />
      <Text className="text-white text-2xl font-bold">Bem-vindo de volta!</Text>
    </View>
  );
}
```

### Boas Práticas de Uso

- Evite redimensionamentos via código fixo:
  - Prefira usar classes "className" com estilos responsivos (w-24, h-24, etc.) em vez de valores absolutos.
- Manter proporção da logo:
  - Use "resizedMode=contain" para preservar o formato original da imagem.
- Reutilização:
  - Ideal para ser usado em múltiplas telas sem duplicação de código.

### Resumo

O componente Logo fornece uma forma simples e padronizada
de exibir a marca visual do aplicativo, garantindo consistência
e flexibilidade na aplicação de estilos e redimensionamento.

## 8. Descrição Geral "WarningNotFound.jsx"

O componente WarningNotFound é utilizado para exibir uma mensagem de aviso quando nenhum item é encontrado em uma listagem, busca ou consulta.
Ele serve como um feedback visual simples e direto para o usuário, mantendo consistência no design da interface.

### Comportamento Visual

- O componente centraliza vertical e horizontalmente o texto de aviso.
- Possui espaçamento superior (mt-20) para garantir que a mensagem não fique colada ao topo.
- A cor do texto (text-white) mantém contraste com o fundo, funcionando bem tanto em temas claros quanto escuros (com ajustes possíveis via tema).

### Exemplo de Uso

```js
import WarningNotFound from "../components/ui/WarningNotFound";

export default function ProductListScreen({ products }) {
  if (products.length === 0) {
    return <WarningNotFound message="Nenhum produto disponível no momento." />;
  }

  return (
    <FlatList
      data={products}
      renderItem={({ item }) => <CardProductList item={item} />}
    />
  );
}
```

### Resumo

O componente "WarningNotFound" exibe uma mensagem informativa
quando não há dados disponíveis, garantindo clareza e consistência
no feedback ao usuário.