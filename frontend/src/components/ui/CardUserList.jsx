import React from "react";
import { View, Text, TouchableOpacity } from "react-native";

/**
 * Componente responsável pelo card de usuário
 * 
 * Recebe 2 atributos.
 * O primeiro chamado item é o objeto de usuário que é usado para ver as informações exibidas.
 * O segundo chamado onEdit é responsável por receber a função de edicão do item.
 * 
 * @param {object} props
 * @param {object} props.item
 * @param {function} props.onEdit
 */
export default function CardUserList({ item, onEdit }) {
  return (
    <View className="flex-1 flex-col bg-light-card dark:bg-dark-card rounded-xl m-2 overflow-hidden">
      <View className="flex-1 p-3">
        <Text
          className="text-light-text-primary dark:text-dark-text-primary font-bold text-lg"
          numberOfLines={1}
        >
          {item.name}
        </Text>
        <Text className="text-light-text-secondary dark:text-dark-text-secondary text-sm">
          ID: {item.id}
        </Text>
        <Text className="text-light-text-secondary dark:text-dark-text-secondary text-sm">
          Registro: {item.register_number}
        </Text>
      </View>

      <TouchableOpacity
        onPress={() => onEdit(item)}
        className="py-2 px-5 m-3 rounded-full bg-light-secondary dark:bg-dark-secondary"
      >
        <Text className="text-white font-bold text-center">Editar</Text>
      </TouchableOpacity>
    </View>
  );
}
