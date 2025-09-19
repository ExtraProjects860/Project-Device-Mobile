import { Text, TouchableOpacity } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { useState } from "react";
import { useNavigate } from "react-router-native";
import Loading from "../components/Loading";

export default function HomeScreen() {
  const navigate = useNavigate();
  const [isLoading, setIsLoading] = useState(false);

  const handleNavigateSobre = () => {
    setIsLoading(true);

    setTimeout(() => {
      navigate("/sobre");
    }, 2000);
  };

  if (isLoading) return <Loading/>;

  return (
    <SafeAreaView className="flex-1 items-center justify-center bg-white">
      <Text className="text-xl font-bold text-blue-500">
        Welcome to Nativewind!
      </Text>
      <TouchableOpacity onPress={handleNavigateSobre} className="mt-4">
        <Text className="text-lg text-red-500">Ir para a tela Sobre</Text>
      </TouchableOpacity>
    </SafeAreaView>
  );
}
