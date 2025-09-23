import { SafeAreaProvider } from "react-native-safe-area-context";
import { View } from "react-native";
import { NativeRouter, Routes, Route } from "react-router-native";
import HomeScreen from "./src/screens/HomeScreen";
import ProductsScreen from "./src/screens/ProductsScreen";
import NotFoundScreen from "./src/screens/404";
import UsersScreen from "./src/screens/UsersScreen";
import WishListScreen from "./src/screens/WishListScreen";
import NoticesScreen from "./src/screens/NoticesScreeen";
import NewPasswordScreen from "./src/screens/NewPasswordScreen";
import ForgotPasswordScreen from "./src/screens/ForgotPasswordScreen";
import "./global.css";

export default function App() {
  return (
    <SafeAreaProvider>
      <View className="flex-1">
        <NativeRouter
          future={{ v7_startTransition: true, v7_relativeSplatPath: true }}
        >
          <Routes>
            <Route path="/" element={<HomeScreen />} />
            <Route path="products" element={<ProductsScreen />} />
            <Route path="users" element={<UsersScreen />} />
            <Route path="wishlist" element={<WishListScreen />} />
            <Route path="notices" element={<NoticesScreen />} />
            <Route path="new-password" element={<NewPasswordScreen />} />
            <Route path="forgot-password" element={<ForgotPasswordScreen />} />
            <Route />
            <Route path="*" element={<NotFoundScreen />}></Route>
          </Routes>
        </NativeRouter>
      </View>
    </SafeAreaProvider>
  );
}
