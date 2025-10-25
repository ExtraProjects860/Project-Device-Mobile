import React from "react";
import { Navigate, Outlet } from "react-router-native";

/**
 *
 * @param  {object} props
 * @param  {} props.accessToken
 * @param  {} props.checkInternetConection
 */
export default function OnlineRoutes({ accessToken, checkInternetConection }) {
  if (checkInternetConection) {
    return <Outlet />;
  }

  const redirectPath = accessToken ? "/home" : "/login";
  return <Navigate to={redirectPath} replace />;
}
