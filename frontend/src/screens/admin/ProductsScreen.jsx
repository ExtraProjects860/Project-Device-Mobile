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
import {
  addInWishListRequest,
  getItemsWishListRequest,
} from "../../lib/wishListRequests.js";
import { useError } from "../../context/ErrorContext.js";
import ModalCheck from "../../components/modals/ModalCheck.jsx";

export default function ProductsScreen() {
  const themeColors = useThemeColors();
  const { listKey, handleRefresh } = useHandleRefresh();
  const { accessToken, userData } = useAppContext();
  const { showErrorModal } = useError();

  const isAdmin = userData?.role === "ADMIN" || userData?.role === "SUPERADMIN";

  const [isCreateProductModalVisible, setCreateProductVisible] =
    useState(false);
  const [isUpdateProductModalVisible, setUpdateProductVisible] =
    useState(false);
  const [selectedProduct, setSelectedProduct] = useState(null);

  const [isSuccessModalVisible, setSuccessModalVisible] = useState(false);
  const [successMessage, setSuccessMessage] = useState("");

  const [searchTerm, setSearchTerm] = useState("");
  const [debouncedSearchTerm, setDebouncedSearchTerm] = useState(searchTerm);
  const [itemsOrder, setItemsOrder] = useState("DESC");

  const [wishlistProductIds, setWishlistProductIds] = useState(new Set());

  const fetchWishlistIds = useCallback(async () => {
    if (isAdmin) return;

    try {
      const wishlistData = await getItemsWishListRequest(
        500,
        1,
        accessToken,
        "",
        "ASC"
      );

      if (wishlistData?.data) {
        const ids = wishlistData.data.map((item) => item.id);
        setWishlistProductIds(new Set(ids));
      }
    } catch (error) {
      console.error("Erro ao buscar IDs da wishlist:", error);
    }
  }, [accessToken, isAdmin]);

  useEffect(() => {
    fetchWishlistIds();
  }, [fetchWishlistIds, listKey]);

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

  const handleAddWishList = async (product) => {
    try {
      if (!product?.id) {
        showErrorModal("Produto inválido.");
        return;
      }

      if (wishlistProductIds.has(product.id)) {
        showErrorModal("Este item já está na sua lista de desejos.");
        return;
      }

      await addInWishListRequest(product.id, accessToken);

      setWishlistProductIds((prevIds) => new Set(prevIds).add(product.id));

      setSuccessMessage("Produto adicionado à sua lista de desejos.");
      setSuccessModalVisible(true);
    } catch (error) {
      console.error("Erro ao adicionar à wishlist:", error);
      const errorMessage =
        error.response?.data?.error ||
        "Não foi possível adicionar o item. Tente novamente.";

      if (error.response?.status === 409) {
        setWishlistProductIds((prevIds) => new Set(prevIds).add(product.id));
      } else {
        showErrorModal(errorMessage);
      }
    }
  };

  const handleToggleOrder = () => {
    setItemsOrder((prevOrder) => (prevOrder === "ASC" ? "DESC" : "ASC"));
  };

  return (
    <Background>
      <ModalCreate
        visible={isCreateProductModalVisible}
        onClose={() => setCreateProductVisible(false)}
        onProductCreated={handleRefresh}
      />
      <ModalUpdateProduct
        visible={isUpdateProductModalVisible}
        onClose={() => setUpdateProductVisible(false)}
        product={selectedProduct}
        onProductUpdated={handleRefresh}
      />
      <ModalCheck
        visible={isSuccessModalVisible}
        message={successMessage}
        onClose={() => setSuccessModalVisible(false)}
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
        key={listKey}
        callbackFetch={fetchProductsCallback}
        CardListRender={({ item }) => (
          <CardProductList
            item={item}
            onClick={isAdmin ? handleEditProduct : handleAddWishList}
            isAdded={!isAdmin && wishlistProductIds.has(item.id)}
          />
        )}
      />
    </Background>
  );
}
