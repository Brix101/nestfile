import { Navigate, Outlet } from "react-router-dom";

import { useUser } from "@/hooks/useUser";

function ProtectedRoute() {
  const { user, isLoaded } = useUser();

  if (isLoaded && !user) {
    return <Navigate to="/" />;
  }

  return <Outlet />;
}

export default ProtectedRoute;
