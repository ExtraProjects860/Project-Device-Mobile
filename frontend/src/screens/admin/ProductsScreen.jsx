import React, { useState } from "react";
import { View, Text } from "react-native";
import Background from "../../components/ui/Background";
import { NavBar } from "../../components/Navbar";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import SearchBar from "../../components/SearchBar";
import ListItems from "../../components/ListItems";
import ButtonAdd from "../../components/ui/ButtonAdd";
import CardProductList from "../../components/ui/CardProductList";
import { useThemeColors } from "../../hooks/useThemeColors.js";
import { getProductsRequest } from "../../lib/productsRequests.js";
import { useHandleRefresh } from "../../hooks/useHandleRefresh.js";
import ModalCreate from "../../components/modals/ModalCreateProduct";
import ModalUpdateProduct from "../../components/modals/ModalUpdateProduct";

export default function ProductsScreen() {
  const themeColors = useThemeColors();
  const { listKey, handleRefresh } = useHandleRefresh();

  const [isCreateProductModalVisible, setCreateProductVisible] =
    useState(false);
  const [isUpdateProductModalVisible, setUpdateProductVisible] =
    useState(false);
  const [selectedProduct, setSelectedProduct] = useState(null);

  const handleEditProduct = (product) => {
    setSelectedProduct(product);
    setUpdateProductVisible(true);
  };

  return (
    <Background>
      <ModalCreate
        visible={isCreateProductModalVisible}
        onClose={() => setCreateProductVisible(false)}
      />
      <ModalUpdateProduct
        visible={isUpdateProductModalVisible}
        onClose={() => setUpdateProductVisible(false)}
        product={selectedProduct}
        onProductUpdated={handleRefresh}
      />
      <NavBar />

      <View className="flex-row gap-2 m-6 items-center justify-center">
        <Icon name="basket-outline" size={30} color={themeColors.header} />
        <Text className="text-white font-bold text-3xl">Produtos</Text>
      </View>

      <View className="items-center mb-2">
        <SearchBar
          buttonAdd={
            <ButtonAdd
              onPress={() => setCreateProductVisible(true)}
              name={"basket-plus-outline"}
            />
          }
        />
      </View>

      <ListItems
        ref={listKey}
        callbackFetch={getProductsRequest}
        CardListRender={({ item }) => (
          <CardProductList item={item} onEdit={handleEditProduct} />
        )}
      />
    </Background>
  );
}
