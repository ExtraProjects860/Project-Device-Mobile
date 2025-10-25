import React, { Children } from "react";
import { Navigate, Outlet } from "react-router-native";

/**
 *
 * @param  {object} props
 * @param  {} props.accessToken
 * @param  {} props.redirectPath
 * @param  {} props.childern
 */
export default function ProtectedRoutes({
  accessToken,
  redirectPath = "/",
  children,
}) {
  if (!accessToken) {
    return <Navigate to={redirectPath} replace />;
  }
  return children ? children : <Outlet />;
}
