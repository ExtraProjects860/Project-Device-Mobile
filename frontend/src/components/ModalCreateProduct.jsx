import React, { useState } from "react";
import {
  View,
  Text,
  TextInput,
  TouchableOpacity,
  Modal,
  Pressable,
  Image,
  Alert,
} from "react-native";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import * as ImagePicker from "expo-image-picker";

export default function ModalCreate ({ visible, onClose }) {
  const [imagemUri, setImagemUri] = useState(null);
  const handleAddProduct = () => {
    console.log("Produto cadastrado com Sucesso!");
  };

  const pickImage = async () => {
    const { status } = await ImagePicker.requestMediaLibraryPermissionsAsync();
    if (status !== "granted") {
      Alert.alert("Desculpe, precisamos da permissão para acessar suas fotos!");
      return;
    }

    let result = await ImagePicker.launchImageLibraryAsync({
      mediaTypes: ['images'],
      allowsEditing: true,
      aspect: [1, 1], 
      quality: 1,

    });

    if (!result.canceled) {
      setImagemUri(result.assets[0].uri);
    }
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
            <Text className="text-xl font-bold text-light-text-primary">Cadastrar Produto</Text>
            <Pressable onPress={onClose} className="p-1">
              <Icon name="close" size={24} />
            </Pressable>
          </View>

          {/* Nome do produto */}
          <View className="mb-4">
            <View className="flex-row items-center mb-2">
              <Text className="ml-2 text-gray-strong text-xl font-semibold">
                Nome do Produto:
              </Text>
            </View>
            <TextInput
              className="bg-gray-soft rounded-lg p-4 text-base"
              placeholder="Nome do Produto"
            />
          </View>

          {/* Descrição do produto */}
          <View className="mb-6">
            <View className="flex-row items-center mb-2">
              <Text className="ml-2 text-gray-strong text-xl font-semibold">
                Descrição do Produto:
              </Text>
            </View>
            <TextInput
              className="bg-gray-soft rounded-lg p-10 text-base"
              placeholder="Descrição do Produto"
            />
          </View>
          {/* Quantidade */}
          <View className="mb-6">
            <View className="flex-row items-center mb-2">
              <Text className="ml-2 text-gray-strong text-xl font-semibold">
                Quantidade do Produto:
              </Text>
            </View>
            <TextInput
              className="bg-gray-soft rounded-lg p-4 text-base"
              placeholder="Quantidade do Produto"
            />
          </View>
          {/* Marca */}
          <View className="mb-6">
            <View className="flex-row items-center mb-2">
              <Text className="ml-2 text-gray-strong text-xl font-semibold">
                Marca do Produto:
              </Text>
            </View>
            <TextInput
              className="bg-gray-soft rounded-lg p-4 text-base"
              placeholder="Marca do Produto"
            />
          </View>
          {/*Campo de foto*/}
          <View className="mb-6">
            <View className="flex-row items-center mb-2">
              <Text className="ml-2 text-gray-strong text-xl font-semibold">
                Foto:
              </Text>
            </View>
            <TouchableOpacity
              onPress={pickImage}
              className="bg-gray-soft h-32 rounded-lg items-center justify-center"
            >
              {imagemUri ? (
                <Image source={{ uri: imagemUri }} className="w-full h-full rounded-2xl" />
              ) : (
                <View className="items-center">
                  <Icon name="image-plus" size={40} color="#a0a0a0" />
                  <Text className="text-gray-500 mt-2">Adicionar foto</Text>
                </View>
              )}
            </TouchableOpacity>
          </View>

          {/* Salvar */}
          <TouchableOpacity
            onPress={handleAddProduct}
            className="bg-light-primary p-3 rounded-full flex-row justify-center items-center"
          >
            <Icon name="check-bold" size={20} color="white" />
            <Text className="text-light-text-inverted text-center font-bold text-base ml-2">
              Adcionar
            </Text>
          </TouchableOpacity>
        </View>
      </View>
    </Modal>
  );
}
