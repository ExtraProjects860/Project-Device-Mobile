import { Text, View, Alert } from "react-native";
import Header from "../components/header";
import Background from "../components/ui/Background";

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
