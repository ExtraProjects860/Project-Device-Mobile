import { useCallback, useState } from "react";

export function useLoading(initialState = false) {
  const [isLoading, setIsLoading] = useState(initialState);

  const withLoading = useCallback(async (asyncFunction) => {
    try {
      setIsLoading(true);
      await asyncFunction();
    } finally {
      setIsLoading(false);
    }
  }, []);

  return { isLoading, withLoading };
}
