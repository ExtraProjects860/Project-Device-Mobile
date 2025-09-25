import React from "react";
import { View, Image, TouchableOpacity } from "react-native";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import { useThemeColors } from "../hooks/useThemeColors.js";
import logo from "../assets/images/logo.png"

/**
 *
 * @param {object} props
 * * @param {function} props.onMenuPress
 */

export default function Header({ onMenuPress }) {
  const themeColors = useThemeColors();

  return (
    <View className="flex-row items-center justify-between bg-light-secundary dark:bg-dark-sencondary px-4 py-2 shadow-md pt-10">
      <Image source={logo} className="w-24 h-24" resizeMode="contain" />
      <TouchableOpacity onPress={onMenuPress} className="p-[5px]">
        <Icon name="menu" size={40} color={themeColors.header} />
      </TouchableOpacity>
    </View>
  );
}
