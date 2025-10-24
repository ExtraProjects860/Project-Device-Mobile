import React from "react";
import { View, Text, Modal, Pressable } from "react-native";
import Icon from "react-native-vector-icons/Ionicons";

/**
 * Modal de confirmação para Logout.
 *
 * @param {object} props
 * @param {boolean} props.visible
 * @param {function} props.onClose
 * @param {function} props.onConfirm
 */
export default function ModalLogout({ visible, onClose, onConfirm }) {
  return (
    <Modal
      animationType="fade"
      transparent={true}
      visible={visible}
      onRequestClose={onClose}
    >
      <View className="flex-1 justify-center items-center bg-black/50 px-5">
        <View className="w-full max-w-sm bg-light-card dark:bg-dark-card rounded-2xl p-6 items-center shadow-lg">
          <View className="h-14 w-14 justify-center items-center rounded-full bg-red-100 dark:bg-red-900/50 mb-4">
            <Icon name="log-out-outline" size={30} color="#E91D62" />
          </View>
          <Text className="text-xl font-bold text-light-text-primary dark:text-dark-text-primary mb-2">
            Confirmar Saída
          </Text>
          <Text className="text-sm text-light-text-secondary dark:text-dark-text-secondary text-center mb-6">
            Você tem certeza que deseja sair da sua conta?
          </Text>
          <View className="flex-row w-full gap-3">
            <Pressable
              className="flex-1 p-3 rounded-lg border border-gray-300 dark:border-gray-600 bg-white dark:bg-dark-card active:bg-gray-100"
              onPress={onClose}
            >
              <Text className="text-center font-semibold text-light-text-primary dark:text-dark-text-primary">
                Cancelar
              </Text>
            </Pressable>
            <Pressable
              className="flex-1 p-3 rounded-lg bg-light-secondary dark:bg-dark-secondary active:opacity-80"
              onPress={onConfirm}
            >
              <Text className="text-center font-semibold text-white">Sair</Text>
            </Pressable>
          </View>
        </View>
      </View>
    </Modal>
  );
}
