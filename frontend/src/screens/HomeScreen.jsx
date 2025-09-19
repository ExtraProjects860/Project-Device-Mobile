import { Text, TouchableOpacity, View, StyleSheet, Alert } from "react-native"; 
import { SafeAreaView } from "react-native-safe-area-context";
import { useState } from "react";
import { useNavigate } from "react-router-native";
import Loading from "../components/Loading";
import Header from "../components/header";

export default function HomeScreen() {

  const handleMenuPress = () => {
    Alert.alert("Menu", "O menu foi pressionado!"); 
  };

  return (
    <View style={styles.screenContainer}>
      <Header onMenuPress={handleMenuPress} />
      
      <View style={styles.content}>
        <Text>Conteúdo da sua tela aqui...</Text>
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  screenContainer: {
    flex: 1,
    backgroundColor: '#fff',
  },
  content: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
});