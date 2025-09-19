import { SafeAreaProvider } from "react-native-safe-area-context";
import { View } from "react-native";
import { NativeRouter, Routes, Route } from "react-router-native";
import HomeScreen from "./src/screens/HomeScreen";
import SobreScreen from "./src/screens/SobreScreen";
import "./global.css";

export default function App() {
  return (
    <SafeAreaProvider>
      <View className="flex-1">
        <NativeRouter>
          <Routes>
            <Route path="/" element={<HomeScreen />} />
            <Route path="/sobre" element={<SobreScreen />} />
          </Routes>
        </NativeRouter>
      </View>
    </SafeAreaProvider>
  );
}
