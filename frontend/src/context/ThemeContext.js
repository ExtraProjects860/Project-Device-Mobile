import React, { createContext, useState, useContext, useEffect } from "react";
import { useColorScheme } from "nativewind";

const ThemeContext = createContext();

export function ThemeProvider({ children }) {
  const { colorScheme, setColorScheme } = useColorScheme();
  const [isThemeDark, setIsThemeDark] = useState(colorScheme === "dark");

  useEffect(() => {
    setIsThemeDark(colorScheme === "dark");
  }, [colorScheme]);

  const toggleTheme = () => {
    const newScheme = colorScheme === "dark" ? "light" : "dark";
    setColorScheme(newScheme);
  };

  const value = {
    isThemeDark,
    toggleTheme,
  };

  return (
    <ThemeContext.Provider value={value}>{children}</ThemeContext.Provider>
  );
}

export function useTheme() {
  return useContext(ThemeContext);
}
