import { useNavigate } from "react-router-native";

export function useNavigateTo() {
  const navigate = useNavigate();

  function goTo(screen, params = {}) {
    navigate(screen, { state: params });
  }

  return goTo;
}
