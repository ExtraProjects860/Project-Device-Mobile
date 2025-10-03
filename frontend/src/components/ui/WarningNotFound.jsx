import { View, Text } from "react-native";

export default function WarningNotFound({ message }) {
  return (
    <View className="flex-1 items-center justify-center mt-20">
      <Text className="text-white">{message}</Text>
    </View>
  );
}
