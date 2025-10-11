import React from "react";
import { useCallback } from "react";
import { View, Text } from "react-native";
import Background from "../components/ui/Background";
import { NavBar } from "../components/Navbar";
import PageLoader from "../context/PageLoader.js";

export default function ProductsScreen() {
  const fetchNotices = useCallback(async () => {
    await new Promise((resolve) => setTimeout(resolve, 1000));
  }, []);

  return (
    <PageLoader fetchData={fetchNotices}>
      <Background>
        <NavBar />
        <View className="flex-1 items-center justify-center">
          <Text className="text-xl color-white font-bold text-slate-800">
            Esta é a Tela de Notícias
          </Text>
        </View>
      </Background>
    </PageLoader>
  );
}
