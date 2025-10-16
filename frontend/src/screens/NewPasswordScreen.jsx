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
            Etapa 2 - Alterar Senha
          </Text>
        </View>
        <View className="w-96 h-0.5 bg-light-card mb-2" />

        {/* Codigo */}
        <View>
          <View className="flex-row items-center mb-2">
            <MaterialIcons name="check" size={24} color="white" />
            <Text className="text-light-text-inverted font-bold pl-2 text-2xl">
              Código:
            </Text>
          </View>
          <View className="flex-row items-center bg-white rounded-full px-4 py-1 w-5/6">
            <TextInput
              className="-py-1 ml-2 flex-1 text-light-text-secondary font-semibold text-2xl"
              placeholder="Código"
              placeholderTextColor="gray"
            />
          </View>
        </View>
        {/* Nova senha */}
        <View>
          <View className="flex-row items-center mb-2">
            <MaterialIcons name="lock" size={24} color="white" />
            <Text className="text-light-text-inverted font-bold pl-2 text-2xl">
              Nova senha:
            </Text>
          </View>
          <View className="flex-row items-center bg-white rounded-full px-4 py-1 w-5/6">
            <TextInput
              className="-py-1 ml-2 flex-1 text-light-text-secondary font-semibold text-2xl"
              secureTextEntry={!showPassword}
              placeholder="*********"
              placeholderTextColor="gray"
            />
            <TouchableOpacity onPress={() => setShowPassword(!showPassword)}>
              <Feather
                name={showPassword ? "eye-off" : "eye"}
                size={20}
                color="#475569"
              />
            </TouchableOpacity>
          </View>
        </View>
        {/* Confirmar senha */}
        <View>
          <View className="flex-row items-center mb-2">
            <MaterialIcons name="lock" size={24} color="white" />
            <Text className="text-light-text-inverted font-bold pl-2 text-2xl">
              Confirmar senha:
            </Text>
          </View>
          <View className="flex-row items-center bg-white rounded-full px-4 py-1 w-5/6">
            <TextInput
              className="-py-1 ml-2 flex-1 text-light-text-secondary font-semibold text-2xl"
              secureTextEntry={!showPassword}
              placeholder="*********"
              placeholderTextColor="gray"
            />
            <TouchableOpacity onPress={() => setShowPassword(!showPassword)}>
              <Feather
                name={showPassword ? "eye-off" : "eye"}
                size={20}
                color="#475569"
              />
            </TouchableOpacity>
          </View>
        </View>
        {/* Enviar e rodapé */}
        <View className="w-full items-center flex-1">
          <TouchableOpacity
            onPress={() => goTo("/")}
            className="mt-2 py-2 px-20 bg-light-secondary rounded-full items-center"
          >
            <Text className="text-white text-2xl font-bold">ENVIAR</Text>
          </TouchableOpacity>
            <Text className="text-white text-s py-2 px-20"> {/* arrumar cores */}
              Digite o código enviado para o e-mail e defina sua nova senha.
            </Text>
        </View>
      </View>
    </Background>
  );
}