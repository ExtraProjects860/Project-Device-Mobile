import React from "react";
import Header from "../components/Header.jsx";
import Menu from "../components/Menu";
import { useMenu } from "../hooks/useMenu.js";

export function NavBar() {
  const { isVisible, openMenu, closeMenu } = useMenu();

  return (
    <>
      <Header onMenuPress={openMenu} />
      <Menu visible={isVisible} closeMenu={closeMenu} />
    </>
  );
}
