import React from "react";
import { View, Image, TouchableOpacity, StatusBar } from "react-native";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import { useThemeColors } from "../hooks/useThemeColors.js";
import logo from "../assets/images/logo.png";
import { useAppContext } from "../context/AppContext";

/**
 * Componente responsável pelo cabeçalho principal do app
 *
 * Recebe 1 atributo.
 * O atributo chamdo onMenuPress é responsável por passar a função de abertra de menu ao clicar no icone.
 *
 * @param {object} props
 * @param {function} props.onMenuPress
 */

export default function Header({ onMenuPress }) {
  const themeColors = useThemeColors();
  const { checkInternetConection } = useAppContext();

  const menuIconVisibilit = () => {
    if (checkInternetConection) {
      return (
        <TouchableOpacity onPress={onMenuPress} className="p-[5px]">
          <Icon name="menu" size={40} color={themeColors.header} />
        </TouchableOpacity>
      );
    }
    return null;
  };

  return (
    <View className="flex-row items-center justify-between bg-light-secondary dark:bg-dark-sencondary px-4 py-2 shadow-md pt-10">
      <StatusBar
        barStyle={"light-content"}
        translucent={true}
        className="bg-light-secondary"
      />

      <Image source={logo} className="w-24 h-24" resizeMode="contain" />

      {menuIconVisibilit()}
    </View>
  );
}
