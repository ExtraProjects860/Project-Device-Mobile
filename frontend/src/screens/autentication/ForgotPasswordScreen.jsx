import React, { useState } from "react";
import { View, Text, TextInput, TouchableOpacity } from "react-native";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import Background from "../../components/ui/Background";
import Logo from "../../components/ui/Logo";
import ModalCheck from "../../components/modals/ModalCheck.jsx";
import { useError } from "../../context/ErrorContext.js";
import { requestToken } from "../../lib/authRequests.js";
import { useNavigateTo } from "../../hooks/useNavigateTo";
import { useThemeColors } from "../../hooks/useThemeColors.js";

export default function ForgotPasswordScreen() {
  const goTo = useNavigateTo();
  const [email, setEmail] = useState("");
  const [erro, setErro] = useState("");
  const [successMessage, setSuccessMessage] = useState("");
  const [isSuccessVisible, setSuccessVisible] = useState(false);
  const { showErrorModal } = useError();
  const themeColors = useThemeColors();

  const handleChangeEmail = async () => {
    try {
      if (!handleVerifyEmail()) return;
      const data = await requestToken(email);
      if (data) {
        setSuccessMessage(
          "E-mail enviado com sucesso! Por favor, verifique sua caixa de entrada."
        );
        setSuccessVisible(true);
      }
    } catch (error) {
      /* mensagem de erro não bem detalhada para o usuario*/
      showErrorModal(
        `Ocorreu um erro ao tentar enviar o e-mail. Por favor, tente novamente. ${error}`
      );
    }
  };

  const handleCloseSuccessModal = () => {
    setSuccessVisible(false);

    goTo("/new-password", { email });
  };

  const handleVerifyEmail = () => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!email.match(emailRegex)) {
      setErro("Por favor, insira um e-mail válido.");
      return false;
    }
    return true;
  };

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
              color="#fff"
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
              value={email}
              onChangeText={setEmail}
            />
          </View>
          <Text className="ml-2 text-light-text-inverted dark:text-dark-text-primary text-xl font-semibold mb-2">
            {erro}
          </Text>
        </View>

        {/* Enviar e rodapé */}
        <View className="w-full items-center flex-1">
          <View className="flex mt-6 flex-row gap-4">
            <TouchableOpacity
              onPress={() => goTo("/login")}
              className="mb-3 py-2 px-16 bg-light-text-inverted rounded-full items-center"
            >
              <Text className="text-light-text-primary text-2xl font-bold">
                Voltar
              </Text>
            </TouchableOpacity>
            <TouchableOpacity
              onPress={() => handleChangeEmail()}
              className="mb-3 py-2 px-16 bg-light-secondary dark:bg-light-secondary rounded-full items-center"
            >
              <Text className="text-light-text-inverted text-2xl font-bold">
                Enviar
              </Text>
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
