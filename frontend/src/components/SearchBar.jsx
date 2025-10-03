import React, { useState } from "react";
import { View, Text, TouchableOpacity, TextInput } from "react-native";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import { useThemeColors } from "../hooks/useThemeColors.js";

export default function SearchBar() {
  const [search, setSearch] = useState();

  const themeColors = useThemeColors();

  return (
    <View className="flex-row w-full px-5 items-center gap-x-2">
      <View className="flex-1 flex-row items-center bg-white rounded-full px-4 gap-x-3">
        <TouchableOpacity>
          <Icon name="menu" size={24} />
        </TouchableOpacity>
        <TextInput
          placeholder="Buscar"
          returnKeyType="search"
          value={search}
          onChangeText={setSearch}
          className="flex-1 text-base text-light-text-primary dark:text-dark-text-primary"
        />
        <TouchableOpacity className="px-2">
          <Icon name="magnify" size={24} />
        </TouchableOpacity>
      </View>

      <View>
        <TouchableOpacity className="flex-row items-center bg-light-secondary dark:bg-dark-sencondary rounded-full p-2 shadow-md">
          <Icon
            name="account-plus-outline"
            size={24}
            color={themeColors.header}
          />
          <Text className="ml-2 text-white font-semibold">Add</Text>
        </TouchableOpacity>
      </View>
    </View>
  );
}
