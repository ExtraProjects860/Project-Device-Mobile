import React, { useState, useEffect } from "react";
import {
  View,
  Text,
  TouchableOpacity,
  Switch,
  Dimensions,
  Pressable,
} from "react-native";
import Animated, {
  useSharedValue,
  useAnimatedStyle,
  withTiming,
  Easing,
} from "react-native-reanimated";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import { useNavigateTo } from "../hooks/useNavigateTo";

/**
 * @param {object} props
 * @param {boolean} props.visible
 * @param {function} props.onClose
 */
export default function Menu({ visible, onClose }) {
  const goTo = useNavigateTo();

  const [isThemeDark, setIsThemeDark] = useState(false);
  const toggleTheme = () => setIsThemeDark((previousState) => !previousState);

  const { width: screenWidth } = Dimensions.get("window");
  const translateX = useSharedValue(screenWidth);

  const animatedMenuStyle = useAnimatedStyle(() => {
    return {
      transform: [{ translateX: translateX.value }],
    };
  });

  useEffect(() => {
    if (!visible) {
      // Anima de volta para fora da tela
      translateX.value = withTiming(screenWidth, {
        duration: 200,
        easing: Easing.in(Easing.ease),
      });
      return;
    }
    // Anima para a posição 0 (visível)
    translateX.value = withTiming(0, {
      duration: 200,
      easing: Easing.out(Easing.ease),
    });
  }, [visible]);

  return (
    <>
      {/* Overlay */}
      {visible && (
        <Pressable onPress={onClose} className="absolute inset-0 bg-black/50 z-50" />
      )}

      {/* Sidebar */}
      <Animated.View
        style={animatedMenuStyle}
        className="absolute top-0 bottom-0 right-0 h-full w-5/6 bg-magenta p-6 z-50"
      >
        {/* Header */}
        <View className="flex-row items-center justify-between mb-10 rounded-xl">
          <TouchableOpacity
            onPress={() => goTo("/")}
            className="p-2 border-2 border-white rounded-xl"
          >
            <Icon name="home" size={30} color="white" />
          </TouchableOpacity>
          <TouchableOpacity
            onPress={onClose}
            className="p-2 border-2 border-white rounded-xl"
          >
            <Icon name="arrow-left" size={30} color="white" />
          </TouchableOpacity>
        </View>

        {/* Configurações */}
        <Text className="text-white font-bold text-xl mb-0">Configurações</Text>
        <View className="h-px bg-white my-4" />
        <View className="flex-col gap-y-4 mb-6">
          <View className="flex-row gap-x-2 items-center justify-between bg-white rounded-full pl-3">
            <View className="flex-row items-center">
              <Icon name="weather-night" size={24} color="teal" />
              <Text className="ml-2 text-teal font-semibold text-base">
                Tema
              </Text>
            </View>
            <Switch
              trackColor={{ false: "gray", true: "#b0fffcff" }}
              thumbColor={isThemeDark ? "teal" : "#f4f3f4"}
              onValueChange={toggleTheme}
              value={isThemeDark}
            />
          </View>
          <TouchableOpacity className="flex-row items-center bg-white rounded-full p-3">
            <Icon name="lock" size={24} color="teal" />
            <Text className="ml-2 text-teal font-semibold text-base">
              Senha
            </Text>
          </TouchableOpacity>
          <TouchableOpacity className="flex-row items-center bg-white rounded-full p-3">
            <Icon name="logout" size={24} color="teal" />
            <Text className="ml-2 text-teal font-semibold text-base">
              Logout
            </Text>
          </TouchableOpacity>
        </View>

        {/* Itens */}
        <Text className="text-white font-bold text-xl mb-0">Itens</Text>
        <View className="h-px bg-white my-4" />
        <View className="mb-6">
          <TouchableOpacity
            onPress={() => goTo("/productssss")}
            className="flex-row items-center bg-white rounded-full p-3 mb-3"
          >
            <Icon name="shopping" size={24} color="teal" />
            <Text className="ml-2 text-teal font-semibold text-base">
              Produtos
            </Text>
          </TouchableOpacity>
          <TouchableOpacity onPress={() => goTo("/wishlist")} className="flex-row items-center bg-white rounded-full p-3">
            <Icon name="bookmark" size={24} color="teal" />
            <Text className="ml-2 text-teal font-semibold text-base">
              Lista de Desejos
            </Text>
          </TouchableOpacity>
        </View>

        {/* Outros */}
        <Text className="text-white font-bold text-xl mb-0">Outros</Text>
        <View className="h-px bg-white my-4" />
        <View>
          <TouchableOpacity onPress={() => goTo("/notices")} className="flex-row items-center bg-white rounded-full p-3">
            <Icon name="web" size={24} color="teal" />
            <Text className="ml-2 text-teal font-semibold text-base">
              Notícias
            </Text>
          </TouchableOpacity>
        </View>

        {/* Footer */}
        <View className="absolute bottom-5 left-6 right-6 pt-2">
          <View className="h-px bg-white my-4" />
          <View className="flex-1 items-center">
            <Text className="text-white text-right text-s">
              0.0.0.1v - Design
            </Text>
            <Text className="text-white text-right text-s">
              © Direitos Reservados
            </Text>
          </View>
        </View>
      </Animated.View>
    </>
  );
}
