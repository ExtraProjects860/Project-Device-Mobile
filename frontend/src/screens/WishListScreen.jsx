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

export default function WishListScreen() {
  // TODO essa lista aqui, falta o modal de deletar items
  const themeColors = useThemeColors();
  const { listKey, handleRefresh } = useHandleRefresh();
  const { accessToken } = useAppContext();

  const [selectedWishListItem, setSelectedWishListItem] = useState(null);

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

  const handleRemoveProduct = () => {};

  return (
    <Background>
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
        ref={listKey}
        callbackFetch={fetchWishListCallback}
        CardListRender={({ item }) => (
          <CardWishListItem item={item} onRemove={handleRemoveProduct} />
        )}
      />
    </Background>
  );
}
