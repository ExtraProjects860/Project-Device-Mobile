import { useState, useEffect, useCallback } from "react";
import { useAppContext } from "../context/AppContext";
import { useError } from "../context/ErrorContext";
import {
  getRolesRequest,
  getEnterprisesRequest,
} from "../lib/formOptionsRequests.js";

export function useFormOptions() {
  const { accessToken } = useAppContext();
  const { showErrorModal } = useError();

  const [roles, setRoles] = useState([]);
  const [enterprises, setEnterprises] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);

  const fetchOptions = useCallback(async () => {
    if (!accessToken) return;

    setIsLoading(true);
    setError(null);

    const allItemsLimit = 500;
    const firstPage = 1;

    try {
      const [rolesData, enterprisesData] = await Promise.all([
        getRolesRequest(accessToken, allItemsLimit, firstPage),
        getEnterprisesRequest(accessToken, allItemsLimit, firstPage),
      ]);

      setRoles(rolesData);
      setEnterprises(enterprisesData);
    } catch (err) {
      console.error("Erro ao buscar opções do formulário:", err);
      setError(err);
      showErrorModal(
        "Não foi possível carregar as opções de Funções e Empresas.",
      );
    } finally {
      setIsLoading(false);
    }
  }, [accessToken, showErrorModal]);

  useEffect(() => {
    fetchOptions();
  }, [fetchOptions]);

  return {
    roles,
    enterprises,
    isLoading,
    error,
    refetch: fetchOptions,
  };
}
