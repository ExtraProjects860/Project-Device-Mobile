import { useNavigate } from "react-router-native";

export function useNavigateTo() {
  const navigate = useNavigate();

  const goTo = (screen, params = {}) => {
    navigate(screen, { state: params });
  };

  return goTo;
}
