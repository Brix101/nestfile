import { Navigate, Outlet } from "react-router-dom";

import { useUser } from "@/hooks/useUser";

import SiteHeader from "./SiteHeader";

function ProtectedLayout() {
  const { user, isLoaded } = useUser();

  if (isLoaded && !user) {
    return <Navigate to="/" />;
  }

  return (
    <>
      <SiteHeader />
      <Outlet />
    </>
  );
}

export default ProtectedLayout;
