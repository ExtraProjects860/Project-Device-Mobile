import { useState } from "react";
import { Text, View, Image } from "react-native";
import Header from "../components/header";
import Background from "../components/ui/Background";
import Menu from "../components/Menu";
import avatar from "../assets/images/avatar.png";
import Barcode from "@kichiyaki/react-native-barcode-generator";
import Icon from "react-native-vector-icons/FontAwesome";

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
      <View className="flex-col items-center justify-center">
        {/* Nome Cliente */}
        <View className="mt-10 mb-5">
          <Text className="text-white font-bold text-5xl">Fulano de Tal</Text>
        </View>
        <View className="w-11/12 bg-white rounded-3xl p-16 items-center shadow-lg">
          {/* Foto */}
          <View className="w-60 h-60 rounded-full bg-teal p-2 mb-5">
            <Image source={avatar} className="w-full h-full rounded-full" />
          </View>

          {/* Informações Cliente */}
          <View className="items-center mt-4">
            <Text className="text-2xl font-bold text-slate-800">
              Parceiros S.A.
            </Text>
            <Text className="mt-1 text-lg text-slate-600">111.111.111-11</Text>
          </View>

          <View className="w-full h-1 bg-teal my-6" />

          {/* Codigo de Barras */}
          <View className="w-full px-4">
            <Barcode
              value="1231231" 
              format="EAN8" 
              width={4}
              height={70}
              lineColor="#000000"
              background="#FFFFFF"
            />
          </View>
        </View>
      </View>

      {/* Rodapé */}
      <View className="absolute bottom-0 w-full h-80 justify-end items-end p-4">
        <Text className="text-white text-right text-2xl font-bold">Drogaria Jocy</Text>
        <Text className="text-white text-right text-lg">
          Rua Nelson Viana, 652 - Centro, Três Rios
        </Text>
        <View className="flex-row items-center text-lg">
          <Icon name="whatsapp" size={15} color="white" />
          <Text className="text-white ml-2">(24) 99255-6024</Text>
        </View>
      </View>
      <Menu visible={isMenuVisible} onClose={handleMenuClose} />
    </Background>
  );
}
