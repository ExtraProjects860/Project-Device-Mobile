import React, { useEffect } from "react";
import { NativeRouter, Routes, Route, useLocation} from "react-router-native";
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
import { useNavigateTo } from "../hooks/useNavigateTo"

export default function Router() {
  function NavigationGuard({ children }) {
    const { accessToken, isLoading } = useAppContext();
    const goTo = useNavigateTo();
    const location = useLocation();

    useEffect(() => {
      if (isLoading) {
        return;
      }

      if (!accessToken) {
        if (
          location.pathname !== "/login" &&
          location.pathname !== "/forgot-password" &&
          location.pathname !== "/new-password"
        ) {
          goTo("/login");
        }
        return;
      }

      if (
        location.pathname === "/login" ||
        location.pathname === "/forgot-password" ||
        location.pathname === "/new-password" ||
        location.pathname === "/"
      ) {
        goTo("/home");
      }
    }, [isLoading, accessToken, goTo, location.pathname]);

    if (isLoading) {
      return <Loading />;
    }
    return <>{children}</>;
  }
  return (
    <NativeRouter
      future={{ v7_startTransition: true, v7_relativeSplatPath: true }}
    >
      <NavigationGuard>
        <Routes>
          <Route path="/" element={<Loading />} />

          <Route path="login" element={<LoginScreen />} />
          <Route path="home" element={<HomeScreen />} />
          <Route path="products" element={<ProductsScreen />} />
          <Route path="users" element={<UsersScreen />} />
          <Route path="wishlist" element={<WishListScreen />} />
          <Route path="new-password" element={<NewPasswordScreen />} />
          <Route path="forgot-password" element={<ForgotPasswordScreen />} />
          <Route path="*" element={<NotFoundScreen />}></Route>
        </Routes>
      </NavigationGuard>
    </NativeRouter>
  );
}
