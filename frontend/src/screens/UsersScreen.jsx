import React from "react";
import { View, Text, Alert } from "react-native";
import Background from "../components/ui/Background";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import { NavBar } from "../components/Navbar";
import SearchBar from "../components/SearchBar.jsx";
import ListItems from "../components/ListItems.jsx";
import ButtonAdd from "../components/ui/ButtonAdd.jsx";
import { getUsersRequest } from "../lib/UserRequest.js";
import { useThemeColors } from "../hooks/useThemeColors.js";
import { useState } from "react";
import CardUserList from "../components/ui/CardUserList.jsx";
import ModalCreate from "../components/ModalCreateUser.jsx";

export default function UsersScreen() {
  const themeColors = useThemeColors();
  const [isCreateProductModalVisible, setCreateUserVisible] =
    useState(false);
  // const handleEditUser = (user) => {
  //   console.log("Editar usuário:", user.name);
  //   Alert.alert("A Fazer", `Implementar modal de edição.`);
  // };

  return (
    <Background>
      <ModalCreate
        visible={isCreateProductModalVisible}
        onClose={() => setCreateProductVisible(false)}
      />
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
        <SearchBar
          buttonAdd={
            <ButtonAdd
              onPress={() => setCreateUserVisible(true)}
              name={"account-outline"}
            />
          }
        />
      </View>

      {/* Usuários */}
      <ListItems
        callbackFetch={getUsersRequest}
        CardListRender={CardUserList}
      />
    </Background>
  );
}
