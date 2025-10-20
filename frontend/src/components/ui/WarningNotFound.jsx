import { View, Text } from "react-native";

/**
 * Componente responsável por exibir menssagem de nenhum item encontrado
 *
 * Recebe 1 atributo.
 * O atributo chamado message responsável por receber a menssagem a ser exibida, já possui uma menssagem padrão.
 */
export default function WarningNotFound({
  message = "Nenhum item encontrado",
}) {
  return (
    <View className="flex-1 items-center justify-center mt-20">
      <Text className="text-white">{message}</Text>
    </View>
  );
}
