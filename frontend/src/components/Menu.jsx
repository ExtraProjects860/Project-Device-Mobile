import React, { useState } from "react";
import { View, Text, TouchableOpacity, Image, Switch, Modal } from "react-native";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import logo from "../assets/images/logo.png"; 

/**
 * @param {object} props
 * @param {boolean} props.visible
 * @param {function} props.onClose
 */
export default function Menu({ visible, onClose }) {
  const [isThemeDark, setIsThemeDark] = useState(false);
  const toggleTheme = () => setIsThemeDark(previousState => !previousState);

  return (
    <Modal
      animationType="slide"
      transparent={false}
      visible={visible}
      onRequestClose={onClose}
    >
        {/* Header */}
      <View className="flex-1 bg-magenta p-6">
        <View className="flex-row items-center justify-between mb-10">
          <Image source={logo} className="w-24 h-24" resizeMode="contain" />
          <TouchableOpacity onPress={onClose} className="p-2">
            <Icon name="arrow-left" size={30} color="white" />
          </TouchableOpacity>
        </View>
        
        {/* Configurações */}
        <Text className="text-white font-bold text-xl mb-4">Configurações</Text>
        <View className="mb-6">
          <View className="flex-row items-center justify-between bg-white rounded-full p-3 mb-3">
            <View className="flex-row items-center">
              <Icon name="weather-night" size={24} color="#087975" className="mr-3" />
              <Text className="text-teal font-semibold text-base">Tema</Text>
            </View>
            <Switch
              trackColor={{ false: "#767577", true: "#b0fffcff" }}
              thumbColor={isThemeDark ? "#087975" : "#f4f3f4"}
              onValueChange={toggleTheme}
              value={isThemeDark}
            />
          </View>
          <TouchableOpacity className="flex-row items-center bg-white rounded-full p-3">
            <Icon name="lock" size={24} color="#087975" className="mr-3" />
            <Text className="text-teal font-semibold text-base">Senha</Text>
          </TouchableOpacity>
        </View>

        {/* Itens */}
        <Text className="text-white font-bold text-xl mb-4">Itens</Text>
        <View className="mb-6">
          <TouchableOpacity className="flex-row items-center bg-white rounded-full p-3 mb-3">
            <Icon name="swap-horizontal" size={24} color="#087975" className="mr-3" />
            <Text className="text-teal font-semibold text-base">Produtos</Text>
          </TouchableOpacity>
          <TouchableOpacity className="flex-row items-center bg-white rounded-full p-3">
            <Icon name="bookmark" size={24} color="#087975" className="mr-3" />
            <Text className="text-teal font-semibold text-base">Lista de Desejos</Text>
          </TouchableOpacity>
        </View>

        {/* Outros */}
        <Text className="text-white font-bold text-xl mb-4">Outros</Text>
        <View>
          <TouchableOpacity className="flex-row items-center bg-white rounded-full p-3">
            <Icon name="web" size={24} color="#087975" className="mr-3" />
            <Text className="text-teal font-semibold text-base">Notícias</Text>
          </TouchableOpacity>
        </View>

        {/* Footer */}
        <View className="absolute bottom-5 left-6 right-6 border-t border-white/50 pt-2">
          <Text className="text-white text-center text-xs">0.0.0.1v - Design</Text>
          <Text className="text-white text-center text-xs">© Direitos Reservados</Text>
        </View>
      </View>
    </Modal>
  );
}