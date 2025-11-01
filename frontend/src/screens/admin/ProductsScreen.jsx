import React, { useState, useEffect, useCallback } from "react";
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
import { useAppContext } from "../../context/AppContext.js";

export default function ProductsScreen() {
  const themeColors = useThemeColors();
  const { listKey, handleRefresh } = useHandleRefresh();
  const { accessToken, userData } = useAppContext();

  const isAdmin = userData?.role === "ADMIN" || userData?.role === "SUPERADMIN";
  const [isCreateProductModalVisible, setCreateProductVisible] =
    useState(false);
  const [isUpdateProductModalVisible, setUpdateProductVisible] =
    useState(false);
  const [selectedProduct, setSelectedProduct] = useState(null);

  const [searchTerm, setSearchTerm] = useState("");
  const [debouncedSearchTerm, setDebouncedSearchTerm] = useState(searchTerm);
  const [itemsOrder, setItemsOrder] = useState("ASC");

  useEffect(() => {
    const handler = setTimeout(() => {
      setDebouncedSearchTerm(searchTerm);
    }, 500);

    return () => {
      clearTimeout(handler);
    };
  }, [searchTerm]);

  const fetchProductsCallback = useCallback(
    (itemsPerPage, currentPage) => {
      return getProductsRequest(
        itemsPerPage,
        currentPage,
        accessToken,
        debouncedSearchTerm.toUpperCase(),
        itemsOrder
      );
    },
    [debouncedSearchTerm, accessToken, itemsOrder]
  );

  const handleEditProduct = (product) => {
    setSelectedProduct(product);
    setUpdateProductVisible(true);
  };

  const handleToggleOrder = () => {
    setItemsOrder((prevOrder) => (prevOrder === "ASC" ? "DESC" : "ASC"));
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
            isAdmin && (
              <ButtonAdd
                onPress={() => setCreateProductVisible(true)}
                name={"basket-plus-outline"}
              />
            )
          }
          searchValue={searchTerm}
          onSearchChange={setSearchTerm}
          itemsOrder={itemsOrder}
          onToggleOrder={handleToggleOrder}
        />
      </View>

      <ListItems
        ref={listKey}
        callbackFetch={fetchProductsCallback}
        CardListRender={({ item }) => (
          <CardProductList item={item} onEdit={handleEditProduct} />
        )}
      />
    </Background>
  );
}
