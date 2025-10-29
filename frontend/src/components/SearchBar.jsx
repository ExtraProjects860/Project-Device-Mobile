import React, { useState } from "react";
import { View, TouchableOpacity, TextInput } from "react-native";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";

export default function SearchBar({
  buttonAdd,
  searchValue,
  onSearchChange,
  itemsOrder,
  onToggleOrder,
}) {
  const orderIcon = itemsOrder === "ASC" ? "sort-ascending" : "sort-descending";

  return (
    <View className="flex-row w-full px-5 items-center gap-x-2">
      <View className="flex-1 flex-row items-center bg-white rounded-full px-4 gap-x-3">
        <TouchableOpacity onPress={onToggleOrder}>
          <Icon name={orderIcon} size={24} />
        </TouchableOpacity>

        <TextInput
          placeholder="Buscar"
          returnKeyType="search"
          value={searchValue}
          onChangeText={onSearchChange}
          className="flex-1 text-base text-light-text-primary"
        />

        <TouchableOpacity className="px-2">
          <Icon name="magnify" size={24} />
        </TouchableOpacity>
      </View>

      <View>{buttonAdd}</View>
    </View>
  );
}
