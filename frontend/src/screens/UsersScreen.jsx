import React, { useCallback } from "react";
import { View, Text, FlatList, Alert } from "react-native";
import Background from "../components/ui/Background";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import { NavBar } from "../components/Navbar";
import { useThemeColors } from "../hooks/useThemeColors.js";
import SearchBar from "../components/SearchBar.jsx";
import PageLoader from "../context/PageLoader.js";
import UserRequest from "../lib/UserRequest.js";
import ListItems from "../components/ListItems.jsx";

export default function UsersScreen() {
  const themeColors = useThemeColors();

  const fetchUsers = useCallback(async (setData) => {
    try {
      const usersData = await UserRequest.GetUsersRequest();
      setData(usersData);
    } catch (error) {
      console.error("Erro ao buscar usuários:", error);
      Alert.alert("Erro", "Não foi possível carregar a lista de usuários.");
      setData({ data: [] });
    }
  }, []);

  const handleEditUser = (user) => {
    console.log("Editar usuário:", user.name);
    Alert.alert("A Fazer", `Implementar modal de edição.`);
  };

  return (
    <PageLoader fetchData={fetchUsers}>
      {(apiResponse) => (
        <Background>
          <NavBar />

          {/* Título */}
          <View className="flex-row gap-2 m-6 items-center justify-center">
            <Icon
              name="account-group-outline"
              size={30}
              color={themeColors.header}
            />
            <Text className="text-white font-bold text-3xl">Usuários</Text>
          </View>

          {/* Barra de Busca */}
          <View className="items-center mb-4">
            <SearchBar />
          </View>

          {/* Usuários */}
          <FlatList
            data={apiResponse?.data || []}
            keyExtractor={(item) => item.id.toString()}
            renderItem={({ item }) => (
              <ListItems item={item} onEdit={handleEditUser} />
            )}
            ListEmptyComponent={
              <View className="flex-1 items-center justify-center mt-20">
                <Text className="text-gray-400">
                  Nenhum usuário encontrado.
                </Text>
              </View>
            }
          />
        </Background>
      )}
    </PageLoader>
  );
}
