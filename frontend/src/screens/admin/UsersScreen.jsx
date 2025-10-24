import React, { useState } from "react";
import { View, Text } from "react-native";
import Background from "../../components/ui/Background.jsx";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import { NavBar } from "../../components/Navbar.jsx";
import SearchBar from "../../components/SearchBar.jsx";
import ListItems from "../../components/ListItems.jsx";
import ButtonAdd from "../../components/ui/ButtonAdd.jsx";
import { getUsersRequest } from "../../lib/userRequests.js";
import { useThemeColors } from "../../hooks/useThemeColors.js";
import { useHandleRefresh } from "../../hooks/useHandleRefresh.js";
import CardUserList from "../../components/ui/CardUserList.jsx";
import ModalCreateUser from "../../components/modals/ModalCreateUser.jsx";
import ModalUpdateUser from "../../components/modals/ModalUpdateUser.jsx";

export default function UsersScreen() {
  const themeColors = useThemeColors();
  const { listKey, handleRefresh } = useHandleRefresh();
  
  const [isCreateModalVisible, setCreateModalVisible] = useState(false);
  const [isUpdateModalVisible, setUpdateModalVisible] = useState(false);
  const [selectedUser, setSelectedUser] = useState(null);

  const handleEditUser = (user) => {
    setSelectedUser(user);
    setUpdateModalVisible(true);
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
        onUserUpdated={handleRefresh}
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
              name={"account-plus-outline"}
            />
          }
        />
      </View>

      <ListItems
        key={listKey}
        callbackFetch={getUsersRequest}
        CardListRender={CardListRender}
      />
    </Background>
  );
}
