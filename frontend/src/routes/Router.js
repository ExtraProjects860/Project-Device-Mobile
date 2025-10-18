import { NativeRouter, Routes, Route } from "react-router-native";
import HomeScreen from "../screens/HomeScreen";
import ProductsScreen from "../screens/ProductsScreen";
import NotFoundScreen from "../screens/404";
import UsersScreen from "../screens/UsersScreen";
import WishListScreen from "../screens/WishListScreen";
import NoticesScreen from "../screens/NoticesScreeen";
import NewPasswordScreen from "../screens/NewPasswordScreen";
import ForgotPasswordScreen from "../screens/ForgotPasswordScreen";
import LoginScreen from "../screens/LoginScreen";

export default function Router() {
  return (
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
        <Route path="*" element={<NotFoundScreen />}></Route>
      </Routes>
    </NativeRouter>
  );
}
