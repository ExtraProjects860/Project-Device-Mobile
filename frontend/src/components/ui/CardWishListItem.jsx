import React from "react";
import { View, Text, Image, TouchableOpacity } from "react-native";
import { formatPromotion } from "../../lib/utils.js";
import { useThemeColors } from "../../hooks/useThemeColors.js";
import DefaultProduct from "../../assets/images/shopping-bag.png";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";

/**
 * @param {object} props
 * @param {object} props.item
 * @param {function} props.onRemove
 */
export default function CardWishListItem({ item, onRemove }) {
  const themeColors = useThemeColors();

  return (
    <View className="bg-light-card dark:bg-dark-card rounded-xl m-2 overflow-hidden">
      {item.is_promotion_avaible && (
        <View className="bg-light-secondary dark:bg-dark-secondary p-1 items-center justify-center">
          <Text className="font-bold text-white text-sm">% Promoção</Text>
        </View>
      )}

      <View className="flex-row">
        <View className="w-1/3 items-center justify-center">
          <Image
            className="w-32 h-32 bg-white aspect-square rounded-full"
            source={item?.photo_url ? { uri: item.photo_url } : DefaultProduct}
            resizeMode="contain"
          />
        </View>

        <View className="flex-1 p-2 justify-between">
          <View>
            <Text
              className="text-light-text-primary dark:text-dark-text-primary font-bold text-base"
              numberOfLines={2}
            >
              {item?.name || "Nome não disponível"}
            </Text>

            <View className="flex-row flex-wrap items-center">
              {item.is_promotion_avaible ? (
                <>
                  <Text className="line-through font-semibold text-light-text-secondary dark:text-dark-text-secondary">
                    R$ {item.value}
                  </Text>
                  <Text className="text-light-secondary dark:text-dark-secondary font-bold ml-1">
                    | Por: R$ {formatPromotion(item.value, item.discount)}
                  </Text>
                </>
              ) : (
                <Text className="text-light-text-secondary font-semibold dark:text-dark-text-secondary">
                  R$ {item.value}
                </Text>
              )}
            </View>

            <Text
              className="text-light-text-secondary dark:text-dark-text-secondary text-sm"
              numberOfLines={1}
              ellipsizeMode="tail"
            >
              Descrição: {item.description || "Sem descrição no momento"}
            </Text>
          </View>

          <View className="m-2">
            <TouchableOpacity
              onPress={() => onRemove()}
              className="flex-row gap-1 items-center justify-center py-2 px-5 rounded-full bg-light-primary dark:bg-dark-primary"
            >
              <Icon
                name="trash-can-outline"
                size={20}
                color={themeColors.header}
              />
              <Text className="text-white font-bold text-center">
                Remover da Lista
              </Text>
            </TouchableOpacity>
          </View>
        </View>
      </View>
    </View>
  );
}
