import React, { useState, useRef } from "react";
import { View, Text } from "react-native";
import Background from "../components/ui/Background";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import { NavBar } from "../components/Navbar";
import SearchBar from "../components/SearchBar.jsx";
import ListItems from "../components/ListItems.jsx";
import ButtonAdd from "../components/ui/ButtonAdd.jsx";
import { getUsersRequest } from "../lib/UserRequest.js";
import { useThemeColors } from "../hooks/useThemeColors.js";
import CardUserList from "../components/ui/CardUserList.jsx";
import ModalCreateUser from "../components/modals/ModalCreateUser";
import ModalUpdateUser from "../components/modals/ModalUpdateUser";

export default function UsersScreen() {
  const themeColors = useThemeColors();
  const [isCreateModalVisible, setCreateModalVisible] = useState(false);
  const [isUpdateModalVisible, setUpdateModalVisible] = useState(false);
  const [selectedUser, setSelectedUser] = useState(null);
  const listRef = useRef(null);

  const handleEditUser = (user) => {
    setSelectedUser(user);
    setUpdateModalVisible(true);
  };

  const handleRefresh = () => {
    if (listRef.current) {
      listRef.current.refresh();
    }
  };

  const CardListRender = ({ item }) => (
    <CardUserList item={item} onEdit={handleEditUser} />
  );

  return (
    <Background>
      <ModalCreateUser
        visible={isCreateModalVisible}
        onClose={() => setCreateModalVisible(false)}
        onUserCreated={handleRefresh}
      />
      <ModalUpdateUser
        visible={isUpdateModalVisible}
        onClose={() => setUpdateModalVisible(false)}
        user={selectedUser}
        onUserUpdated={() => {
          setUpdateModalVisible(false);
          handleRefresh();
        }}
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
              onPress={() => setCreateModalVisible(true)}
              name={"account-outline"}
            />
          }
        />
      </View>

      <ListItems
        ref={listRef}
        callbackFetch={getUsersRequest}
        CardListRender={CardListRender}
      />
    </Background>
  );
}
