import React, { useState, useEffect } from "react";
import {
  View,
  Text,
  TouchableOpacity,
  Switch,
  Dimensions,
  Pressable,
  ScrollView,
} from "react-native";
import Animated, {
  useSharedValue,
  useAnimatedStyle,
  withTiming,
  Easing,
} from "react-native-reanimated";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import PasswordChange from "../components/PasswordChange";
import { useNavigateTo } from "../hooks/useNavigateTo";
import { useAppContext } from "../context/AppContext.js";
import { useThemeColors } from "../hooks/useThemeColors.js";

/**
 * Componente responsável pelo Menu do app
 *
 * Recebe 2 atributos
 * O primeiro chamado visible responsável por receber o estado de visibilidade do menu
 * O segundo chamado onClose responsável por receber a função que fecha o menu
 *
 * @param {object} props
 * @param {boolean} props.visible
 * @param {function} props.onClose
 */
export default function Menu({ visible, onClose }) {
  const goTo = useNavigateTo();
  const { isThemeDark, toggleTheme, logout, userData } = useAppContext();
  const themeColors = useThemeColors();

  const [isPasswordModalVisible, setPasswordModalVisible] = useState(false);

  const translateX = useSharedValue(screenWidth);
  const { width: screenWidth } = Dimensions.get("window");

  const isAdmin = userData?.role === "ADMIN" || "SUPERADMIN";

  const animatedMenuStyle = useAnimatedStyle(() => {
    return {
      transform: [{ translateX: translateX.value }],
    };
  });

  // Função para animação do menu aparecendo na tela
  useEffect(() => {
    if (!visible) {
      translateX.value = withTiming(screenWidth, {
        duration: 200,
        easing: Easing.in(Easing.ease),
      });
      return;
    }
    translateX.value = withTiming(0, {
      duration: 200,
      easing: Easing.out(Easing.ease),
    });
  }, [visible, translateX, screenWidth]);

  return (
    <>
      {/* Modal de Alteração de Senha */}
      <PasswordChange
        visible={isPasswordModalVisible}
        onClose={() => setPasswordModalVisible(false)}
      />

      {visible && (
        <Pressable
          onPress={onClose}
          className="absolute inset-0 bg-black/40 z-50"
        />
      )}

      {/* Sidebar */}
      <Animated.View
        style={animatedMenuStyle}
        className="absolute top-0 bottom-0 right-0 h-full w-5/6 bg-light-secondary dark:bg-dark-primary p-6 z-50 pt-10 flex flex-col"
      >
        <ScrollView>
          {/* Header */}
          <View className="flex-row items-center justify-between mb-10 rounded-xl">
            <TouchableOpacity
              onPress={() => goTo("/home")}
              className="p-2 border-2 border-light-text-inverted dark:border-dark-text-primary rounded-xl"
            >
              <Icon name="home" size={30} color={themeColors.header} />
            </TouchableOpacity>

            <TouchableOpacity
              onPress={onClose}
              className="p-2 border-2 border-light-text-inverted dark:border-dark-text-primary rounded-xl"
            >
              <Icon name="arrow-left" size={30} color={themeColors.header} />
            </TouchableOpacity>
          </View>

          {/* Configurações */}
          <View>
            <Text className="text-light-text-inverted dark:text-dark-text-primary font-bold text-xl mb-0">
              Configurações
            </Text>

            <View className="h-px bg-light-text-inverted dark:bg-dark-text-secondary my-4" />

            <View className="flex-col gap-y-4 mb-6">
              <View className="flex-row gap-x-2 items-center justify-between bg-light-card dark:bg-dark-card rounded-full pl-3">
                <View className="flex-row items-center">
                  <Icon
                    name="weather-night"
                    size={24}
                    color={themeColors.primary}
                  />
                  <Text className="ml-2 text-light-text-primary dark:text-dark-text-primary font-semibold text-base">
                    Tema
                  </Text>
                </View>

                <Switch
                  trackColor={themeColors.switch.track}
                  thumbColor={themeColors.switch.thumb}
                  onValueChange={toggleTheme}
                  value={isThemeDark}
                />
              </View>

              <TouchableOpacity
                onPress={() => setPasswordModalVisible(true)}
                className="flex-row items-center bg-light-card dark:bg-dark-card rounded-full p-3"
              >
                <Icon name="lock" size={24} color={themeColors.primary} />
                <Text className="ml-2 text-light-primary dark:text-dark-text-primary font-semibold text-base">
                  Senha
                </Text>
              </TouchableOpacity>

              <TouchableOpacity
                onPress={async () => {
                  await logout();
                  goTo("/login");
                }}
                className="flex-row items-center bg-light-card dark:bg-dark-card rounded-full p-3"
              >
                <Icon name="logout" size={24} color={themeColors.primary} />

                <Text className="ml-2 text-light-primary dark:text-dark-text-primary font-semibold text-base">
                  Logout
                </Text>
              </TouchableOpacity>
            </View>
          </View>

          {/* Itens */}
          <View>
            <Text className="text-light-text-inverted dark:text-dark-text-primary font-bold text-xl mb-0">
              Itens
            </Text>

            <View className="h-px bg-light-text-inverted dark:bg-dark-text-secondary my-4" />

            <View className="mb-6">
              <TouchableOpacity
                onPress={() => goTo("/products")}
                className="flex-row items-center bg-light-card dark:bg-dark-card rounded-full p-3 mb-3"
              >
                <Icon name="shopping" size={24} color={themeColors.primary} />

                <Text className="ml-2 text-light-primary dark:text-dark-text-primary font-semibold text-base">
                  Produtos
                </Text>
              </TouchableOpacity>

              <TouchableOpacity
                onPress={() => goTo("/wishlist")}
                className="flex-row items-center bg-light-card dark:bg-dark-card rounded-full p-3"
              >
                <Icon name="bookmark" size={24} color={themeColors.primary} />

                <Text className="ml-2 text-light-primary dark:text-dark-text-primary font-semibold text-base">
                  Lista de Desejos
                </Text>
              </TouchableOpacity>
            </View>
          </View>

          {/* Seção do Administrador */}
          {isAdmin && (
            <View>
              <Text className="text-light-text-inverted dark:text-dark-text-primary font-bold text-xl mb-0 mt-2">
                Admin
              </Text>

              <View className="h-px bg-light-text-inverted dark:bg-dark-text-secondary my-4" />

              <View className="mb-6">
                <TouchableOpacity
                  onPress={() => goTo("/users")}
                  className="flex-row items-center bg-light-card dark:bg-dark-card rounded-full p-3 mb-3"
                >
                  <Icon
                    name="account-group"
                    size={24}
                    color={themeColors.primary}
                  />

                  <Text className="ml-2 text-light-primary dark:text-dark-text-primary font-semibold text-base">
                    Gerenciar Usuários
                  </Text>
                </TouchableOpacity>

                <TouchableOpacity
                  onPress={() => goTo("/products")}
                  className="flex-row items-center bg-light-card dark:bg-dark-card rounded-full p-3"
                >
                  <Icon
                    name="basket-outline"
                    size={24}
                    color={themeColors.primary}
                  />

                  <Text className="ml-2 text-light-primary dark:text-dark-text-primary font-semibold text-base">
                    Gerenciar Produtos
                  </Text>
                </TouchableOpacity>
              </View>
            </View>
          )}
        </ScrollView>

        {/* Footer */}
        <View>
          <View className="h-px bg-light-text-inverted dark:bg-dark-text-secondary my-4" />

          <View>
            <Text className="text-light-text-inverted dark:text-dark-text-primary text-right text-xs">
              0.0.0.1v - Design
            </Text>

            <Text className="text-light-text-inverted dark:text-dark-text-primary text-right text-xs">
              © Direitos Reservados
            </Text>
          </View>
        </View>
      </Animated.View>
    </>
  );
}
