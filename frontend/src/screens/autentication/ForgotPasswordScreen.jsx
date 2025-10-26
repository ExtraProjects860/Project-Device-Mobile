import React from "react";
import {
  View,
  Text,
  TextInput,
  TouchableOpacity,
  StatusBar,
} from "react-native";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import Background from "../../components/ui/Background";
import Logo from "../../components/ui/Logo";
import { useNavigateTo } from "../../hooks/useNavigateTo";
import { useThemeColors } from "../../hooks/useThemeColors.js";

export default function ForgotPasswordScreen() {
  const goTo = useNavigateTo();

  const themeColors = useThemeColors();

  return (
    <Background>
      <StatusBar
        barStyle={"light-content"}
        translucent={true}
        className="bg-light-primary dark:bg-dark-primary"
      />
      <View className="flex-1 justfy-center items-center align-middle mt-16">
        {/* Logo */}
        <Logo className="size-48 mb-4" resizedMode="center" />

        {/* Titulo */}
        <View>
          <Text className="text-light-text-inverted font-bold text-4xl mb-4">
            Etapa 1 - Alterar Senha
          </Text>
        </View>
        <View className="w-96 h-0.5 bg-light-card mb-16" />

        {/* Campo E-mail */}
        <View className="mb-6">
          <View className="flex-row items-center mb-2">
            <Icon
              name="email-outline"
              size={24}
              color={themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"}
            />
            <Text className="text-light-text-inverted font-bold pl-2 text-2xl">
              E-mail:
            </Text>
          </View>
          <View className="flex-row items-center bg-white rounded-full px-4 py-1 w-5/6">
            <TextInput
              className="-py-1 ml-2 flex-1 text-light-text-secondary font-semibold text-2xl"
              placeholder="Exemplo@gmail.com"
              placeholderTextColor={
                themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"
              }
              keyboardType="email-address"
            />
          </View>
        </View>

        {/* Enviar e rodapé */}
        <View className="w-full items-center flex-1">
          <View className="flex mt-6 flex-row gap-4">
            <TouchableOpacity
              onPress={() => goTo("/new-password")}
              className="mb-3 py-2 px-16 bg-light-secondary rounded-full items-center"
            >
              <Text className="text-white text-2xl font-bold">Enviar</Text>
            </TouchableOpacity>
            <TouchableOpacity
              onPress={() => goTo("/login")}
              className="mb-3 py-2 px-16 h-12 bg-light-primary rounded-full items-center"
            >
              <Text className="text-white text-2xl font-bold">Voltar</Text>
            </TouchableOpacity>
          </View>
          <Text className="text-white text-s py-2 px-20">
            Obs: Chegará um código no E-mail cadastrado para proseguir com a
            alteração da senha.
          </Text>
        </View>
      </View>
    </Background>
  );
}
