import { useState, useCallback, useRef } from "react";

const initialData = {
  data: [],
  refreshing: false,
  loadingMore: false,
  currentPage: 1,
  totalPages: 1,
  totalResult: 0,
};

export function usePagination(callbackFetch) {
  const [listItems, setListItems] = useState(initialData.data);
  const [currentPage, setCurrentPage] = useState(initialData.currentPage);
  const [totalPages, setTotalPages] = useState(initialData.totalPages);
  const [totalResult, setTotalResult] = useState(initialData.totalResult);
  const [isLoadingMore, setIsLoadingMore] = useState(initialData.loadingMore);
  const [isRefreshing, setRefreshing] = useState(initialData.refreshing);

  const [error, setError] = useState(null);

  const isFetchingRef = useRef(false);
  const flatListRef = useRef(null);
  const allItemsLoaded = currentPage >= totalPages && listItems.length > 0;
  const itemsPerPage = 20;

  const initialLoad = useCallback(async () => {
    setError(null);
    try {
      const result = await callbackFetch(itemsPerPage, 1);
      if (result?.data) {
        setListItems(result.data);
        setCurrentPage(result.current_page);
        setTotalPages(result.total_pages);
        setTotalResult(result.total_items);
      }
    } catch (err) {
      console.error("Erro ao buscar dados iniciais:", err);
      setError("Não foi possível carregar os itens. Verifique sua conexão.");
    }
  }, [callbackFetch]);

  const loadMore = async () => {
    if (isLoadingMore || currentPage >= totalPages) return;

    try {
      isFetchingRef.current = true;
      setIsLoadingMore(true);
      setError(null);

      const result = await callbackFetch(itemsPerPage, currentPage + 1);
      if (result.data && result.data.length > 0) {
        setListItems((prevItems) => [...prevItems, ...result.data]);
        setCurrentPage(result.current_page);
      }
    } catch (err) {
      console.error("Erro ao buscar a próxima página:", err);
      setError("Não foi possível carregar mais itens.");
    } finally {
      isFetchingRef.current = false;
      setIsLoadingMore(false);
    }
  };

  const handleRefresh = useCallback(async () => {
    setRefreshing(true);
    await initialLoad();
    setRefreshing(false);
  }, [initialLoad]);

  const clearError = () => {
    setError(null);
  };

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
    error,
    initialLoad,
    loadMore,
    handleRefresh,
    scrollToTop,
    clearError,
  };
}
