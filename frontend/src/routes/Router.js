import React, { useEffect } from "react";
import { NativeRouter, Routes, Route, useLocation } from "react-router-native";
import HomeScreen from "../screens/HomeScreen";
import ProductsScreen from "../screens/admin/ProductsScreen";
import NotFoundScreen from "../screens/404";
import UsersScreen from "../screens/admin/UsersScreen";
import WishListScreen from "../screens/WishListScreen";
import NewPasswordScreen from "../screens/autentication/NewPasswordScreen";
import ForgotPasswordScreen from "../screens/autentication/ForgotPasswordScreen";
import LoginScreen from "../screens/autentication/LoginScreen";
import Loading from "../components/ui/Loading";
import { useAppContext } from "../context/AppContext";
import { useNavigateTo } from "../hooks/useNavigateTo";
import ProtectedRoutes from "./ProtectedRoutes";
import OnlineRoutes from "./OnlineRoutes";

export function AppRoutes() {
  const { accessToken, checkInternetConection } = useAppContext();
  return (
    <Routes>
      <Route path="/" element={<LoginScreen />} />
      <Route path="login" element={<LoginScreen />} />
      <Route path="home" element={<HomeScreen />} />

      <Route
        element={
          <OnlineRoutes
            accessToken={accessToken}
            checkInternetConection={checkInternetConection}
          />
        }
      >
        <Route path="new-password" element={<NewPasswordScreen />} />
        <Route path="forgot-password" element={<ForgotPasswordScreen />} />

        <Route element={<ProtectedRoutes accessToken={accessToken} />}>
          <Route path="products" element={<ProductsScreen />} />
          <Route path="users" element={<UsersScreen />} />
          <Route path="wishlist" element={<WishListScreen />} />
        </Route>
      </Route>

      <Route path="*" element={<NotFoundScreen />}></Route>
    </Routes>
  );
}

export default function Router() {
  const { isLoading } = useAppContext();

  if (isLoading) {
    return <Loading />;
  }

  return (
    <NativeRouter
      future={{ v7_startTransition: true, v7_relativeSplatPath: true }}
    >
      <AppRoutes />
    </NativeRouter>
  );
}
