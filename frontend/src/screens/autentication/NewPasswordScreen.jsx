import React, { useState } from "react";
import { useLocation } from "react-router-native";
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
import ModalCheck from "../../components/modals/ModalCheck.jsx";
import { useError } from "../../context/ErrorContext.js";
import { resetPasswordRequest } from "../../lib/authRequests.js";
import { useNavigateTo } from "../../hooks/useNavigateTo";
import { useThemeColors } from "../../hooks/useThemeColors";

export default function NewPasswordScreen() {
  
  const goTo = useNavigateTo();
  const themeColors = useThemeColors();
  const location = useLocation();
  const {email} = location.state;
  const {showErrorModal} = useError();

  const [token, setToken] = useState("");
  const [newPassword, setNewPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  const [showPassword, setShowPassword] = useState(false);
  const [successMessage, setSuccessMessage] = useState("");
  const [isSuccessVisible, setSuccessVisible] = useState(false);

  const handleChangePassword = async () => {
    try {
      const data = await resetPasswordRequest(email, token, newPassword);
      if (data) {
        setSuccessMessage(
          "Senha alterada com sucesso! Por favor, faça login com sua nova senha."
        );
        setSuccessVisible(true);
      }
    } catch (error) {
      showErrorModal(
        "Ocorreu um erro ao tentar alterar a senha. Por favor, tente novamente."
      );
    }
  };


  const handleCloseSuccessModal = () => {
    setSuccessVisible(false);

    goTo("/login");
  };

    const verifyMatchPasswords = () => {
      if (newPassword !== confirmPassword) {
        showErrorModal("As senhas não coincidem. Por favor, tente novamente.");
        return;
      }
      handleChangePassword();
    }
  return (
    <Background>

      <ModalCheck
        visible={isSuccessVisible}
        message={successMessage}
        onClose={handleCloseSuccessModal}
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
        <View className="mb-6">
          <View className="flex-row items-center mb-2">
            <Icon
              name="check"
              size={24}
              color={themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"}
            />
            <Text className="text-light-text-inverted font-bold pl-2 text-2xl">
              Código:
            </Text>
          </View>
          <View className="flex-row items-center bg-white rounded-full px-4 py-1 w-5/6">
            <TextInput
              className="-py-1 ml-2 flex-1 text-light-text-secondary font-semibold text-2xl"
              placeholder="Código"
              placeholderTextColor={
                themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"}
              value ={token}
              onChangeText ={setToken}
            />
          </View>
        </View>

        {/* Nova senha */}
        <View className="mb-6">
          <View className="flex-row items-center mb-2">
            <Icon
              name="lock-outline"
              size={24}
              color={themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"}
            />
            <Text className="text-light-text-inverted font-bold pl-2 text-2xl">
              Nova senha:
            </Text>
          </View>
          <View className="flex-row items-center bg-white rounded-full px-4 py-1 w-5/6">
            <TextInput
              className="-py-1 ml-2 flex-1 text-light-text-secondary font-semibold text-2xl"
              secureTextEntry={!showPassword}
              placeholder="*********"
              placeholderTextColor={
                themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"}
              value ={newPassword}
              onChangeText ={setNewPassword}
            />
            <TouchableOpacity onPress={() => setShowPassword()}>
              <Icon
                name={showPassword ? "eye-off-outline" : "eye"}
                size={20}
                color="#475569"
              />
            </TouchableOpacity>
          </View>
        </View>

        {/* Confirmar senha */}
        <View className="mb-6">
          <View className="flex-row items-center mb-2">
            <Icon
              name="lock-outline"
              size={24}
              color={themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"}
            />
            <Text className="text-light-text-inverted font-bold pl-2 text-2xl">
              Confirmar senha:
            </Text>
          </View>
          <View className="flex-row items-center bg-white rounded-full px-4 py-1 w-5/6">
            <TextInput
              className="-py-1 ml-2 flex-1 text-light-text-secondary font-semibold text-2xl"
              secureTextEntry={!showPassword}
              placeholder="*********"
              placeholderTextColor={
                themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"
              }
              value ={confirmPassword}
              onChangeText ={setConfirmPassword}
            />
            <TouchableOpacity onPress={() => setShowPassword(!showPassword)}>
              <Icon
                name={showPassword ? "eye-off-outline" : "eye"}
                size={20}
                color="#475569"
              />
            </TouchableOpacity>
          </View>
        </View>

        {/* Enviar e rodapé */}
        <View className="w-full items-center flex-1">
          <View className="flex mt-6 flex-row gap-4">
            <TouchableOpacity
              onPress={verifyMatchPasswords}
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
