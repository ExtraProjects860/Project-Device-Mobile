import React from "react";
import { View, Text, Image, TouchableOpacity } from "react-native";
import Background from "../components/ui/Background";
import { useNavigateTo } from "../hooks/useNavigateTo";
import Icon from "react-native-vector-icons/MaterialCommunityIcons";
import logo from "../assets/images/logo.png";

export default function NotFoundScreen() {
  const goTo = useNavigateTo();

  return (
    <Background>
      <View className="flex justify-between items-center py-10 gap-20">
        <Image source={logo} className="w-70 h-60" resizeMode="contain" />

        <View className="items-center w-full">
          <View className="w-4/5 h-px bg-white mb-8" />
          <Text className="text-white font-bold text-9xl">404</Text>
          <Text className="text-white text-3xl m-6 text-center">
            Ops.. Essa Página Não Existe
          </Text>
          {/* TODO vou precisar colocar verificação aqui depois, pq no momento a página leve sempre pro login*/}
          <TouchableOpacity onPress={() => goTo("/login")}>
            <Text className="underline decoration-solid text-white text-2xl">
              Retorne a <Icon name="home" size={20} color="white" /> Clicando
              Aqui!
            </Text>
          </TouchableOpacity>
        </View>
        <View />
      </View>
    </Background>
  );
}
