import React, { useState } from "react";
import {
  View,
  Text,
  TextInput,
  TouchableOpacity,
  StatusBar,
} from "react-native";
import { MaterialIcons, Feather } from "@expo/vector-icons";
import Background from "../components/ui/Background";
import Logo from "../components/ui/Logo";
import { useNavigateTo } from "../hooks/useNavigateTo";

export default function LoginScreen() {
  const [showPassword, setShowPassword] = useState(false);
  const goTo = useNavigateTo();
  
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
            <MaterialIcons name="email" size={24} color="white" />
            <Text className="text-light-text-inverted font-bold pl-2 text-2xl">
              E-mail:
            </Text>
          </View>
          <View className="flex-row items-center bg-white rounded-full px-4 py-1 w-5/6">
            <TextInput
              className="-py-1 ml-2 flex-1 text-light-text-secondary font-semibold text-2xl"
              placeholder="Exemplo@gmail.com"
              placeholderTextColor="gray"
              keyboardType="email-address"
            />
          </View>
        </View>

        {/* Enviar e rodapé */}
        <View className="w-full items-center flex-1">
          <TouchableOpacity
            onPress={() => goTo("/new-password")}
            className="mt-20 mb-3 py-2 px-20 bg-light-secondary rounded-full items-center"
          >
            <Text className="text-white text-2xl font-bold">ENVIAR</Text>
          </TouchableOpacity>
            <Text className="text-white text-s py-2 px-20">
              Obs.: Chegará um código no E-mail cadastrado para proseguir com a alteração da senha.
            </Text>
        </View>
      </View>
    </Background>
  );
}
