import { SafeAreaProvider } from "react-native-safe-area-context";
import { View } from "react-native";
import { NativeRouter, Routes, Route } from "react-router-native";
import { ThemeProvider, useTheme } from "./src/context/ThemeContext.js";
import HomeScreen from "./src/screens/HomeScreen";
import ProductsScreen from "./src/screens/ProductsScreen";
import NotFoundScreen from "./src/screens/404";
import UsersScreen from "./src/screens/UsersScreen";
import WishListScreen from "./src/screens/WishListScreen";
import NoticesScreen from "./src/screens/NoticesScreeen";
import NewPasswordScreen from "./src/screens/NewPasswordScreen";
import ForgotPasswordScreen from "./src/screens/ForgotPasswordScreen";
import LoginScreen from "./src/screens/LoginScreen.jsx";
import "./global.css";
import { ErrorProvider } from "./src/context/ErrorContext.js";

function AppContent() {
  const { themeClass } = useTheme();

  return (
    <ErrorProvider>
      <View className={`flex-1 ${themeClass}`}>
        <NativeRouter
          future={{ v7_startTransition: true, v7_relativeSplatPath: true }}
        >
          <Routes>
            <Route path="/" element={<LoginScreen />} />
            <Route path="home" element={<HomeScreen />} />
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
    </ErrorProvider>
  );
}

export default function App() {
  return (
    <SafeAreaProvider>
      <ThemeProvider>
        <AppContent />
      </ThemeProvider>
    </SafeAreaProvider>
  );
}
