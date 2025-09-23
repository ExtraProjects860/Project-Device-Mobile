import React from "react";
import { View, Image, StatusBar } from "react-native";
import footer from "../../assets/images/footer.png";

export default function Background({ children, className = "" }) {
  return (
    <View className="flex-1 bg-teal -z-10">
      {children}
      <StatusBar barStyle={"light-content"} translucent={true} backgroundColor={"#E91D62"} />
      <View className="absolute bottom-0 w-full h-[400px] -z-10">
        <Image source={footer} className="w-full h-full" resizeMode="stretch" />
      </View>
    </View>
  );
}
