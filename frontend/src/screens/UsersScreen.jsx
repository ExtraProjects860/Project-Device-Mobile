import React from "react";
import { View, Text, Alert } from "react-native";
import Background from "../components/ui/Background";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import { NavBar } from "../components/Navbar";
import { useThemeColors } from "../hooks/useThemeColors.js";
import SearchBar from "../components/SearchBar.jsx";
import { getUsersRequest } from "../lib/UserRequest.js";
import ListItems from "../components/ListItems.jsx";

export default function UsersScreen() {
  const themeColors = useThemeColors();

  // const handleEditUser = (user) => {
  //   console.log("Editar usuário:", user.name);
  //   Alert.alert("A Fazer", `Implementar modal de edição.`);
  // };

  return (
    <Background>
      <NavBar />

      <View className="flex-row gap-2 m-6 items-center justify-center">
        <Icon
          name="account-group-outline"
          size={30}
          color={themeColors.header}
        />
        <Text className="text-white font-bold text-3xl">Usuários</Text>
      </View>

      <View className="items-center mb-2">
        <SearchBar />
      </View>

      {/* Usuários */}
      <ListItems
        callbackFetch={getUsersRequest}
      />
    </Background>
  );
}
