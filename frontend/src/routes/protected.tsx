import React from "react";
import { Outlet, RouteObject } from "react-router-dom";

import { Loader } from "@/components/Loader";
import { lazyImport } from "@/utils/lazyImport";
import { MainLayout } from "@/components/layout/MainLayout";

const { Files } = lazyImport(() => import("@/features/files"), "Files");

const App = () => {
  return (
    <MainLayout>
      <React.Suspense
        fallback={
          <div className="fixed top-0 left-0 z-50 h-screen w-full flex justify-center items-center">
            <Loader />
          </div>
        }
      >
        <Outlet />
      </React.Suspense>
    </MainLayout>
  );
};

export const protectedRoutes: RouteObject[] = [
  {
    path: "/",
    element: <App />,
    children: [{ index: true, element: <Files /> }],
  },
];
