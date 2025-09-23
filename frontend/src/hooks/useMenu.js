import { useState } from "react";

export function useMenu(initialState = false) {
  const [isVisible, setIsVisible] = useState(initialState);

  const openMenu = () => setIsVisible(true);
  const closeMenu = () => setIsVisible(false);

  return {
    isVisible,
    openMenu,
    closeMenu
  };
}
