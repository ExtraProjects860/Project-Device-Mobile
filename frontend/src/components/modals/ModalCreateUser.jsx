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
import ModalCheck from "./ModalCheck";
import { createUserRequest } from "../../lib/userRequests.js";
import { useError } from "../../context/ErrorContext.js";
import { useThemeColors } from "../../hooks/useThemeColors.js";
import { useAppContext } from "../../context/AppContext.js";
import User from "../../lib/class/User.js";

import { Picker } from "@react-native-picker/picker";

export default function ModalCreateUser({
  visible,
  onClose,
  onUserCreated,
  roles, 
  enterprises, 
}) {
  const { accessToken } = useAppContext();
  const { showErrorModal } = useError();
  const themeColors = useThemeColors();

  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [cpf, setCpf] = useState("");
  const [registerNumber, setRegisterNumber] = useState("");
  const [roleId, setRoleId] = useState("");
  const [enterpriseId, setEnterpriseId] = useState("");
  const [photoAsset, setPhotoAsset] = useState(null);

  const [successMessage, setSuccessMessage] = useState("");
  const [isSuccessVisible, setSuccessVisible] = useState(false);

  const [errors, setErrors] = useState({});

  const handleCloseSuccessModal = () => {
    setSuccessVisible(false);
    onUserCreated();
    onClose();
  };

  const handleBlur = (fieldName, value) => {
    const error = User.validateField(fieldName, value);
    setErrors((prevErrors) => ({
      ...prevErrors,
      [fieldName]: error,
    }));
  };

  const handleCreateUser = async () => {
    const formData = { name, cpf, email, registerNumber, roleId };
    const validationErrors = User.validateAll(formData);

    const hasErrors = Object.values(validationErrors).some(
      (error) => error !== null
    );

    if (hasErrors) {
      setErrors(validationErrors);
      return;
    }

    try {
      const userData = new User(
        name,
        cpf,
        email,
        registerNumber,
        roleId,
        enterpriseId,
        photoAsset
      );

      await createUserRequest(userData.toJSON(), accessToken);
      setSuccessMessage("Usuário cadastrado com sucesso!");
      setSuccessVisible(true);
    } catch (error) {
      showErrorModal(
        "Não foi possível cadastrar o usuário. Verifique os dados e tente novamente."
      );
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

    const result = await ImagePicker.launchImageLibraryAsync({
      mediaTypes: ImagePicker.MediaTypeOptions.Images,
      allowsEditing: true,
      aspect: [1, 1],
      quality: 1,
    });

    if (!result.canceled) {
      setPhotoAsset(result.assets[0].uri);
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
        visible={visible && !isSuccessVisible}
        onRequestClose={onClose}
      >
        <View className="flex-1 justify-center items-center bg-black/40 px-5">
          <View className="bg-light-card dark:bg-dark-card w-full p-6 rounded-2xl max-h-[90%]">
            {/* Header */}
            <View className="flex-row justify-between items-center mb-6">
              <View className="w-8" />
              <Text className="text-xl font-bold text-light-text-primary dark:text-dark-text-primary">
                Cadastrar Usuário
              </Text>
              <Pressable onPress={onClose} className="p-1">
                <Icon name="close" size={24} color={themeColors.primary} />
              </Pressable>
            </View>

            <ScrollView showsVerticalScrollIndicator={false}>
              {/* Nome do Usuário */}
              <View className="mb-4">
                <Text className="ml-2 text-light-text-primary dark:text-dark-text-primary text-xl font-semibold mb-2">
                  Nome:
                </Text>
                <TextInput
                  className={`bg-gray-soft rounded-lg p-4 text-base text-light-text-primary 
                    ${errors.name ? "border border-red-500" : ""}
                  `}
                  placeholder="Nome Completo"
                  placeholderTextColor={
                    themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"
                  }
                  value={name}
                  onChangeText={setName}
                  onBlur={() => handleBlur("name", name)}
                />
                {errors.name && (
                  <Text className="text-red-500 text-sm ml-2 mt-1">
                    {errors.name}
                  </Text>
                )}
              </View>

              {/* E-mail */}
              <View className="mb-4">
                <Text className="ml-2 text-light-text-primary dark:text-dark-text-primary text-xl font-semibold mb-2">
                  E-mail:
                </Text>
                <TextInput
                  className={`bg-gray-soft rounded-lg p-4 text-base text-light-text-primary
                    ${errors.email ? "border border-red-500" : ""}
                  `}
                  placeholder="exemplo@email.com"
                  placeholderTextColor={
                    themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"
                  }
                  keyboardType="email-address"
                  autoCapitalize="none"
                  value={email}
                  onChangeText={setEmail}
                  onBlur={() => handleBlur("email", email)}
                />
                {errors.email && (
                  <Text className="text-red-500 text-sm ml-2 mt-1">
                    {errors.email}
                  </Text>
                )}
              </View>

              {/* CPF */}
              <View className="mb-4">
                <Text className="ml-2 text-light-text-primary dark:text-dark-text-primary text-xl font-semibold mb-2">
                  CPF:
                </Text>
                <TextInput
                  className={`bg-gray-soft rounded-lg p-4 text-base text-light-text-primary
                    ${errors.cpf ? "border border-red-500" : ""}
                  `}
                  placeholder="000.000.000-00"
                  placeholderTextColor={
                    themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"
                  }
                  keyboardType="numeric"
                  value={cpf}
                  onChangeText={setCpf}
                  onBlur={() => handleBlur("cpf", cpf)}
                />
                {errors.cpf && (
                  <Text className="text-red-500 text-sm ml-2 mt-1">
                    {errors.cpf}
                  </Text>
                )}
              </View>

              {/* Número de Registro */}
              <View className="mb-4">
                <Text className="ml-2 text-light-text-primary dark:text-dark-text-primary text-xl font-semibold mb-2">
                  Nº de Registro:
                </Text>
                <TextInput
                  className={`bg-gray-soft rounded-lg p-4 text-base text-light-text-primary
                    ${errors.registerNumber ? "border border-red-500" : ""}
                  `}
                  placeholder="7 dígitos"
                  placeholderTextColor={
                    themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"
                  }
                  keyboardType="numeric"
                  value={registerNumber}
                  onChangeText={setRegisterNumber}
                  maxLength={7}
                  onBlur={() => handleBlur("registerNumber", registerNumber)}
                />
                {errors.registerNumber && (
                  <Text className="text-red-500 text-sm ml-2 mt-1">
                    {errors.registerNumber}
                  </Text>
                )}
              </View>

              {/* 3. SUBSTITUIR TEXTINPUT POR PICKER (ROLE ID) */}
              <View className="mb-4">
                <Text className="ml-2 text-light-text-primary dark:text-dark-text-primary text-xl font-semibold mb-2">
                  Função:
                </Text>
                <View
                  className={`bg-gray-soft rounded-lg text-base 
                    ${errors.roleId ? "border border-red-500" : ""}
                  `}
                >
                  <Picker
                    selectedValue={roleId}
                    onValueChange={(itemValue) => setRoleId(itemValue)}
                    onBlur={() => handleBlur("roleId", roleId)}
                    dropdownIconColor={
                      themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"
                    }
                    style={{
                      color:
                        themeColors.primary === "#FFFFFF"
                          ? "#000000"
                          : "#000000",
                    }}
                  >
                    <Picker.Item label="Selecione uma função..." value="" />
                    {roles.map((role) => (
                      <Picker.Item
                        key={role.id}
                        label={role.name}
                        value={role.id.toString()}
                      />
                    ))}
                  </Picker>
                </View>
                {errors.roleId && (
                  <Text className="text-red-500 text-sm ml-2 mt-1">
                    {errors.roleId}
                  </Text>
                )}
              </View>

              {/* 4. SUBSTITUIR TEXTINPUT POR PICKER (ENTERPRISE ID) */}
              <View className="mb-4">
                <Text className="ml-2 text-light-text-primary dark:text-dark-text-primary text-xl font-semibold mb-2">
                  Empresa (Opcional):
                </Text>
                <View className="bg-gray-soft rounded-lg text-base">
                  <Picker
                    selectedValue={enterpriseId}
                    onValueChange={(itemValue) => setEnterpriseId(itemValue)}
                    dropdownIconColor={
                      themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"
                    }
                    style={{
                      color:
                        themeColors.primary === "#FFFFFF"
                          ? "#000000"
                          : "#000000",
                    }}
                  >
                    <Picker.Item label="Nenhuma (opcional)" value="" />
                    {enterprises.map((enterprise) => (
                      <Picker.Item
                        key={enterprise.id}
                        label={enterprise.name}
                        value={enterprise.id.toString()}
                      />
                    ))}
                  </Picker>
                </View>
              </View>

              {/* Campo de foto */}
              <View className="mb-6">
                <Text className="ml-2 text-light-text-primary dark:text-dark-text-primary text-xl font-semibold mb-2">
                  Foto (Opcional):
                </Text>

                <TouchableOpacity
                  onPress={pickImage}
                  className="bg-gray-soft h-32 rounded-lg items-center justify-center"
                >
                  {photoAsset ? (
                    <Image
                      source={{ uri: photoAsset }}
                      className="w-full h-full rounded-lg"
                    />
                  ) : (
                    <View className="items-center">
                      <Icon
                        name="image-plus"
                        size={40}
                        color={
                          themeColors.primary === "#FFFFFF"
                            ? "#A0A0A0"
                            : "#6B7280"
                        }
                      />
                      <Text
                        style={{
                          color:
                            themeColors.primary === "#FFFFFF"
                              ? "#A0A0A0"
                              : "#6B7280",
                        }}
                      >
                        Adicionar foto
                      </Text>
                    </View>
                  )}
                </TouchableOpacity>
              </View>
            </ScrollView>

            {/* Cadastrar */}
            <TouchableOpacity
              onPress={handleCreateUser}
              className="bg-light-primary dark:bg-dark-secondary p-3 rounded-full flex-row justify-center items-center mt-4"
            >
              <Icon name="check-bold" size={20} color="white" />
              <Text className="text-light-text-inverted text-center font-bold text-base ml-2">
                Cadastrar
              </Text>
            </TouchableOpacity>
          </View>
        </View>
      </Modal>
    </>
  );
}
