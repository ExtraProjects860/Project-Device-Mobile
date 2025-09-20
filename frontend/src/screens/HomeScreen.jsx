import { useState } from "react";
import { Text, View } from "react-native";
import Header from "../components/Header";
import Background from "../components/ui/Background";
import Menu from "../components/Menu";

export default function HomeScreen() {
  const [isMenuVisible, setMenuVisible] = useState(false);

  const handleMenuPress = () => {
    setMenuVisible(true);
  };

  const handleMenuClose = () => {
    setMenuVisible(false);
  };

  return (
    <Background>
      <Header onMenuPress={handleMenuPress} />
      <View className="flex-1 justify-center items-center">
        <Text>Conteúdo da sua tela aqui...</Text>
      </View>

      <Menu visible={isMenuVisible} onClose={handleMenuClose} />
    </Background>
  );
}