import { RouterProvider, createBrowserRouter } from "react-router-dom";
import React from "react";

import { useUser } from "@/hooks/useUser";
import { Loader } from "@/components/Loader";

import { protectedRoutes } from "./protected";
import { publicRoutes } from "./public";

export function AppRoutes() {
  const user = useUser();

  const routes = user.isSignedIn ? protectedRoutes : publicRoutes;

  const router = createBrowserRouter([...routes], {});

  return (
    <React.Suspense
      fallback={
        <div className="fixed top-0 left-0 z-50 h-screen w-full flex justify-center items-center">
          <Loader />
        </div>
      }
    >
      <RouterProvider router={router} />
    </React.Suspense>
  );
}
