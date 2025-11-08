import React, { useState, useEffect, useCallback } from "react";
import { View, Text } from "react-native";
import Background from "../components/ui/Background";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import { NavBar } from "../components/Navbar";
import SearchBar from "../components/SearchBar";
import CardWishListItem from "../components/ui/CardWishListItem";
import ListItems from "../components/ListItems";
import { useThemeColors } from "../hooks/useThemeColors.js";
import { useHandleRefresh } from "../hooks/useHandleRefresh.js";
import { getItemsWishListRequest } from "../lib/wishListRequests.js";
import { useAppContext } from "../context/AppContext.js";
import { Alert } from "react-native";
import { deleteWishListRequest } from "../lib/wishListRequests.js";
import { useError } from "../context/ErrorContext.js";
import ModalWarning from "../components/modals/ModalWarning.jsx";

export default function WishListScreen() {
  const themeColors = useThemeColors();
  const { listKey, handleRefresh } = useHandleRefresh();
  const { accessToken } = useAppContext();
  const { showErrorModal } = useError();

  const [searchTerm, setSearchTerm] = useState("");
  const [debouncedSearchTerm, setDebouncedSearchTerm] = useState(searchTerm);
  const [itemsOrder, setItemsOrder] = useState("ASC");

  const [itemToRemove, setItemToRemove] = useState(null);
  const [isWarningVisible, setWarningVisible] = useState(false);

  useEffect(() => {
    const handler = setTimeout(() => {
      setDebouncedSearchTerm(searchTerm);
    }, 500);

    return () => {
      clearTimeout(handler);
    };
  }, [searchTerm]);

  const fetchWishListCallback = useCallback(
    (itemsPerPage, currentPage) => {
      return getItemsWishListRequest(
        itemsPerPage,
        currentPage,
        accessToken,
        debouncedSearchTerm.toUpperCase(),
        itemsOrder
      );
    },
    [debouncedSearchTerm, accessToken, itemsOrder]
  );

  const handleToggleOrder = () => {
    setItemsOrder((prevOrder) => (prevOrder === "ASC" ? "DESC" : "ASC"));
  };

  const handleRemoveProduct = (item) => {
    if (!item?.id) {
      showErrorModal("Item inválido.");
      return;
    }
    setItemToRemove(item);
    setWarningVisible(true);
  };

  const confirmRemoveProduct = async () => {
    if (!itemToRemove) return;

    try {
      await deleteWishListRequest(itemToRemove.id, accessToken);

      handleRefresh();
    } catch (error) {
      console.error("Erro ao remover da wishlist:", error);
      const errorMessage =
        error.response?.data?.error ||
        "Não foi possível remover o item. Tente novamente.";
      showErrorModal(errorMessage);
    } finally {
      setWarningVisible(false);
      setItemToRemove(null);
    }
  };

  return (
    <Background>
      <ModalWarning
        visible={isWarningVisible}
        message="Tem certeza que deseja remover este item da sua lista de desejos?"
        onClose={() => {
          setWarningVisible(false);
          setItemToRemove(null);
        }}
        onConfirm={confirmRemoveProduct}
      />
      <NavBar />

      <View className="flex-row gap-2 m-6 items-center justify-center">
        <Icon name="bookmark-outline" size={30} color={themeColors.header} />
        <Text className="text-white font-bold text-3xl">Lista de Desejos</Text>
      </View>

      <View className="items-center mb-2">
        <SearchBar
          searchValue={searchTerm}
          onSearchChange={setSearchTerm}
          itemsOrder={itemsOrder}
          onToggleOrder={handleToggleOrder}
        />
      </View>

      <ListItems
        key={listKey}
        callbackFetch={fetchWishListCallback}
        CardListRender={({ item }) => (
          <CardWishListItem
            item={item}
            onRemove={() => handleRemoveProduct(item)}
          />
        )}
      />
    </Background>
  );
}
