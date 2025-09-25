import React, { useState, useEffect } from "react";
import { useLoading } from "../hooks/useLoading";
import Loading from "../components/Loading";

export default function PageLoader({ fetchData, children }) {
  const { isLoading, withLoading } = useLoading();
  const [data, setData] = useState(null);

  useEffect(() => {
    withLoading(async () => {
      await fetchData(setData);
    });
  }, [fetchData, withLoading]);

  if (isLoading) return <Loading />;

  return <>{children(data)}</>
}
