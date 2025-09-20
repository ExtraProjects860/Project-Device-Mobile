import { View, Image } from "react-native";
const footer = require("../assets/images/footer.png");

export default function background({children, className = ""}) {
  return (
    <View className="flex-1 bg-teal -z-10">
      {children}
      <View className="absolute bottom-0 w-full h-80 -z-10">
        <Image
          source={footer}
          className="w-full h-full"
          resizeMode="stretch"
        />
      </View>
    </View>
  );
}