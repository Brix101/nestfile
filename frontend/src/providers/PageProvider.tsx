import { Suspense, lazy } from "react";
import {
  Navigate,
  RouterProvider,
  createBrowserRouter,
} from "react-router-dom";

import ResourcesLayout from "@/components/layout/ResourcesLayout";
import { useUser } from "@/hooks/useUser";

const LoginPage = lazy(() => import("@/pages/Login"));
const FileListingPage = lazy(() => import("@/pages/files/FileListing"));

function PageProvider() {
  const { isSignedIn } = useUser();

  const routes = createBrowserRouter(
    [
      {
        index: true,
        element: isSignedIn ? <Navigate to="/files" /> : <LoginPage />,
      },
      {
        path: "/files",
        element: <ResourcesLayout />,
        children: [{ index: true, element: <FileListingPage /> }],
      },
    ],
    {},
  );

  return (
    <Suspense fallback={<Loading />}>
      <RouterProvider router={routes} />
    </Suspense>
  );
}

function Loading() {
  return (
    <div className="fixed top-0 left-0 z-50 h-screen w-full flex justify-center items-center">
      <div id="loading" className="bg-background">
        <div className="spinner">
          <div className="bounce1 bg-primary"></div>
          <div className="bounce2 bg-primary"></div>
          <div className="bounce3 bg-primary"></div>
        </div>
      </div>
    </div>
  );
}

export default PageProvider;
