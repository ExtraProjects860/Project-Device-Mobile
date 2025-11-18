import { useState } from "react";

export function useMenu() {
  const [isVisible, setIsVisible] = useState(false);

  const openMenu = () => setIsVisible(true);
  const closeMenu = () => setIsVisible(false);

  return {
    isVisible,
    openMenu,
    closeMenu,
  };
}
