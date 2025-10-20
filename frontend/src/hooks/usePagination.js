import { useState, useCallback, useRef } from "react";
import { useError } from "../context/ErrorContext";
import { useAppContext } from "../context/AppContext";

const initialData = {
  data: [],
  refreshing: false,
  loadingMore: false,
  currentPage: 1,
  totalPages: 1,
  totalResult: 0,
};

export function usePagination(callbackFetch) {
  const { accessToken } = useAppContext();

  const [listItems, setListItems] = useState(initialData.data);
  const [currentPage, setCurrentPage] = useState(initialData.currentPage);
  const [totalPages, setTotalPages] = useState(initialData.totalPages);
  const [totalResult, setTotalResult] = useState(initialData.totalResult);
  const [isLoadingMore, setIsLoadingMore] = useState(initialData.loadingMore);
  const [isRefreshing, setRefreshing] = useState(initialData.refreshing);

  const { showErrorModal } = useError();

  const isFetchingRef = useRef(false);
  const flatListRef = useRef(null);
  const allItemsLoaded = currentPage >= totalPages && listItems.length > 0;
  const itemsPerPage = 20;

  const initialLoad = useCallback(async () => {
    try {
      if (typeof callbackFetch === "function") {
        const result = await callbackFetch(itemsPerPage, 1, accessToken);
        if (result?.data) {
          setListItems(result.data);
          setCurrentPage(result.current_page);
          setTotalPages(result.total_pages);
          setTotalResult(result.total_items);
        }
      } else {
        console.error("callbackFetch is not a function");
      }
    } catch (err) {
      console.error("Erro ao buscar dados iniciais:", err);
      const message =
        err.message ||
        "Não foi possível carregar os itens. Verifique sua conexão.";
      showErrorModal(message, initialLoad);
    }
  }, [callbackFetch, showErrorModal, accessToken]);

  const loadMore = useCallback(async () => {
    if (isLoadingMore || currentPage >= totalPages) return;

    try {
      isFetchingRef.current = true;
      setIsLoadingMore(true);

      const result = await callbackFetch(
        itemsPerPage,
        currentPage + 1,
        accessToken,
      );
      if (result.data && result.data.length > 0) {
        setListItems((prevItems) => [...prevItems, ...result.data]);
        setCurrentPage(result.current_page);
      }
    } catch (err) {
      console.error("Erro ao buscar a próxima página:", err);
      showErrorModal("Não foi possível carregar mais itens.", loadMore);
    } finally {
      isFetchingRef.current = false;
      setIsLoadingMore(false);
    }
  }, [
    isLoadingMore,
    currentPage,
    totalPages,
    callbackFetch,
    showErrorModal,
    accessToken,
  ]);

  const handleRefresh = useCallback(async () => {
    setRefreshing(true);
    await initialLoad();
    setRefreshing(false);
  }, [initialLoad]);

  const scrollToTop = () => {
    flatListRef.current?.scrollToOffset({ offset: 0, animated: true });
  };

  return {
    listItems,
    currentPage,
    totalPages,
    totalResult,
    flatListRef,
    allItemsLoaded,
    isLoadingMore,
    isRefreshing,
    initialLoad,
    loadMore,
    handleRefresh,
    scrollToTop,
  };
}
