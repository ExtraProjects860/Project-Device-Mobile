import React, { useState, useEffect, useCallback } from "react";
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
import { useAppContext } from "../../context/AppContext.js";
import { useFormOptions } from "../../hooks/useFormOptions.js";

export default function UsersScreen() {
  const themeColors = useThemeColors();
  const { listKey, handleRefresh } = useHandleRefresh();
  const { accessToken } = useAppContext();

  const { roles, enterprises, isLoading: isLoadingOptions } = useFormOptions();

  const [isCreateModalVisible, setCreateModalVisible] = useState(false);
  const [isUpdateModalVisible, setUpdateModalVisible] = useState(false);
  const [selectedUser, setSelectedUser] = useState(null);

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

  const fetchUsersCallback = useCallback(
    (itemsPerPage, currentPage) => {
      return getUsersRequest(
        itemsPerPage,
        currentPage,
        accessToken,
        debouncedSearchTerm.toUpperCase(),
        itemsOrder
      );
    },
    [debouncedSearchTerm, accessToken, itemsOrder]
  );

  const handleEditUser = (user) => {
    setSelectedUser(user);
    setUpdateModalVisible(true);
  };

  const handleToggleOrder = () => {
    setItemsOrder((prevOrder) => (prevOrder === "ASC" ? "DESC" : "ASC"));
  };

  return (
    <Background>
      <ModalCreateUser
        visible={isCreateModalVisible}
        onClose={() => setCreateModalVisible(false)}
        onUserCreated={handleRefresh}
        roles={roles || []}
        enterprises={enterprises || []}
      />
      <ModalUpdateUser
        visible={isUpdateModalVisible}
        onClose={() => setUpdateModalVisible(false)}
        user={selectedUser}
        onUserUpdated={handleRefresh}
        roles={roles || []}
        enterprises={enterprises || []}
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
          searchValue={searchTerm}
          onSearchChange={setSearchTerm}
          itemsOrder={itemsOrder}
          onToggleOrder={handleToggleOrder}
        />
      </View>

      <ListItems
        ref={listKey}
        callbackFetch={fetchUsersCallback}
        CardListRender={({ item }) => (
          <CardUserList item={item} onEdit={handleEditUser} />
        )}
      />
    </Background>
  );
}
