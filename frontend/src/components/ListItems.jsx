import React, { useImperativeHandle, forwardRef } from "react";
import {
  View,
  Text,
  FlatList,
  TouchableOpacity,
  RefreshControl,
} from "react-native";
import PageLoader from "../context/PageLoader.js";
import Loading from "./ui/Loading.jsx";
import WarningNotFound from "./ui/WarningNotFound.jsx";
import { usePagination } from "../hooks/usePagination.js";
import ModalErrors from "./modals/ModalErrors";

const ListItems = forwardRef(
  ({ callbackFetch, CardListRender, onEditItem }, ref) => {
    const {
      listItems,
      isLoadingMore,
      isRefreshing,
      flatListRef,
      totalResult,
      allItemsLoaded,
      error,
      initialLoad,
      loadMore,
      handleRefresh,
      scrollToTop,
      clearError,
    } = usePagination(callbackFetch);

    useImperativeHandle(ref, () => ({
      refresh: handleRefresh,
    }));

    const renderInFooter = () => {
      if (isLoadingMore) {
        return <Loading />;
      }
      if (allItemsLoaded) {
        return (
          <View className="w-full items-center">
            <TouchableOpacity
              onPress={scrollToTop}
              className="rounded-full m-2 py-2"
            >
              <Text className="text-white text-xl underline underline-offset-1 font-bold">
                Voltar ao Topo
              </Text>
            </TouchableOpacity>
          </View>
        );
      }
      return null;
    };

    return (
      <PageLoader fetchData={initialLoad}>
        <ModalErrors
          visible={!!error}
          message={error}
          onClose={clearError}
          onRetry={handleRefresh}
        />

        {isRefreshing && listItems.length > 0 ? (
          <Loading />
        ) : (
          <View className="flex-1">
            <View className="w-full items-center my-4">
              <Text className="text-white">
                Total de itens sendo exibidos: {listItems.length} de{" "}
                {totalResult}
              </Text>
            </View>

            <FlatList
              contentContainerStyle={{ paddingBottom: 40 }}
              ref={flatListRef}
              data={listItems}
              numColumns={2}
              keyExtractor={(item) => item.id.toString()}
              onEndReached={loadMore}
              onEndReachedThreshold={0.5}
              ListFooterComponent={renderInFooter}
              renderItem={({ item }) => <CardListRender item={item} />}
              refreshControl={
                <RefreshControl
                  refreshing={isRefreshing}
                  onRefresh={handleRefresh}
                />
              }
              ListEmptyComponent={
                !error && <WarningNotFound/>
              }
            />
          </View>
        )}
      </PageLoader>
    );
  }
);

export default ListItems;
