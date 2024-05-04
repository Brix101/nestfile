import React from "react";
import { Outlet } from "react-router-dom";

import { Loader } from "@/components/Loader";
import { MainLayout } from "@/components/layout/MainLayout";
import { lazyImport } from "@/utils/lazyImport";

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

export const protectedRoutes = [
  {
    path: "/",
    element: <App />,
    children: [{ index: true, element: <Files /> }],
  },
];
