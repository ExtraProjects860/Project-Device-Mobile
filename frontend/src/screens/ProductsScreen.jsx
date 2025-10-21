import React from "react";
import { View, Text } from "react-native";
import Background from "../components/ui/Background";
import { NavBar } from "../components/Navbar";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import SearchBar from "../components/SearchBar.jsx";
import ListItems from "../components/ListItems.jsx";
import ButtonAdd from "../components/ui/ButtonAdd.jsx";
import CardProductList from "../components/ui/CardProductList.jsx";
import { useThemeColors } from "../hooks/useThemeColors.js";
import { getProductsRequest } from "../lib/productsRequests.js";
import { useState } from "react";
import ModalCreate from "../components/modals/ModalCreateProduct";
import ModalUpdateProduct from "../components/modals/ModalUpdateProduct.jsx";
import { useRef } from "react";

export default function ProductsScreen() {
  const themeColors = useThemeColors();
  const [isCreateProductModalVisible, setCreateProductVisible] =
    useState(false);
  const [isUpdateProductModalVisible, setUpdateProductVisible] =
    useState(false);
  const [selectedProduct, setSelectedProduct] = useState(null);
  const listRef = useRef(null);

  const handleEditProduct = (product) => {
    setSelectedProduct(product);
    setUpdateProductVisible(true);
  };

  const handleRefresh = () => {
    if (listRef.current) {
      listRef.current.refresh();
    }
  };

  const CardProductRender = ({ item }) => (
    <CardProductList item={item} onEdit={handleEditProduct} />
  );

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
        <Icon name="shopping-outline" size={30} color={themeColors.header} />
        <Text className="text-white font-bold text-3xl">Produtos</Text>
      </View>

      <View className="items-center mb-2">
        <SearchBar
          buttonAdd={
            <ButtonAdd
              onPress={() => setCreateProductVisible(true)}
              name={"shopping-outline"}
            />
          }
        />
      </View>

      <ListItems
        ref={listRef}
        callbackFetch={getProductsRequest}
        CardListRender={CardProductRender}
      />
    </Background>
  );
}
