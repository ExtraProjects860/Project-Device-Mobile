import { useState, useCallback, useRef } from "react";

const initialData = {
  data: [],
  refreshing: false,
  totalResult: 0,
  loadingMore: false,
  currentPage: 1,
  totalPages: 1,
};

export function usePagination(callbackFetch) {
  const [listItems, setListItems] = useState(initialData.data);
  const [currentPage, setCurrentPage] = useState(initialData.currentPage);
  const [totalPages, setTotalPages] = useState(initialData.totalPages);
  const [isLoadingMore, setIsLoadingMore] = useState(initialData.loadingMore);
  const [isRefreshing, setRefreshing] = useState(initialData.refreshing);

  const isFetchingRef = useRef(false);
  const flatListRef = useRef(null);
  const allItemsLoaded = currentPage >= totalPages && listItems.length > 0;
  const itemsPerPage = 20;

  const initialLoad = useCallback(async () => {
    try {
      const result = await callbackFetch(itemsPerPage, 1);
      if (result?.data) {
        setListItems(result.data);
        setCurrentPage(result.current_page);
        setTotalPages(result.total_pages);
      }
    } catch (error) {
      // TODO colocar erro na tela como modal para usuário, ou seja, fazer um modal padrão para erros
      console.error("Erro ao buscar dados iniciais:", error);
    }
  }, [callbackFetch]);

  const loadMore = async () => {
    if (isLoadingMore || currentPage >= totalPages) {
      return;
    }

    try {
      isFetchingRef.current = true;
      setIsLoadingMore(true);

      const result = await callbackFetch(itemsPerPage, currentPage + 1);
      if (result.data && result.data.length > 0) {
        setListItems((prevItems) => {
          return [...prevItems, ...result.data];
        });
        setCurrentPage(result.current_page);
      }
    } catch (error) {
      // TODO colocar erro na tela como modal para usuário, ou seja, fazer um modal padrão para erros
      console.error("Erro ao buscar a próxima página:", error);
    } finally {
      isFetchingRef.current = false;
      setIsLoadingMore(false);
    }
  };

  const handleRefresh = useCallback(async () => {
    setRefreshing(true);
    try {
      await initialLoad();
    } catch (error) {
      // TODO colocar erro na tela como modal para usuário, ou seja, fazer um modal padrão para erros
      console.error("Erro no refresh:", error);
    } finally {
      setRefreshing(false);
    }
  }, [initialLoad]);

  const scrollToTop = () => {
    flatListRef.current?.scrollToOffset({ offset: 0, animated: true });
  };

  return {
    listItems,
    currentPage,
    totalPages,
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
