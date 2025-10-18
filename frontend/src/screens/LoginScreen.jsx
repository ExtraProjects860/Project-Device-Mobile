import React, { useState } from "react";
import {
  View,
  Text,
  TextInput,
  TouchableOpacity,
  StatusBar,
  Switch,
} from "react-native";
import { MaterialIcons, Feather } from "@expo/vector-icons";
import Background from "../components/ui/Background";
import Logo from "../components/ui/Logo";
import ModalCheck from "../components/modals/ModalCheck";
import ModalErrors from "../components/modals/ModalErrors";
import { loginRequest } from "../lib/AuthRequest.js";
import { getInfoUserRequest } from "../lib/UserRequest.js";
import { useNavigateTo } from "../hooks/useNavigateTo.js";
import { useAppContext } from "../context/AppContext.js";

export default function LoginScreen() {
  const { updateToken, updateUser } = useAppContext();
  const [showPassword, setShowPassword] = useState(false);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [rememberMe, setRememberMe] = useState(false);

  const [successMessage, setSuccessMessage] = useState("");
  const [isSuccessVisible, setSuccessVisible] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");
  const [isErrorVisible, setErrorVisible] = useState(false);

  const goTo = useNavigateTo();

  const getToken = async () => {
    const userData = {
      email,
      password,
      remember_me: rememberMe,
    };

    const token = await loginRequest(userData);

    if (!token) {
      throw new Error("Não foi possível recuperar o token de acesso")
    } 

    await updateToken(token);
    
    return token;
  };

  const getUserData = async (accessToken) => {
    const userData = await getInfoUserRequest(accessToken);

    if (!userData) {
      throw new Error("Não foi possível recuperar os dados do usuário")
    }

    await updateUser(userData);
  };

  const handleLogin = async () => {
    if (!email || !password) {
      setErrorVisible(true);
      setErrorMessage("Por favor, preencha todos os campos obrigatórios.");
      return;
    }

    try {
      const token = await getToken();
      await getUserData(token);

      setSuccessMessage("Usuário logado com sucesso!");
      setSuccessVisible(true);
    } catch (error) {
      const errorMessage =
        error.response?.data?.error ||
        "Não foi possível realizar login. Verifique os dados e tente novamente.";

      setErrorVisible(true);
      setErrorMessage(errorMessage);
      console.error("Erro ao logar usuário:", errorMessage);
    }
  };

  const handleCloseSuccessModal = () => {
    setSuccessVisible(false);

    goTo("/home");
  };

  const handleCloseErrorModal = () => {
    setErrorVisible(false);
  };

  return (
    <Background>
      <ModalCheck
        visible={isSuccessVisible}
        message={successMessage}
        onClose={handleCloseSuccessModal}
      />

      <ModalErrors
        visible={isErrorVisible}
        message={errorMessage}
        onClose={handleCloseErrorModal}
        onRetry={handleLogin}
      />

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
            Olá! Seja Bem Vindo!
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
              value={email}
              onChangeText={setEmail}
            />
          </View>
        </View>

        {/* Campo Senha */}
        <View className="mb-6">
          <View className="flex-row items-center mb-2">
            <MaterialIcons name="lock" size={24} color="white" />
            <Text className="text-light-text-inverted font-bold pl-2 text-2xl">
              Senha:
            </Text>
          </View>
          <View className="flex-row items-center bg-white rounded-full px-4 py-1 w-5/6">
            <TextInput
              className="-py-1 ml-2 flex-1 text-light-text-secondary font-semibold text-2xl"
              secureTextEntry={!showPassword}
              placeholder="*********"
              placeholderTextColor="gray"
              value={password}
              onChangeText={setPassword}
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

        <View className="mb-6 flex-row items-center">
          <Switch value={rememberMe} onValueChange={setRememberMe} />
          <Text className="text-light-text-inverted font-bold">
            Lembrar de Mim
          </Text>
        </View>

        {/* Botões */}
        <View className="w-full items-center flex-1">
          <TouchableOpacity
            onPress={handleLogin}
            className="mt-10 mb-3 py-2 px-20 bg-light-secondary rounded-full items-center"
          >
            <Text className="text-white text-2xl font-bold">Entrar</Text>
          </TouchableOpacity>

          <TouchableOpacity onPress={() => goTo("/forgot-password")}>
            <Text className="text-white text-s underline underline-offset-1">
              Esqueceu a Senha? Clique Aqui!
            </Text>
          </TouchableOpacity>
        </View>
      </View>
    </Background>
  );
}
