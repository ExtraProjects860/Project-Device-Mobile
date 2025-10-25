import React, { useState } from "react";
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

export default function ProductsScreen() {
  // TODO essa lista aqui, falta o modal de deletar items
  const themeColors = useThemeColors();
  const { listKey, handleRefresh } = useHandleRefresh();

  const [selectedWishListItem, setSelectedWishListItem] = useState(null);

  const handleRemoveProduct = () => {};

  return (
    <Background>
      <NavBar />

      <View className="flex-row gap-2 m-6 items-center justify-center">
        <Icon name="bookmark-outline" size={30} color={themeColors.header} />
        <Text className="text-white font-bold text-3xl">Lista de Desejos</Text>
      </View>

      <View className="items-center mb-2">
        <SearchBar />
      </View>

      <ListItems
        ref={listKey}
        callbackFetch={getItemsWishListRequest}
        CardListRender={({ item }) => (
          <CardWishListItem item={item} onRemove={handleRemoveProduct} />
        )}
      />
    </Background>
  );
}
