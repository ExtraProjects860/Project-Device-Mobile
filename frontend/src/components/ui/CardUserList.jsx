import React from "react";
import { View, Text, Image, TouchableOpacity } from "react-native";
import { formatDate, formatCPF } from "../../lib/utils.js";
import { useThemeColors } from "../../hooks/useThemeColors.js";
import DefaultAvatar from "../../assets/images/avatar.png";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";

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
  const themeColors = useThemeColors();

  return (
    <View className="flex-row bg-light-card dark:bg-dark-card rounded-xl m-2 overflow-hidden">
      <View className="w-1/3 p-2 items-center justify-center">
        <Image
          className="w-32 h-32 bg-white aspect-square rounded-full"
          source={item?.photo_url ? { uri: item?.photo_url } : DefaultAvatar}
          resizeMode="contain"
        />
      </View>

      <View className="flex-1 p-3 justify-between">
        <View>
          <Text
            className="text-light-text-primary dark:text-dark-text-primary font-bold text-lg"
            numberOfLines={2}
          >
            {item?.name || "Nome não disponível"}
          </Text>

          <Text className="text-light-text-secondary dark:text-dark-text-secondary text-sm font-bold">
            Permissão: {item.role}
          </Text>

          <Text className="text-light-text-secondary dark:text-dark-text-secondary text-sm">
            Empresa: {item?.enterprise || "Sem Empresa"}
          </Text>

          <Text className="text-light-text-secondary dark:text-dark-text-secondary text-sm">
            Registro: {item.register_number} | CPF: {formatCPF(item.cpf)}
          </Text>

          <Text className="text-light-text-secondary dark:text-dark-text-secondary text-sm">
            Entrou em: {formatDate(item.created_at)}
          </Text>
        </View>

        <View className="m-2">
          <TouchableOpacity
            onPress={() => onEdit(item)}
            className="flex-row gap-1 items-center justify-center py-2 px-5 rounded-full bg-light-secondary dark:bg-dark-secondary"
          >
            <Icon
              name="square-edit-outline"
              size={20}
              color={themeColors.header}
            />
            <Text className="text-white font-bold text-center">Editar</Text>
          </TouchableOpacity>
        </View>
      </View>
    </View>
  );
}
