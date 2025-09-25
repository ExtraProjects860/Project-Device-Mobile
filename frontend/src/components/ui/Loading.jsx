import React, { useEffect, useState } from "react";
import { ActivityIndicator, View, Text } from "react-native";

export default function Loading() {
  const [dots, setDots] = useState("");

  useEffect(() => {
    let count = 0;
    const interval = setInterval(() => {
      count = (count + 1) % 4;
      setDots(".".repeat(count));
    }, 300);

    return () => clearInterval(interval);
  }, []);

  return (
    <View className="flex-1 bg-light-primary dark:bg-dark-primary items-center justify-center">
      <ActivityIndicator size="large" color="#E91D62" />
      <Text className="text-light-text-inverted dark:text-dark-text-primary text-xl">Carregando{dots}</Text>
    </View>
  );
}
