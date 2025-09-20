import { Text, TouchableOpacity, View, StyleSheet, Alert } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { useState } from "react";
import { useNavigate } from "react-router-native";
import Loading from "../components/Loading";
import Header from "../components/header";
import Background from "../components/background";

export default function HomeScreen() {
  
  const handleMenuPress = () => {
    Alert.alert("Menu", "O menu foi pressionado!");
  };

  return (
    <Background>
      <Header onMenuPress={handleMenuPress} />
      <View className="flex-1 justify-center items-center">
        <Text>Conteúdo da sua tela aqui...</Text>
      </View>
    </Background>
  );
}
