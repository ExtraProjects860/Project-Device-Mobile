import React from "react";
import { useCallback } from "react";
import { View, Text } from "react-native";
import Background from "../components/ui/Background";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import { NavBar } from "../components/Navbar";
import { useThemeColors } from "../hooks/useThemeColors.js";
import SearchBar from "../components/SearchBar.jsx";
import PageLoader from "../context/PageLoader.js";

export default function UsersScreen() {
  const themeColors = useThemeColors();

  const fetchUsers = useCallback(async (setData) => {
    await new Promise((resolve) => setTimeout(resolve, 1000));
    setData({});
  }, []);

  return (
    <PageLoader fetchData={fetchUsers}>
      {(data) => (
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
          <View className="items-center ">
            <SearchBar />
          </View>
        </Background>
      )}
    </PageLoader>
  );
}
