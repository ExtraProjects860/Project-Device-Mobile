import React, { useState, useEffect } from "react";
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
import { updateUserRequest } from "../../lib/userRequests.js";
import { useError } from "../../context/ErrorContext.js";
import { useThemeColors } from "../../hooks/useThemeColors.js";
import ModalCheck from "./ModalCheck";
import { useAppContext } from "../../context/AppContext.js";
import User from "../../lib/class/User.js";

export default function ModalUpdateUser({
  visible,
  onClose,
  user,
  onUserUpdated,
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
  const [photoUri, setPhotoUri] = useState(null);

  const [isSuccessVisible, setSuccessVisible] = useState(false);
  const [successMessage, setSuccessMessage] = useState("");

  const [errors, setErrors] = useState({});

  useEffect(() => {
    if (user) {
      setName(user.name || "");
      setEmail(user.email || "");
      setCpf(user.cpf || "");
      setRegisterNumber(user.register_number?.toString() || "");
      setRoleId(user.role_id?.toString() || "");
      setEnterpriseId(user.enterprise_id?.toString() || "");
      setPhotoUri(user.photo_url || null);
      setErrors({});
    }
  }, [user]);

  const handleCloseSuccessModal = () => {
    setSuccessVisible(false);
    onUserUpdated();
    onClose();
  };

  const handleBlur = (fieldName, value) => {
    const error = User.validateField(fieldName, value);
    setErrors((prevErrors) => ({
      ...prevErrors,
      [fieldName]: error,
    }));
  };

  const handleUpdateUser = async () => {
    const formData = { name, cpf, email, registerNumber, roleId };
    const validationErrors = User.validateAll(formData);
    const hasErrors = Object.values(validationErrors).some(
      (error) => error !== null
    );

    if (hasErrors) {
      setErrors(validationErrors);
      return;
    }

    const newFormData = { ...formData, enterpriseId, photoUri };
    const updatedUserData = User.getChangedFields(user, newFormData);

    if (Object.keys(updatedUserData).length === 0) {
      setSuccessMessage("Nenhum campo foi modificado.");
      setSuccessVisible(true);
      return;
    }

    try {
      await updateUserRequest(user.id, updatedUserData, accessToken);
      setSuccessMessage("Usuário atualizado com sucesso!");
      setSuccessVisible(true);
    } catch (error) {
      showErrorModal(
        "Não foi possível atualizar o usuário. Verifique os dados e tente novamente."
      );
      console.error("Erro ao atualizar usuário:", error);
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
      setPhotoUri(result.assets[0].uri);
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
                Atualizar Usuário
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
                    themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7D80"
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

              {/* Role ID */}
              <View className="mb-4">
                <Text className="ml-2 text-light-text-primary dark:text-dark-text-primary text-xl font-semibold mb-2">
                  ID da Função:
                </Text>
                <TextInput
                  className={`bg-gray-soft rounded-lg p-4 text-base text-light-text-primary
                    ${errors.roleId ? "border border-red-500" : ""}
                  `}
                  placeholder="Ex: 1 para Admin, 2 para Usuário"
                  placeholderTextColor={
                    themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"
                  }
                  keyboardType="numeric"
                  value={roleId}
                  onChangeText={setRoleId}
                  onBlur={() => handleBlur("roleId", roleId)}
                />
                {errors.roleId && (
                  <Text className="text-red-500 text-sm ml-2 mt-1">
                    {errors.roleId}
                  </Text>
                )}
              </View>

              {/* Enterprise ID (Opcional) */}
              <View className="mb-4">
                <Text className="ml-2 text-light-text-primary dark:text-dark-text-primary text-xl font-semibold mb-2">
                  ID da Empresa (Opcional):
                </Text>
                <TextInput
                  className="bg-gray-soft rounded-lg p-4 text-base text-light-text-primary"
                  placeholder="Deixe em branco se não aplicável"
                  placeholderTextColor={
                    themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"
                  }
                  keyboardType="numeric"
                  value={enterpriseId}
                  onChangeText={setEnterpriseId}
                />
              </View>

              {/* Campo de foto */}
              <View className="mb-6">
                <Text className="ml-2 text-light-text-primary dark:text-dark-text-primary text-xl font-semibold mb-2">
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
                        className="mt-2"
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

            {/* Salvar */}
            <TouchableOpacity
              onPress={handleUpdateUser}
              className="bg-light-primary dark:bg-dark-secondary p-3 rounded-full flex-row justify-center items-center mt-4"
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
