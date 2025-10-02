import React from "react";
import { View, Text, TouchableOpacity } from "react-native";

/**
 * @param {object} props
 * @param {object} props.item
 * @param {function} props.onEdit
 */
export default function ListItems({ item, onEdit }) {
  return (
    <View className="flex-row justify-between items-center p-4 rounded-xl mb-3 mx-4 bg-light-card dark:bg-dark-card">
      <View>
        <Text className="text-light-text-primary dark:text-dark-text-primary font-bold text-base">
          {item.name}
        </Text>
        <Text className="text-light-text-secondary dark:text-dark-text-secondary">
          {item.register_number}
        </Text>
      </View>
      <TouchableOpacity
        onPress={() => onEdit(item)}
        className="py-2 px-5 rounded-full bg-light-secondary dark:bg-dark-secondary"
      >
        <Text className="text-white font-bold">Editar</Text>
      </TouchableOpacity>
    </View>
  );
}
