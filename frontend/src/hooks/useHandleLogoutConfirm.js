import { useAppContext } from "../context/AppContext";
import { useNavigateTo } from "./useNavigateTo";

export function useHandleLogoutConfirm () {
  const goTo = useNavigateTo();
  const { manuallyLogout } = useAppContext();

  const handleLogout = async () => {
    await manuallyLogout();
    goTo("/login");
  };
  
  return handleLogout;
}
 
