import React from "react";
import { View, Text } from "react-native";
import Background from "../components/ui/Background";
import { NavBar } from "../components/Navbar";

export default function ProductsScreen() {
  return (
    <Background>
      <NavBar />
      <View className="flex-1 items-center justify-center">
        <Text className="text-xl color-white font-bold text-slate-800">
          Esta é a Tela de Produtos
        </Text>
      </View>
    </Background>
  );
}
