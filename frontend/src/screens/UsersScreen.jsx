import React from "react";
import { useCallback } from "react";
import { View, Text } from "react-native";
import Background from "../components/ui/Background";
import { NavBar } from "../components/Navbar";
import { useNavigateTo } from "../hooks/useNavigateTo.js";
import PageLoader from "../components/PageLoader";

export default function UsersScreen() {
  const fetchUsers = useCallback(async (setData) => {
    await new Promise((resolve) => setTimeout(resolve, 1000));
    setData({});
  }, []);

  return (
    <PageLoader fetchData={fetchUsers}>
      {(data) => (
        <Background>
          <NavBar />
          <View className="flex-1 items-center justify-center">
            <Text className="text-xl color-white font-bold text-slate-800">
              Esta é a Tela de Usuários
            </Text>
          </View>
        </Background>
      )}
    </PageLoader>
  );
}
