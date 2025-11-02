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
  Switch,
} from "react-native";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import * as ImagePicker from "expo-image-picker";

import ModalCheck from "./ModalCheck";
// Use o nome de arquivo que você tem (productsRequests ou productRequests)
import { createProductRequest } from "../../lib/productsRequests.js";
import { useError } from "../../context/ErrorContext.js";
import { useThemeColors } from "../../hooks/useThemeColors.js";
import { useAppContext } from "../../context/AppContext.js";
// Use o nome de classe que você tem (Products ou Product)
import Product from "../../lib/class/Product.js";

export default function ModalCreateProduct({
  visible,
  onClose,
  onProductCreated,
}) {
  const { accessToken } = useAppContext();
  const { showErrorModal } = useError();
  const themeColors = useThemeColors();

  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [value, setValue] = useState("");
  const [discount, setDiscount] = useState("");
  const [quantity, setQuantity] = useState("");
  const [isAvailable, setIsAvailable] = useState(true);
  const [isPromotionAvailable, setIsPromotionAvailable] = useState(false);
  const [photoUrl, setPhotoUrl] = useState(null);

  const [successMessage, setSuccessMessage] = useState("");
  const [isSuccessVisible, setSuccessVisible] = useState(false);
  const [errors, setErrors] = useState({});

  const handleCloseSuccessModal = () => {
    setSuccessVisible(false);
    onProductCreated();
    onClose();
  };

  const handleBlur = (fieldName, fieldValue) => {
    const error = Product.validateField(fieldName, fieldValue);
    setErrors((prevErrors) => ({
      ...prevErrors,
      [fieldName]: error,
    }));
  };

  const handleCreateProduct = async () => {
    const formData = { name, description, value, quantity, discount };
    const validationErrors = Product.validateAll(formData);

    const hasErrors = Object.values(validationErrors).some(
      (error) => error !== null,
    );

    if (hasErrors) {
      setErrors(validationErrors);
      return;
    }

    try {
      const productData = new Product(
        name,
        description,
        value,
        discount,
        quantity,
        isAvailable,
        isPromotionAvailable,
        null, // Passa null aqui, pois a foto não vai mais no JSON
      );

      // --- INÍCIO DA CORREÇÃO ---
      // Envia o JSON e a photoUrl (URI) como argumentos separados
      await createProductRequest(productData.toJSON(), photoUrl, accessToken);
      // --- FIM DA CORREÇÃO ---

      setSuccessMessage("Produto cadastrado com sucesso!");
      setSuccessVisible(true);
    } catch (error) {
      showErrorModal(
        "Não foi possível cadastrar o produto. Verifique os dados e tente novamente.",
      );
      if (error.response) {
        console.error(
          "Erro 422 - Detalhes da Validação da API:",
          JSON.stringify(error.response.data, null, 2),
        );
      } else {
        console.error("Erro ao criar produto:", error);
      }
    }
  };

  const pickImage = async () => {
    const { status } = await ImagePicker.requestMediaLibraryPermissionsAsync();
    if (status !== "granted") {
      Alert.alert(
        "Permissão necessária",
        "Desculpe, precisamos da permissão para acessar suas fotos!",
      );
      return;
    }

    // --- INÍCIO DA CORREÇÃO ---
    // Corrigido o warning de "MediaTypeOptions"
    const result = await ImagePicker.launchImageLibraryAsync({
      mediaTypes: ImagePicker.MediaType.Images,
      allowsEditing: true,
      aspect: [1, 1],
      quality: 1,
    });
    // --- FIM DA CORREÇÃO ---

    if (!result.canceled) {
      setPhotoUrl(result.assets[0].uri);
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
                Cadastrar Produto
              </Text>
              <Pressable onPress={onClose} className="p-1">
                <Icon name="close" size={24} color={themeColors.primary} />
              </Pressable>
            </View>

            <ScrollView showsVerticalScrollIndicator={false}>
              {/* Nome do Produto */}
              <View className="mb-4">
                <Text className="ml-2 text-light-text-primary dark:text-dark-text-primary text-xl font-semibold mb-2">
                  Nome:
                </Text>
                <TextInput
                  className={`bg-gray-soft rounded-lg p-4 text-base text-light-text-primary 
                    ${errors.name ? "border border-red-500" : ""}
                  `}
                  placeholder="Nome do Produto"
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

              {/* Descrição */}
              <View className="mb-4">
                <Text className="ml-2 text-light-text-primary dark:text-dark-text-primary text-xl font-semibold mb-2">
                  Descrição:
                </Text>
                <TextInput
                  className={`bg-gray-soft rounded-lg p-4 text-base text-light-text-primary h-24
                    ${errors.description ? "border border-red-500" : ""}
                  `}
                  placeholder="Descrição do Produto"
                  placeholderTextColor={
                    themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"
                  }
                  multiline={true}
                  textAlignVertical="top"
                  value={description}
                  onChangeText={setDescription}
                  onBlur={() => handleBlur("description", description)}
                />
                {errors.description && (
                  <Text className="text-red-500 text-sm ml-2 mt-1">
                    {errors.description}
                  </Text>
                )}
              </View>

              {/* Valor */}
              <View className="mb-4">
                <Text className="ml-2 text-light-text-primary dark:text-dark-text-primary text-xl font-semibold mb-2">
                  Valor (R$):
                </Text>
                <TextInput
                  className={`bg-gray-soft rounded-lg p-4 text-base text-light-text-primary 
                    ${errors.value ? "border border-red-500" : ""}
                  `}
                  placeholder="29.99"
                  placeholderTextColor={
                    themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"
                  }
                  keyboardType="numeric"
                  value={value}
                  onChangeText={setValue}
                  onBlur={() => handleBlur("value", value)}
                />
                {errors.value && (
                  <Text className="text-red-500 text-sm ml-2 mt-1">
                    {errors.value}
                  </Text>
                )}
              </View>

              {/* Desconto */}
              <View className="mb-4">
                <Text className="ml-2 text-light-text-primary dark:text-dark-text-primary text-xl font-semibold mb-2">
                  Desconto (Opcional):
                </Text>
                <TextInput
                  className={`bg-gray-soft rounded-lg p-4 text-base text-light-text-primary 
                    ${errors.discount ? "border border-red-500" : ""}
                  `}
                  placeholder="0.1 para 10%"
                  placeholderTextColor={
                    themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"
                  }
                  keyboardType="numeric"
                  value={discount}
                  onChangeText={setDiscount}
                  onBlur={() => handleBlur("discount", discount)}
                />
                {errors.discount && (
                  <Text className="text-red-500 text-sm ml-2 mt-1">
                    {errors.discount}
                  </Text>
                )}
              </View>

              {/* Quantidade */}
              <View className="mb-4">
                <Text className="ml-2 text-light-text-primary dark:text-dark-text-primary text-xl font-semibold mb-2">
                  Quantidade:
                </Text>
                <TextInput
                  className={`bg-gray-soft rounded-lg p-4 text-base text-light-text-primary 
                    ${errors.quantity ? "border border-red-500" : ""}
                  `}
                  placeholder="100"
                  placeholderTextColor={
                    themeColors.primary === "#FFFFFF" ? "#A0A0A0" : "#6B7280"
                  }
                  keyboardType="numeric"
                  value={quantity}
                  onChangeText={setQuantity}
                  onBlur={() => handleBlur("quantity", quantity)}
                />
                {errors.quantity && (
                  <Text className="text-red-500 text-sm ml-2 mt-1">
                    {errors.quantity}
                  </Text>
                )}
              </View>

              {/* Switches para Booleans */}
              <View className="mb-4 flex-row justify-between items-center p-2">
                <Text className="ml-2 text-light-text-primary dark:text-dark-text-primary text-xl font-semibold">
                  Disponível?
                </Text>
                <Switch
                  trackColor={{ false: "#767577", true: themeColors.secondary }}
                  thumbColor={isAvailable ? themeColors.primary : "#f4f3f4"}
                  onValueChange={setIsAvailable}
                  value={isAvailable}
                />
              </View>

              <View className="mb-4 flex-row justify-between items-center p-2">
                <Text className="ml-2 text-light-text-primary dark:text-dark-text-primary text-xl font-semibold">
                  Em promoção?
                </Text>
                <Switch
                  trackColor={{ false: "#767577", true: themeColors.secondary }}
                  thumbColor={
                    isPromotionAvailable ? themeColors.primary : "#f4f3f4"
                  }
                  onValueChange={setIsPromotionAvailable}
                  value={isPromotionAvailable}
                />
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
                  {photoUrl ? (
                    <Image
                      source={{ uri: photoUrl }}
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
              onPress={handleCreateProduct}
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