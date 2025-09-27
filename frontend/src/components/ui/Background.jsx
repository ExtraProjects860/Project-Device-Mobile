import React from "react";
import { View, Image } from "react-native";
import footer from "../../assets/images/footer.png";

export default function Background({ children, className = "" }) {
  return (
    <View className={`flex-1 bg-light-primary dark:bg-dark-primary -z-10 ${className}`}>
      {children}
      <View className="absolute bottom-0 w-full h-[400px] -z-10">
        <Image source={footer} className="w-full h-full" resizeMode="stretch" />
      </View>
    </View>
  );
}