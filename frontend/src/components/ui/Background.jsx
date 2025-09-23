import React from "react";
import { View, Image } from "react-native";
import footer from "../../assets/images/footer.png";

export default function Background({children, className = ""}) {
  return (
    <View className="flex-1 bg-teal -z-10">
      {children}
      <View className="absolute bottom-0 w-full h-[400px] -z-10">
        <Image
          source={footer}
          className="w-full h-full"
          resizeMode="stretch"
        />
      </View>
    </View>
  );
}