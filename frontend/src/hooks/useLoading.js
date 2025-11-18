import { useCallback, useState } from "react";
import { useError } from "../context/ErrorContext";

export function useLoading(initialState = false) {
  const [isLoading, setIsLoading] = useState(initialState);
  const { showErrorModal } = useError();

  const withLoading = useCallback(
    async (asyncFunction) => {
      try {
        setIsLoading(true);
        await asyncFunction();
      } catch (error) {
        console.error("Erro capturado pelo withLoading:", error);
        showErrorModal("Erro ao carregar componente.");
      } finally {
        setIsLoading(false);
      }
    },
    [showErrorModal],
  );

  return { isLoading, withLoading };
}
