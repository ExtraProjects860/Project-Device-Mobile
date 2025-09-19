import { Text, TouchableOpacity } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { useNavigate } from "react-router-native";

export default function SobreScreen() {
  const navigate = useNavigate();

  const handleNavigateHome = () => {
    navigate("/");
  }

  return (
    <SafeAreaView className="flex-1 items-center justify-center bg-slate-100">
      <Text className="text-xl font-bold text-slate-800">
        Esta é a Tela Sobre
      </Text>
      <TouchableOpacity onPress={handleNavigateHome} className="mt-4">
        <Text className="mt-4 text-lg text-blue-500">Voltar para Home</Text>
      </TouchableOpacity>
    </SafeAreaView>
  );
}
