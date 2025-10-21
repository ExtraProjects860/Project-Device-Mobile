import React from "react";
import { View, Text } from "react-native";
import Background from "../components/ui/Background";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import { NavBar } from "../components/Navbar";
import SearchBar from "../components/SearchBar";
import ButtonAdd from "../components/ui/ButtonAdd";
import { useThemeColors } from "../hooks/useThemeColors.js";

export default function ProductsScreen() {
  // TODO essa lista aqui, pode ser implementada só pra teste de renderização
  // A questão toda que será necessário implementar o login para tokens
  // Manter o usuário na tela atual
  // Fazer com que ele não seja deslogado e verifique se o token dele está ok
  // botão de lembrar de mim
  // logout para limpar
  // adicionar secureStore
  // Implementar tela offline para user
  const themeColors = useThemeColors();

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
        <SearchBar buttonAdd={<ButtonAdd />} />
      </View>
    </Background>
  );
}
