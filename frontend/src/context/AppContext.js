import React, { createContext, useContext, useEffect, useState } from "react";
import { useColorScheme } from "nativewind";
import { Storage } from "../lib/Storage.js";

const AppContext = createContext();

export function AppProvider({ children }) {
  const { colorScheme, setColorScheme } = useColorScheme();
  const [accessToken, setAcessToken] = useState(null);
  const [userData, setUserData] = useState({});
  const [theme, setTheme] = useState("");

  const isThemeDark = colorScheme === "dark";

  useEffect(() => {
    const loadingStorageData = async () => {
      const storedToken = await Storage.getItem("token");
      const storedUserData = await Storage.getItem("user");
      const storedTheme = await Storage.getItem("theme");

      if (storedToken) {
        setAcessToken(storedToken);
      }
      if (storedUserData) {
        setUserData(storedUserData);
      }
      if (storedTheme) {
        setTheme(storedTheme);
        setColorScheme(storedTheme);
      }
    };

    loadingStorageData();
  }, []);

  const updateToken = async (newToken) => {
    setToken(newToken);
    await Storage.setItem("token", newToken);
  };

  const updateUser = async (newUser) => {
    setUser(newUser);
    await Storage.setItem("user", newUser);
  };

  const updateTheme = async (newTheme) => {
    setTheme(newTheme);
    setColorScheme(newTheme);
    await Storage.setItem("theme", newTheme);
  };

  const toggleTheme = async () => {
    const newTheme = colorScheme === "dark" ? "light" : "dark";
    await updateTheme(newTheme);
  };

  const logout = async () => {
    setAcessToken(null);
    setUserData(null);
    await Storage.removeItem("token");
    await Storage.removeItem("user");
  };

  const contexValues = {
    accessToken,
    userData,
    theme: colorScheme,
    isThemeDark,
    updateToken,
    updateUser,
    updateTheme,
    toggleTheme,
    logout,
  };

  return (
    <AppContext.Provider value={contexValues}>{children}</AppContext.Provider>
  );
}

export const useAppContext = () => useContext(AppContext);
