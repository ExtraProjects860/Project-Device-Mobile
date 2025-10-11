import React from "react";
import { View, Text, TouchableOpacity } from "react-native";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import { useThemeColors } from "../../hooks/useThemeColors.js";
import { useState } from "react";

export default function ButtonAdd({ onPress, name, text }) {
  const themeColors = useThemeColors();

  return (
    <TouchableOpacity
      className="flex-row items-center bg-light-secondary dark:bg-dark-sencondary rounded-full p-2 shadow-md"
      onPress={onPress}
    >
      <Icon name={name} size={24} color={themeColors.header} />
      <Text className="ml-2 text-white font-semibold">Add</Text>
    </TouchableOpacity>
  );
}
