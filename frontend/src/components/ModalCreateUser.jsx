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
  ScrollView,
} from "react-native";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import * as ImagePicker from "expo-image-picker";
import { createUserRequest } from "../lib/UserRequest.js";

export default function ModalCreateUser({ visible, onClose, onUserCreated }) {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [cpf, setCpf] = useState("");
  const [registerNumber, setRegisterNumber] = useState("");
  const [roleId, setRoleId] = useState("");
  const [enterpriseId, setEnterpriseId] = useState("");
  const [photoUri, setPhotoUri] = useState(null);

  const handleCreateUser = async () => {
    if (!name || !email || !cpf || !registerNumber || !roleId) {
      Alert.alert("Erro", "Por favor, preencha todos os campos obrigatórios.");
      return;
    }

    const temporaryPassword = cpf.replace(/\D/g, "");
    if (temporaryPassword.length !== 11) {
      Alert.alert("Erro", "O CPF deve conter 11 dígitos.");
      return;
    }

    try {
      const userData = {
        name,
        email,
        cpf,
        password: temporaryPassword,
        register_number: parseInt(registerNumber, 10),
        role_id: parseInt(roleId, 10),
        enterprise_id: enterpriseId ? parseInt(enterpriseId, 10) : null,
        photo_url: photoUri,
      };

      await createUserRequest(userData);

      Alert.alert("Sucesso!", "Usuário cadastrado com sucesso!");
      onUserCreated();
      onClose();
    } catch (error) {
      const errorMessage =
        error.response?.data?.error ||
        "Não foi possível cadastrar o usuário. Verifique os dados e tente novamente.";

      showErrorModal(errorMessage);
      console.error("Erro ao criar usuário:", error);
    }
  };

  const pickImage = async () => {
    const { status } = await ImagePicker.requestMediaLibraryPermissionsAsync();
    if (status !== "granted") {
      Alert.alert(
        "Permissão necessária",
        "Desculpe, precisamos da permissão para acessar suas fotos!"
      );
      return;
    }

    let result = await ImagePicker.launchImageLibraryAsync({
      mediaTypes: ImagePicker.MediaTypeOptions.Images,
      allowsEditing: true,
      aspect: [1, 1],
      quality: 1,
    });

    if (!result.canceled) {
      setPhotoUri(result.assets[0].uri);
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
        <View className="bg-light-card w-full p-6 rounded-2xl max-h-[90%]">
          {/* Header */}
          <View className="flex-row justify-between items-center mb-6">
            <View className="w-8" />
            <Text className="text-xl font-bold text-light-text-primary">
              Cadastrar Usuário
            </Text>
            <Pressable onPress={onClose} className="p-1">
              <Icon name="close" size={24} />
            </Pressable>
          </View>

          <ScrollView showsVerticalScrollIndicator={false}>
            {/* Nome do Usuário */}
            <View className="mb-4">
              <Text className="ml-2 text-gray-strong text-xl font-semibold mb-2">
                Nome:
              </Text>
              <TextInput
                className="bg-gray-soft rounded-lg p-4 text-base"
                placeholder="Nome Completo"
                value={name}
                onChangeText={setName}
              />
            </View>

            {/* E-mail */}
            <View className="mb-4">
              <Text className="ml-2 text-gray-strong text-xl font-semibold mb-2">
                E-mail:
              </Text>
              <TextInput
                className="bg-gray-soft rounded-lg p-4 text-base"
                placeholder="exemplo@email.com"
                keyboardType="email-address"
                autoCapitalize="none"
                value={email}
                onChangeText={setEmail}
              />
            </View>

            {/* CPF */}
            <View className="mb-4">
              <Text className="ml-2 text-gray-strong text-xl font-semibold mb-2">
                CPF:
              </Text>
              <TextInput
                className="bg-gray-soft rounded-lg p-4 text-base"
                placeholder="000.000.000-00"
                keyboardType="numeric"
                value={cpf}
                onChangeText={setCpf}
              />
            </View>

            {/* Número de Registro */}
            <View className="mb-4">
              <Text className="ml-2 text-gray-strong text-xl font-semibold mb-2">
                Nº de Registro:
              </Text>
              <TextInput
                className="bg-gray-soft rounded-lg p-4 text-base"
                placeholder="Apenas números"
                keyboardType="numeric"
                value={registerNumber}
                onChangeText={setRegisterNumber}
              />
            </View>

            {/* Role ID */}
            <View className="mb-4">
              <Text className="ml-2 text-gray-strong text-xl font-semibold mb-2">
                ID da Funcionário:
              </Text>
              <TextInput
                className="bg-gray-soft rounded-lg p-4 text-base"
                placeholder="Ex: 1 para Admin, 2 para Usuário"
                keyboardType="numeric"
                value={roleId}
                onChangeText={setRoleId}
              />
            </View>

            {/* Enterprise ID (Opcional) */}
            <View className="mb-4">
              <Text className="ml-2 text-gray-strong text-xl font-semibold mb-2">
                ID da Empresa (Opcional):
              </Text>
              <TextInput
                className="bg-gray-soft rounded-lg p-4 text-base"
                placeholder="Deixe em branco se não aplicável"
                keyboardType="numeric"
                value={enterpriseId}
                onChangeText={setEnterpriseId}
              />
            </View>

            {/* Campo de foto */}
            <View className="mb-6">
              <Text className="ml-2 text-gray-strong text-xl font-semibold mb-2">
                Foto:
              </Text>
              <TouchableOpacity
                onPress={pickImage}
                className="bg-gray-soft h-32 rounded-lg items-center justify-center"
              >
                {photoUri ? (
                  <Image
                    source={{ uri: photoUri }}
                    className="w-full h-full rounded-lg"
                  />
                ) : (
                  <View className="items-center">
                    <Icon name="image-plus" size={40} color="#a0a0a0" />
                    <Text className="text-gray-500 mt-2">Adicionar foto</Text>
                  </View>
                )}
              </TouchableOpacity>
            </View>
          </ScrollView>

          {/* Cadastrar */}
          <TouchableOpacity
            onPress={handleCreateUser}
            className="bg-light-primary p-3 rounded-full flex-row justify-center items-center mt-4"
          >
            <Icon name="check-bold" size={20} color="white" />
            <Text className="text-light-text-inverted text-center font-bold text-base ml-2">
              Cadastrar
            </Text>
          </TouchableOpacity>
        </View>
      </View>
    </Modal>
  );
}
