import React, { createContext, useState, useContext } from "react";

const ThemeContext = createContext();

export function ThemeProvider({ children }) {
  const [isThemeDark, setIsThemeDark] = useState(false);

  const toggleTheme = () => {
    setIsThemeDark((previousState) => !previousState);
  };

  const value = {
    isThemeDark,
    toggleTheme,
    themeClass: isThemeDark ? "dark" : "",
  };

  return (
    <ThemeContext.Provider value={value}>{children}</ThemeContext.Provider>
  );
}

export function useTheme() {
  return useContext(ThemeContext);
}
