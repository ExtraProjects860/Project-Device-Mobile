# Visão Geral "screens" Telas do Aplicativo
Esta seção resume as funcionalidades e componentes das principais telas do aplicativo fornecidas.

# 1. Visão Geral "WishListScreen.jsx"
Esta tela permite que o usuário visualize e gerencie sua Lista de Desejos.

* Funcionalidade Principal: A tela busca e exibe itens da lista de desejos usando o callback fetchWishListCallback, que chama getItemsWishListRequest. A lista é renderizada usando o componente ListItems.

* Estado e Filtragem:

  * Gerencia o estado de busca (searchTerm) e utiliza um debounce de 500ms antes de atualizar o termo de busca real (debouncedSearchTerm).
  * Permite ao usuário alternar a ordem dos itens entre ASC e DESC através de handleToggleOrder.
  * Utiliza o componente SearchBar para entrada de pesquisa e controle de ordenação.

* Componentes de UI: Utiliza Background, NavBar e CardWishListItem para renderizar cada item.
```jsx
import CardWishListItem from "../components/ui/CardWishListItem";
//
<ListItems
        ref={listKey}
        callbackFetch={fetchWishListCallback}
        CardListRender={({ item }) => (
          <CardWishListItem item={item} onRemove={handleRemoveProduct} />
)}
```

# 2. Visão Geral "404.jsx"

Esta tela é exibida quando o usuário tenta acessar uma rota inexistente no aplicativo.

* Design: Exibe uma mensagem de erro estilizada com o código "404" em destaque, a mensagem "Ops.. Essa Página Não Existe", e o logo do aplicativo.

* Navegação: Oferece um link clicável (TouchableOpacity) que usa o hook useNavigateTo() para levar o usuário de volta para a rota /login (Tela de Login).

# 3. Visão Geral "HomeScreen.jsx"
Esta tela é a página inicial e o principal painel de identificação do cliente. Ela exibe informações pessoais e de afiliação do usuário logado.

* Dados Exibidos: O componente obtém os dados do usuário (userData) através do hook useAppContext().
  * Exibe o nome do usuário (ou "Nome não disponível" se ausente).
  * Exibe a foto do usuário (photo_url) ou uma imagem padrão (DefaultAvatar).
  * Exibe a Empresa (enterprise) à qual o usuário está associado (ou "Nenhuma Empresa Encontrada").
  * Exibe o CPF do usuário, formatado usando formatCPF.

### Identificação: 
Inclui um código de barras "Barcode" gerado a partir do "register_number" do usuário, utilizando o formato EAN8.
```jsx
          <View className="w-full px-4">
            <Barcode
              value={userData?.register_number || "0000000"}
              format="EAN8"
              width={4}
              height={70}
              lineColor="black"
              background="white"
            />
          </View>
        </View>
      </View>
```

- >Layout: Utiliza o componente Background e NavBar. Contém um rodapé fixo com informações de contato e endereço físico.