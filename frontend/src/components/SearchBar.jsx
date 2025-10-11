import React, { useState } from "react";
import { View, TouchableOpacity, TextInput } from "react-native";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";

export default function SearchBar({ buttonAdd }) {
  const [search, setSearch] = useState();

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
      <View>{buttonAdd}</View>
    </View>
  );
}
