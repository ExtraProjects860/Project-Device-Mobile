import React, { useState, useEffect } from "react";
import {
  View,
  Text,
  TouchableOpacity,
  Image,
  Switch,
  Modal,
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
import logo from "../assets/images/logo.png";

/**
 * @param {object} props
 * @param {boolean} props.visible
 * @param {function} props.onClose
 */
export default function Menu({ visible, onClose }) {
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
    if (visible) {
      // Anima para a posição 0 (visível)
      translateX.value = withTiming(0, {
        duration: 150,
        easing: Easing.out(Easing.ease),
      });
    } else {
      // Anima de volta para fora da tela
      translateX.value = withTiming(screenWidth, {
        duration: 150,
        easing: Easing.in(Easing.ease),
      });
    }
  }, [visible]);

  return (
    <Modal
      animationType="fade"
      transparent={true}
      visible={visible}
      onRequestClose={onClose}
    >
      <Pressable onPress={onClose} className="flex-1 bg-black/50">
        <Animated.View
          onStartShouldSetResponder={() => true}
          style={[animatedMenuStyle]}
          className="absolute top-0 bottom-0 right-0 h-full w-5/6"
        >
          {/* Header */}
          <View className="flex-1 bg-magenta p-6">
            <View className="flex-row items-center justify-between mb-10 rounded-xl">
              <Image source={logo} className="w-24 h-24" resizeMode="contain" />
              <TouchableOpacity onPress={onClose} className="p-2 bg-white rounded-xl">
                <Icon name="arrow-left" size={40} color="#E91D62"/>
              </TouchableOpacity>
            </View>

            {/* Configurações */}
            <Text className="text-white font-bold text-xl mb-0">
              Configurações
            </Text>
            <View className="h-px bg-white my-4" />
            <View className="mb-6">
              <View className="flex-row items-center justify-between bg-white rounded-full pl-3 mb-3">
                <View className="flex-row items-center">
                  <Icon
                    name="weather-night"
                    size={24}
                    color="teal"
                    className="mr-3"
                  />
                  <Text className="text-teal font-semibold text-base">
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
                <Icon name="lock" size={24} color="teal" className="mr-3" />
                <Text className="text-teal font-semibold text-base">Senha</Text>
              </TouchableOpacity>
            </View>

            {/* Itens */}
            <Text className="text-white font-bold text-xl mb-0">Itens</Text>
            <View className="h-px bg-white my-4" />
            <View className="mb-6">
              <TouchableOpacity className="flex-row items-center bg-white rounded-full p-3 mb-3">
                <Icon
                  name="swap-horizontal"
                  size={24}
                  color="teal"
                  className="mr-3"
                />
                <Text className="text-teal font-semibold text-base">
                  Produtos
                </Text>
              </TouchableOpacity>
              <TouchableOpacity className="flex-row items-center bg-white rounded-full p-3">
                <Icon
                  name="bookmark"
                  size={24}
                  color="teal"
                  className="mr-3"
                />
                <Text className="text-teal font-semibold text-base">
                  Lista de Desejos
                </Text>
              </TouchableOpacity>
            </View>

            {/* Outros */}
            <Text className="text-white font-bold text-xl mb-0">Outros</Text>
            <View className="h-px bg-white my-4" />
            <View>
              <TouchableOpacity className="flex-row items-center bg-white rounded-full p-3">
                <Icon name="web" size={24} color="teal" className="mr-3" />
                <Text className="text-teal font-semibold text-base">
                  Notícias
                </Text>
              </TouchableOpacity>
            </View>

            {/* Footer */}
            <View className="absolute bottom-5 left-6 right-6 pt-2">
              <View className="h-px bg-white my-4"/>
              <Text className="text-white text-right text-s">
                0.0.0.1v - Design
              </Text>
              <Text className="text-white text-right text-s">
                © Direitos Reservados
              </Text>
            </View>
          </View>
        </Animated.View>
      </Pressable>
    </Modal>
  );
}
