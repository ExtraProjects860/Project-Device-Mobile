import React, { useState } from "react";
import {
  View,
  Text,
  TextInput,
  TouchableOpacity,
  Modal,
  Pressable,
} from "react-native";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import ModalCheck from "./ModalCheck";
import { useThemeColors } from "../../hooks/useThemeColors.js";
import { useError } from "../../context/ErrorContext.js";
import { useAppContext } from "../../context/AppContext.js";
import { resetPasswordInternalRequest } from "../../lib/authRequests.js";

export default function PasswordChangeModal({ visible, onClose }) {
  const { accessToken } = useAppContext();

  const { showErrorModal } = useError();
  const themeColors = useThemeColors();

  const [showPassword, setShowPassword] = useState(false);
  const [newPassword, setNewPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  const [successMessage, setSuccessMessage] = useState("");
  const [isSuccessVisible, setSuccessVisible] = useState(false);

  const handleCloseSuccessModal = () => {
    setSuccessVisible(false);
    onClose();
  };

  const handlePasswordChange = async () => {
    if (newPassword === "" || confirmPassword === "") {
      return;
    }

    if (newPassword !== confirmPassword) {
      showErrorModal("As senhas não conferem, verifique e tente novamente");
      return;
    }

    try {
      await resetPasswordInternalRequest(newPassword, accessToken);

      setSuccessMessage(
        "Senha alterada com sucesso! Utilize-a no próximo login!"
      );
      setSuccessVisible(true);
    } catch (error) {
      const errorMessage =
        error.response?.data?.error ||
        "Não foi possível realizar a troca da senha. Verifique os dados e tente novamente.";
      showErrorModal(errorMessage);
      console.error("Erro ao trocar a senha:", error);
    }
  };

  return (
    <>
      <ModalCheck
        visible={isSuccessVisible}
        message={successMessage}
        onClose={handleCloseSuccessModal}
      />

      <Modal
        animationType="fade"
        transparent={true}
        visible={visible}
        onRequestClose={onClose}
      >
        <View className="flex-1 justify-center items-center bg-black/40 px-5">
          <View className="bg-light-card dark:bg-dark-card w-full p-6 rounded-2xl">
            {/* Header */}
            <View className="flex-row justify-between items-center mb-6">
              <View className="w-8" />

              <Text className="text-xl font-bold text-light-text-primary dark:text-dark-text-primary">
                Troca de Senha
              </Text>

              <Pressable onPress={onClose} className="p-1">
                <Icon
                  color={
                    themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"
                  }
                  name="close"
                  size={24}
                />
              </Pressable>
            </View>

            {/* Nova Senha */}
            <View className="mb-6">
              <View className="flex-row items-center mb-2">
                <Icon
                  name="lock-outline"
                  size={24}
                  color={
                    themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"
                  }
                />
                <Text className="text-light-text-inverted font-bold pl-2 text-2xl">
                  Nova Senha:
                </Text>
              </View>

              <View className="flex-row items-center bg-white rounded-full px-4 py-1 w-full">
                <TextInput
                  className="-py-1 ml-2 flex-1 text-light-text-secondary font-semibold text-2xl"
                  secureTextEntry={!showPassword}
                  placeholder="*********"
                  placeholderTextColor={
                    themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"
                  }
                  value={newPassword}
                  onChangeText={setNewPassword}
                />

                <TouchableOpacity
                  onPress={() => setShowPassword(!showPassword)}
                >
                  <Icon
                    name={showPassword ? "eye-off-outline" : "eye"}
                    size={20}
                    color="#475569"
                  />
                </TouchableOpacity>
              </View>
            </View>

            {/* Confirmar Senha */}
            <View className="mb-6">
              <View className="flex-row items-center mb-2">
                <Icon
                  name="lock-outline"
                  size={24}
                  color={
                    themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"
                  }
                />
                <Text className="text-light-text-inverted font-bold pl-2 text-2xl">
                  Senha de Confirmação:
                </Text>
              </View>

              <View className="flex-row items-center bg-white rounded-full px-4 py-1 w-full">
                <TextInput
                  className="-py-1 ml-2 flex-1 text-light-text-secondary font-semibold text-2xl"
                  secureTextEntry={!showPassword}
                  placeholder="*********"
                  placeholderTextColor={
                    themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"
                  }
                  value={confirmPassword}
                  onChangeText={setConfirmPassword}
                />

                <TouchableOpacity
                  onPress={() => setShowPassword(!showPassword)}
                >
                  <Icon
                    name={showPassword ? "eye-off-outline" : "eye"}
                    size={20}
                    color="#475569"
                  />
                </TouchableOpacity>
              </View>
            </View>

            {/* Salvar */}
            <TouchableOpacity
              onPress={handlePasswordChange}
              className="bg-light-primary dark:bg-dark-secondary p-3 rounded-full flex-row justify-center items-center"
            >
              <Icon name="check-bold" size={20} color="white" />

              <Text className="text-light-text-inverted text-center font-bold text-base ml-2">
                Salvar
              </Text>
            </TouchableOpacity>
          </View>
        </View>
      </Modal>
    </>
  );
}
