import React, { useEffect } from "react";
import { useLoading } from "../hooks/useLoading";
import Loading from "../components/ui/Loading";

export default function PageLoader({ fetchData, children }) {
  const { isLoading, withLoading } = useLoading();

  useEffect(() => {
    withLoading(fetchData);
  }, [fetchData, withLoading]);

  if (isLoading) return <Loading />;

  return <>{children}</>
}
