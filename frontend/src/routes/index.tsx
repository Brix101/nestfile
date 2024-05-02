import {
  Navigate,
  RouterProvider,
  createBrowserRouter,
} from "react-router-dom";

import ProtectedLayout from "@/components/layout/ProtectedLayout";

import { useUser } from "@/hooks/useUser";

import { lazyImport } from "@/utils/lazyImport";
import { Loader } from "@/components/Loader";
import React from "react";

const { Login } = lazyImport(() => import("@/pages/Login"), "Login");
const { FileListing } = lazyImport(
  () => import("@/pages/files/FileListing"),
  "FileListing",
);

export function AppRoutes() {
  const { isSignedIn } = useUser();

  const routes = createBrowserRouter(
    [
      {
        index: true,
        element: isSignedIn ? <Navigate to="/files" /> : <Login />,
      },
      {
        path: "/files",
        element: <ProtectedLayout />,
        children: [{ index: true, element: <FileListing /> }],
      },
    ],
    {},
  );

  return (
    <React.Suspense
      fallback={
        <div className="fixed top-0 left-0 z-50 h-screen w-full flex justify-center items-center">
          <Loader />
        </div>
      }
    >
      <RouterProvider router={routes} />
    </React.Suspense>
  );
}
