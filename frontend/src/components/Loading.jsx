import { ActivityIndicator, View } from "react-native";

export default function Loading() {
  return (
    <View className="flex-1 bg-[rgba(0, 0, 0, 0.5)] items-center justify-center">
      <ActivityIndicator size="large" color="#E91D62" />
    </View>
  );
}
