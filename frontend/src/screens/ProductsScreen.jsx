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
import { getProductsRequest } from "../lib/ProductsRequest.js";

export default function ProductsScreen() {
  const themeColors = useThemeColors();

  return (
    <Background>
      <NavBar />

      <View className="flex-row gap-2 m-6 items-center justify-center">
        <Icon name="shopping-outline" size={30} color={themeColors.header} />
        <Text className="text-white font-bold text-3xl">Produtos</Text>
      </View>

      <View className="items-center mb-2">
        <SearchBar buttonAdd={<ButtonAdd />} />
      </View>

      <ListItems
        callbackFetch={getProductsRequest}
        CardListRender={CardProductList}
      />
    </Background>
  );
}
