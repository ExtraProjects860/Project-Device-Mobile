import React, { useCallback } from "react";
import { Text, View, Image } from "react-native";
import Background from "../components/ui/Background";
import { NavBar } from "../components/Navbar";
import avatar from "../assets/images/avatar.png";
import Barcode from "@kichiyaki/react-native-barcode-generator";
import Icon from "react-native-vector-icons/FontAwesome";
import PageLoader from "../components/PageLoader";

export default function HomeScreen() {
  const fetchHomeData = useCallback(async (setData) => {
    await new Promise((resolve) => setTimeout(resolve, 1000));
    setData({});
  }, []);

  return (
    <PageLoader fetchData={fetchHomeData}>
      {(data) => (
        <Background>
          <NavBar />
          <View className="flex-col items-center justify-center">
            {/* Nome Cliente */}
            <View className="mt-4 mb-5">
              <Text className="text-white font-bold text-4xl">
                Fulano de Tal
              </Text>
            </View>
            <View className="w-10/12 bg-white rounded-3xl p-12 items-center shadow-lg">
              {/* Foto */}
              <View className="w-60 h-60 rounded-full bg-teal p-2 mb-5">
                <Image source={avatar} className="w-full h-full rounded-full" />
              </View>

              {/* Informações Cliente */}
              <View className="items-center mt-4">
                <Text className="text-2xl font-bold text-slate-800">
                  Parceiros S.A.
                </Text>
                <Text className="mt-1 text-lg text-slate-600">
                  111.111.111-11
                </Text>
              </View>

              <View className="w-full h-1 bg-teal my-6" />

              {/* Codigo de Barras */}
              <View className="w-full px-4">
                <Barcode
                  value="1231231"
                  format="EAN8"
                  width={4}
                  height={70}
                  lineColor="black"
                  background="white"
                />
              </View>
            </View>
          </View>

          {/* Rodapé */}
          <View className="absolute bottom-0 w-full h-80 justify-end items-end p-4">
            <Text className="text-white text-right text-2xl font-bold">
              Drogaria Jocy
            </Text>
            <Text className="text-white text-right text-lg">
              Rua Nelson Viana, 652 - Centro, Três Rios
            </Text>
            <View className="flex-row items-center text-lg">
              <Icon name="whatsapp" size={15} color="white" />
              <Text className="text-white ml-2">(24) 99255-6024</Text>
            </View>
          </View>
        </Background>
      )}
    </PageLoader>
  );
}
