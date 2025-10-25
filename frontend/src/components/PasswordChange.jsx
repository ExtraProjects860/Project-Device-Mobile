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
import { useThemeColors } from "../hooks/useThemeColors.js";

export default function PasswordChangeModal({ visible, onClose }) {
  const [newPassword, setNewPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  const themeColors = useThemeColors();

  const handlePasswordChange = () => {
    console.log("Senha alterada!");
  };

  return (
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
          <View className="mb-4">
            <View className="flex-row items-center mb-2">
              <Icon
                color={
                  themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"
                }
                name="lock-outline"
                size={24}
              />

              <Text className="ml-2 text-light-text-primary dark:text-dark-text-primary text-xl font-semibold">
                Nova Senha:
              </Text>
            </View>

            <TextInput
              className="bg-gray-soft rounded-full p-4 text-base"
              placeholder="********"
              secureTextEntry
              placeholderTextColor={
                themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"
              }
              value={newPassword}
              onChangeText={setNewPassword}
            />
          </View>

          {/* Confirmar Senha */}
          <View className="mb-6">
            <View className="flex-row items-center mb-2">
              <Icon name="lock-outline" size={24} color="gray" />

              <Text className="ml-2 text-light-text-primary dark:text-dark-text-primary text-xl font-semibold">
                Confirmar Senha:
              </Text>
            </View>

            <TextInput
              className="bg-gray-soft rounded-full p-4 text-base"
              placeholder="********"
              secureTextEntry
              placeholderTextColor={
                themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"
              }
              value={confirmPassword}
              onChangeText={setConfirmPassword}
            />
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
  );
}
