import React from "react";
import { View, Text, TouchableOpacity } from "react-native";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import { useThemeColors } from "../../hooks/useThemeColors.js";
import { useState } from "react";

/**
 * Componente responsável pelo botão de adicionar, exemplo: botão de adicionar usuário na tela UsersScreen
 *
 * Recebe 2 atributos.
 * O primeiro chamado onPress é responsável por passar a função a ser realizada ao clickar no botão.
 * O segundo chamado name é responsável por receber o nome do Icon a ser utilizado no botão tornado assim facil de se reutilizar em várias telas.
 */

export default function ButtonAdd({ onPress, name }) {
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
