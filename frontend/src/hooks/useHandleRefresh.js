import { useState } from "react";

export function useHandleRefresh() {
  const [listKey, setListKey] = useState(0);

  const handleRefresh = () => {
    setListKey((prevKey) => prevKey + 1);
  };

  return {
    listKey,
    handleRefresh,
  };
}
