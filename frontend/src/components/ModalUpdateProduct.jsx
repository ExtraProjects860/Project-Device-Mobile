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
  Switch
} from "react-native";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import * as ImagePicker from "expo-image-picker";
import { updateProductRequest } from "../lib/ProductsRequest.js";
import { useError } from "../context/ErrorContext.js";
import { useThemeColors } from "../hooks/useThemeColors.js";
import ModalCheck from "./ModalCheck.jsx";

export default function ModalUpdateProduct({
  visible,
  onClose,
  product,
  onProductUpdated,
}) {
  const { showErrorModal } = useError();
  const themeColors = useThemeColors();
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [value, setValue] = useState("");
  const [quantity, setQuantity] = useState("");
  const [promotionAvaliable, setPromotionAvaliable] = useState(false);
  const [discont, setDiscont] = useState("");
  const [avaliable, setAvaliable] = useState(true);
  const [photoUri, setPhotoUri] = useState(null);

  const [isSuccessVisible, setSuccessVisible] = useState(false);
  const [successMessage, setSuccessMessage] = useState("");

  useEffect(() => {
    if (product) {
      setName(product.name || "");
      setDescription(product.description || "");
      setValue(product.value?.toString() || "");
      setQuantity(product.quantity?.toString() || "");
      setPromotionAvaliable(product.is_promotion_avaible ?? false);
      setDiscont(product.discount?.toString() || "");
      setAvaliable(product.is_avaible ?? false);
      setPhotoUri(product.photo_url || null);
    }
  }, [product]);

  const handleCloseSuccessModal = () => {
    setSuccessVisible(false);
    onProductUpdated();
    onClose();
  };

  const handleUpdateProduct = async () => {
    if (
      !name &&
      !description &&
      !value &&
      !quantity &&
      !promotionAvaliable &&
      !discont &&
      !avaliable &&
      !photoUri
    ) {
      showErrorModal("Nenhum campo Vazio para atualizar.");
      return;
    }
    const updatedProductData = {};
    if (name !== product.name && name !== "") updatedProductData.name = name;
    if (description !== product.description && description !== "") updatedProductData.description = description;
    if (parseFloat(value, 10) !== product.value && value !== "") updatedProductData.value = parseFloat(value, 10);
    if (parseInt(quantity, 10) !== product.quantity && quantity !== "") updatedProductData.quantity = parseInt(quantity, 10);
    if (promotionAvaliable !== product.is_promotion_avaible && promotionAvaliable !== "") updatedProductData.is_promotion_avaible = promotionAvaliable;
    if (parseFloat(discont, 10) !== product.discount && discont !== "") updatedProductData.discount = parseFloat(discont, 10);
    if (avaliable !== product.is_avaible && avaliable !== "") updatedProductData.is_avaible = avaliable;
    if (photoUri !== product.photo_url) updatedProductData.photo_url = photoUri;


    if (Object.keys(updatedProductData).length === 0) {
      setSuccessMessage("Nenhum campo foi modificado.");
      setSuccessVisible(true);
      return;
    }

    try {
      await updateProductRequest(product.id, updatedProductData);

      setSuccessMessage("Produto atualizado com sucesso!");
      setSuccessVisible(true);
    } catch (error) {
      const errorMessage =
        error.response?.data?.error ||
        "Não foi possível atualizar o produto. Verifique os dados e tente novamente.";
      showErrorModal(errorMessage);
      console.error("Erro ao atualizar produto:", error);
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
    <>
      <ModalCheck
        visible={isSuccessVisible}
        message={successMessage}
        onClose={handleCloseSuccessModal}
      />
      <Modal
        animationType="fade"
        transparent={true}
        visible={visible}
        onRequestClose={onClose}
      >
        <ScrollView>
          <View className="flex-1 justify-center items-center bg-black/40 px-5">
            <View className="bg-light-card w-full p-6 rounded-2xl">
              {/* Header */}
              <View className="flex-row justify-between items-center mb-6">
                <View className="w-8" />
                <Text className="text-xl font-bold text-light-text-primary">
                  Editar Produto
                </Text>
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
                  value={name}
                  onChangeText={setName}
                />
              </View>

              {/* Descrição do produto */}
              <View className="mb-4">
                <View className="flex-row items-center mb-2">
                  <Text className="ml-2 text-gray-strong text-xl font-semibold">
                    Descrição do Produto:
                  </Text>
                </View>
                <TextInput
                  className="bg-gray-soft rounded-lg p-4 h-24 text-base"
                  placeholder="Descrição do Produto"
                  multiline={true}
                  textAlignVertical="top"
                  value={description}
                  onChangeText={setDescription}
                />
              </View>
              {/* Quantidade */}
              <View className="mb-4">
                <View className="flex-row items-center mb-2">
                  <Text className="ml-2 text-gray-strong text-xl font-semibold">
                    Quantidade do Produto:
                  </Text>
                </View>
                <TextInput
                  className="bg-gray-soft rounded-lg p-4 text-base"
                  placeholder="Quantidade do Produto"
                  value={quantity}
                  onChangeText={setQuantity}
                />
              </View>
              {/* Valor */}
              <View className="mb-4">
                <View className="flex-row items-center mb-2">
                  <Text className="ml-2 text-gray-strong text-xl font-semibold">
                    Valor do Produto:
                  </Text>
                </View>
                <TextInput
                  className="bg-gray-soft rounded-lg p-4 text-base"
                  placeholder="Valor do Produto"
                  value={value}
                  onChangeText={setValue}
                />
              </View>
              {/* Promoção */}
              <View className="mb-4">
                <View className="flex-row items-center mb-2">
                  <Text className="ml-2 text-gray-strong text-xl font-semibold">
                    Está em Promoção?:
                  </Text>
                </View>
                <Switch
                  trackColor={themeColors.switch.track}
                  thumbColor={themeColors.switch.thumb}
                  onValueChange={setPromotionAvaliable}
                  value={promotionAvaliable}
                />
              </View>
              {/* Desconto */}
              <View className="mb-4">
                <View className="flex-row items-center mb-2">
                  <Text className="ml-2 text-gray-strong text-xl font-semibold">
                    Porcentagem de Desconto (%):
                  </Text>
                </View>
                <TextInput
                  className="bg-gray-soft rounded-lg p-4 text-base"
                  placeholder="Ex: 10.4 para 10.4%"
                  value={discont}
                  onChangeText={setDiscont}
                />
              </View>
              {/* Disponibilidade */}
              <View className="mb-4">
                <View className="flex-row items-center mb-2">
                  <Text className="ml-2 text-gray-strong text-xl font-semibold">
                    Produto Disponível?:
                  </Text>
                </View>
                <Switch
                  trackColor={themeColors.switch.track}
                  thumbColor={themeColors.switch.thumb}
                  onValueChange={setAvaliable}
                  value={avaliable}
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

              {/* Salvar */}
              <TouchableOpacity
                onPress={handleUpdateProduct}
                className="bg-light-primary p-3 rounded-full flex-row justify-center items-center"
              >
                <Icon name="check-bold" size={20} color="white" />
                <Text className="text-light-text-inverted text-center font-bold text-base ml-2">
                  Adcionar
                </Text>
              </TouchableOpacity>
            </View>
          </View>
        </ScrollView>
      </Modal>
    </>
  );
}
