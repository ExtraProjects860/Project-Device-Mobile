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
  Switch,
} from "react-native";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import * as ImagePicker from "expo-image-picker";

import ModalCheck from "./ModalCheck";
import { updateProductRequest } from "../../lib/productsRequests.js";
import { useError } from "../../context/ErrorContext.js";
import { useThemeColors } from "../../hooks/useThemeColors.js";
import { useAppContext } from "../../context/AppContext.js";
import Product from "../../lib/class/Product.js";

export default function ModalUpdateProduct({
  visible,
  onClose,
  product,
  onProductUpdated,
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

  const [errors, setErrors] = useState({});
  const [isSuccessVisible, setSuccessVisible] = useState(false);
  const [successMessage, setSuccessMessage] = useState("");
  useEffect(() => {
    if (product) {
      setName(product.name || "");
      setDescription(product.description || "");
      setValue(product.value?.toString() || "0");
      setDiscount(product.discount?.toString() || "0");
      setQuantity(product.quantity?.toString() || "0");
      setIsAvailable(product.is_avaible || false);
      setIsPromotionAvailable(product.is_promotion_avaible || false);
      setPhotoUrl(product.photo_url || null); 
    }
  }, [product]);

  const handleCloseSuccessModal = () => {
    setSuccessVisible(false);
    onProductUpdated();
    onClose();
  };

  const handleBlur = (fieldName, fieldValue) => {
    const error = Product.validateField(fieldName, fieldValue);
    setErrors((prevErrors) => ({
      ...prevErrors,
      [fieldName]: error,
    }));
  };

  const handleUpdateProduct = async () => {
    const newFormData = {
      name,
      description,
      value,
      discount,
      quantity,
      isAvailable,
      isPromotionAvailable,
      photoUrl, 
    };

    const validationErrors = Product.validateAll(newFormData);
    const hasErrors = Object.values(validationErrors).some(
      (error) => error !== null,
    );

    if (hasErrors) {
      setErrors(validationErrors);
      return;
    }
    const updatedData = Product.getChangedFields(product, newFormData);
    let newPhotoUri = null;
    if (photoUrl !== product.photo_url) {
      newPhotoUri = photoUrl;
    }

    let dataToSend = updatedData;
    if (Object.keys(updatedData).length === 0 && newPhotoUri) {
      dataToSend = { name: product.name };
    }
    if (Object.keys(dataToSend).length === 0 && !newPhotoUri) {
      onClose();
      return;
    }

    try {
      console.log("Enviando para updateProductRequest:");
      console.log("ID do Produto:", product.id);
      console.log("Dados (JSON):", dataToSend);
      console.log("Nova URI da Foto:", newPhotoUri);

      await updateProductRequest(
        product.id,
        dataToSend, 
        newPhotoUri,
        accessToken,
      );

      setSuccessMessage("Produto atualizado com sucesso!");
      setSuccessVisible(true);
    } catch (error) {
      showErrorModal(
        "Não foi possível atualizar o produto. Tente novamente.",
      );

      if (error.response) {
        console.error(
          `Erro ${error.response.status} - Detalhes da Validação da API:`,
          JSON.stringify(error.response.data, null, 2),
        );
      } else {
        console.error("Erro ao atualizar produto:", error);
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

    const result = await ImagePicker.launchImageLibraryAsync({
      mediaTypes: ImagePicker.MediaType.Images, 
      allowsEditing: true,
      aspect: [1, 1],
      quality: 1,
    });

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
                Editar Produto
              </Text>
              <Pressable onPress={onClose} className="p-1">
                <Icon name="close" size={24} color={themeColors.primary} />
              </Pressable>
            </View>

            <ScrollView showsVerticalScrollIndicator={false}>
              {/* Nome do produto */}
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
                  Desconto:
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
                  Foto:
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
                        Alterar foto
                      </Text>
                    </View>
                  )}
                </TouchableOpacity>
              </View>
            </ScrollView>

            {/* Salvar */}
            <TouchableOpacity
              onPress={handleUpdateProduct}
              className="bg-light-primary dark:bg-dark-secondary p-3 rounded-full flex-row justify-center items-center mt-4"
            >
              <Icon name="check-bold" size={20} color="white" />
              <Text className="text-light-text-inverted text-center font-bold text-base ml-2">
                Salvar Alterações
              </Text>
            </TouchableOpacity>
          </View>
        </View>
      </Modal>
    </>
  );
}