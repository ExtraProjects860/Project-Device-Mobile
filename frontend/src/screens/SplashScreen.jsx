import React, { useEffect } from "react";
import { useAppContext } from "../context/AppContext.js";
import { useNavigateTo } from "../hooks/useNavigateTo.js";
import Loading from "../components/ui/Loading.jsx";

export default function SplashScreen() {
  const { accessToken, isLoading } = useAppContext();
  const goTo = useNavigateTo();

  useEffect(() => {
    if (!isLoading) {
      if (!accessToken) {
        goTo("/login");
        return;
      }

      goTo("/home");
    }
  }, [isLoading, accessToken, goTo]);

  return <Loading/>;
}
