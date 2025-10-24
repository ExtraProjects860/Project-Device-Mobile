import React, { useState, useRef, useEffect } from "react";
import {
  View,
  Text,
  TouchableOpacity,
  Switch,
  Pressable,
  ScrollView,
  Animated,
} from "react-native";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import PasswordChange from "../components/PasswordChange";
import ModalWarning from "./modals/ModalWarning";
import { useNavigateTo } from "../hooks/useNavigateTo";
import { useAppContext } from "../context/AppContext.js";
import { useThemeColors } from "../hooks/useThemeColors.js";

/**
 * Componente responsável pelo Menu do app
 *
 * Recebe 2 atributos
 * O primeiro chamado visible responsável por receber o estado de visibilidade do menu
 * O segundo chamado closeMenu responsável por receber a função que fecha o menu
 *
 * @param {object} props
 * @param {boolean} props.visible
 * @param {function} props.closeMenu
 */
export default function Menu({ visible, closeMenu }) {
  const goTo = useNavigateTo();
  const { isThemeDark, toggleTheme, userData, logout } = useAppContext();
  const themeColors = useThemeColors();

  const [isPasswordModalVisible, setPasswordModalVisible] = useState(false);
  const [isLogoutModalVisible, setLogoutModalVisible] = useState(false);

  const slideAnim = useRef(new Animated.Value(400)).current;
  const fadeAnim = useRef(new Animated.Value(0)).current;

  const isAdmin = userData?.role === "ADMIN" || userData?.role === "SUPERADMIN";

  useEffect(() => {
    Animated.parallel([
      Animated.timing(slideAnim, {
        toValue: visible ? 0 : 400,
        duration: 250,
        useNativeDriver: true,
      }),
      Animated.timing(fadeAnim, {
        toValue: visible ? 1 : 0,
        duration: 250,
        useNativeDriver: true,
      }),
    ]).start();
  }, [visible, slideAnim, fadeAnim]);

  const handleLogoutConfirm = async () => {
    setLogoutModalVisible(false);
    await logout();
    goTo("/login");
  };

  return (
    <>
      <PasswordChange
        visible={isPasswordModalVisible}
        onClose={() => setPasswordModalVisible(false)}
      />

      <ModalWarning
        visible={isLogoutModalVisible}
        message={"Você tem certeza que deseja sair da sua conta?"}
        onClose={() => setLogoutModalVisible(false)}
        onConfirm={handleLogoutConfirm}
      />

      <Animated.View
        style={{ opacity: fadeAnim }}
        pointerEvents={visible ? "auto" : "none"}
        className="absolute inset-0 bg-black/40 z-50"
      >
        <Pressable onPress={closeMenu} className="w-full h-full" />
      </Animated.View>

      {/* Sidebar */}
      <Animated.View
        style={{
          transform: [{ translateX: slideAnim }],
        }}
        pointerEvents={visible ? "auto" : "none"}
        className="absolute top-0 bottom-0 right-0 h-full w-5/6 bg-light-secondary dark:bg-dark-primary p-6 z-50 pt-10 flex flex-col"
      >
        <ScrollView showsVerticalScrollIndicator={false}>
          {/* Header */}
          <View className="flex-row items-center justify-between mb-10 rounded-xl">
            <TouchableOpacity
              onPress={() => goTo("/home")}
              className="p-2 border-2 border-light-text-inverted dark:border-dark-text-primary rounded-xl"
            >
              <Icon name="home-outline" size={30} color={themeColors.header} />
            </TouchableOpacity>

            <TouchableOpacity
              onPress={closeMenu}
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
                  <Text className="ml-2 text-light-primary dark:text-dark-text-primary font-semibold text-base">
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
                <Icon name="lock-outline" size={24} color={themeColors.primary} />
                <Text className="ml-2 text-light-primary dark:text-dark-text-primary font-semibold text-base">
                  Senha
                </Text>
              </TouchableOpacity>

              <TouchableOpacity
                onPress={() => setLogoutModalVisible(true)}
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
                <Icon name="shopping-outline" size={24} color={themeColors.primary} />

                <Text className="ml-2 text-light-primary dark:text-dark-text-primary font-semibold text-base">
                  Produtos
                </Text>
              </TouchableOpacity>

              <TouchableOpacity
                onPress={() => goTo("/wishlist")}
                className="flex-row items-center bg-light-card dark:bg-dark-card rounded-full p-3"
              >
                <Icon name="bookmark-outline" size={24} color={themeColors.primary} />

                <Text className="ml-2 text-light-primary dark:text-dark-text-primary font-semibold text-base">
                  Lista de Desejos
                </Text>
              </TouchableOpacity>
            </View>
          </View>

          {/* Admin */}
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
                    name="account-group-outline"
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
