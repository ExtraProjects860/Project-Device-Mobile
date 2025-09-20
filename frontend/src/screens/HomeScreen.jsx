import React from "react";
import { Text, View, StyleSheet, Alert } from "react-native";
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
}

const styles = StyleSheet.create({
  content: {
    alignItems: "center",
    flex: 1,
    justifyContent: "center",
  },
  screenContainer: {
    backgroundColor: "#fff",
    flex: 1,
  },
});
