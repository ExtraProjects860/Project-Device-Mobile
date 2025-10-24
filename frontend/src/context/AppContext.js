import React, {
  createContext,
  useContext,
  useEffect,
  useState,
  useCallback,
} from "react";
import { useColorScheme } from "nativewind";
import { Storage } from "../lib/storage.js";
import { setupAxiosInterceptors } from "../lib/axios.js";
import { useError } from "./ErrorContext.js";

const AppContext = createContext();

export function AppProvider({ children }) {
  const { colorScheme, setColorScheme } = useColorScheme();
  const [accessToken, setAccessToken] = useState(null);
  const [userData, setUserData] = useState({});
  const [isLoading, setIsLoading] = useState(true);
  const [_, setTheme] = useState("");

  const { showErrorModal } = useError();

  const isThemeDark = colorScheme === "dark";

  useEffect(() => {
    const loadingStorageData = async () => {
      const storedToken = await Storage.getItem("token");
      const storedUserData = await Storage.getItem("user");
      const storedTheme = await Storage.getItem("theme");

      if (storedToken) {
        setAccessToken(storedToken);
      }
      if (storedUserData) {
        setUserData(storedUserData);
      }
      if (storedTheme) {
        setTheme(storedTheme);
        setColorScheme(storedTheme);
      }

      setIsLoading(false);
    };

    loadingStorageData();
  }, []);

  const updateToken = async (newToken) => {
    setAccessToken(newToken);
    await Storage.setItem("token", newToken);
  };

  const updateUser = async (newUser) => {
    setUserData(newUser);
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

  const logout = useCallback(async () => {
    showErrorModal(
      "Sua sessão expirou. Por favor, faça login novamente.",
      null,
    );

    setAccessToken(null);
    setUserData(null);
    await Storage.removeItem("token");
    await Storage.removeItem("user");
    
  }, [showErrorModal]);

  const manuallyLogout = async () => {
    setAccessToken(null);
    setUserData(null);
    await Storage.removeItem("token");
    await Storage.removeItem("user");
  };

  useEffect(() => {
    setupAxiosInterceptors(logout);
  }, [logout]);

  const contexValues = {
    accessToken,
    userData,
    theme: colorScheme,
    isThemeDark,
    isLoading,
    updateToken,
    updateUser,
    updateTheme,
    toggleTheme,
    logout,
    manuallyLogout,
  };

  return (
    <AppContext.Provider value={contexValues}>{children}</AppContext.Provider>
  );
}

export const useAppContext = () => useContext(AppContext);
