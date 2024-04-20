import { useUser } from "@/hooks/useUser";
import React from "react";
import { Navigate } from "react-router-dom";

interface ProtectedRouteProps extends React.PropsWithChildren {}

function ProtectedRoute({ children }: ProtectedRouteProps) {
  const { user, isLoaded } = useUser();

  if (isLoaded && !user) {
    return <Navigate to="/" />;
  }

  return children;
}

export default ProtectedRoute;
