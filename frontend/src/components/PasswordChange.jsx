import { useState } from "react";
import {
  View,
  Text,
  TextInput,
  TouchableOpacity,
  Modal,
  Pressable,
} from "react-native";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";

export default function PasswordChangeModal({ visible, onClose }) {
  const [newPassword, setNewPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

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
        <View className="bg-light-card w-full p-6 rounded-2xl">
          {/* Header */}
          <View className="flex-row justify-between items-center mb-6">
            <View className="w-8" />
            <Text className="text-xl font-bold text-light-text-primary">Troca de Senha</Text>
            <Pressable onPress={onClose} className="p-1">
              <Icon name="close" size={24} />
            </Pressable>
          </View>

          {/* Nova Senha */}
          <View className="mb-4">
            <View className="flex-row items-center mb-2">
              <Icon name="lock-outline" size={24} color="gray" />
              <Text className="ml-2 text-gray-strong text-xl font-semibold">
                Nova Senha:
              </Text>
            </View>
            <TextInput
              className="bg-gray-soft rounded-full p-4 text-base"
              placeholder="********"
              secureTextEntry
              value={newPassword}
              onChangeText={setNewPassword}
            />
          </View>

          {/* Confirmar Senha */}
          <View className="mb-6">
            <View className="flex-row items-center mb-2">
              <Icon name="lock-outline" size={24} color="gray" />
              <Text className="ml-2 text-gray-strong text-xl font-semibold">
                Confirmar Senha:
              </Text>
            </View>
            <TextInput
              className="bg-gray-soft rounded-full p-4 text-base"
              placeholder="********"
              secureTextEntry
              value={confirmPassword}
              onChangeText={setConfirmPassword}
            />
          </View>

          {/* Salvar */}
          <TouchableOpacity
            onPress={handlePasswordChange}
            className="bg-light-primary p-3 rounded-full flex-row justify-center items-center"
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
