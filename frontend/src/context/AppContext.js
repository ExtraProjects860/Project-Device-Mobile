import React, { createContext, useContext, useEffect, useState } from "react";
import { Storage } from "../lib/Storage.js";

const AppContext = createContext();

// TODO vou ter que jogar isso lá pro App.jsx

export function AppProvider({ children }) {
  const [accessToken, setAcessToken] = useState(null);
  const [userData, setUserData] = useState({});
  const [theme, setTheme] = useState("");

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
    await Storage.setItem("theme", newTheme);
  };

  const logout = async () => {
    setAcessToken(null);
    setUserData(null);
    await Storage.removeItem('token');
    await Storage.removeItem('user');
  };

  const contexValues = {
            accessToken,
        userData,
        theme,
        loading,
        updateToken,
        updateUser,
        updateTheme,
        logout,
  }

  return (
    <AppContext.Provider
      value={contexValues}
    >
      {children}
    </AppContext.Provider>
  )
}

export const useAppContext = () => useContext(AppContext);
