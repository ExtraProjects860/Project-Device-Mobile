import "./global.css";
import { View } from "react-native";
import { ErrorProvider } from "./src/context/ErrorContext.js";
import { AppProvider, useAppContext } from "./src/context/AppContext.js";
import Router from "./src/routes/Router.js";

function AppLayout() {
  const { theme } = useAppContext();

  return (
    <View className={`flex-1 ${theme}`}>
      <Router />
    </View>
  );
}

export default function App() {
  return (
    <AppProvider>
      <ErrorProvider>
        <AppLayout/>
      </ErrorProvider>
    </AppProvider>
  );
}
