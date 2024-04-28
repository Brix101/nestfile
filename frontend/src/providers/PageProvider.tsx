import { RouterProvider, createBrowserRouter } from "react-router-dom";
import { Suspense, lazy } from "react";

import ProtectedRoute from "@/components/ProtectedRoute";

const LoginPage = lazy(() => import("@/pages/Login"));
const FileListingPage = lazy(() => import("@/pages/files/FileListing"));

function PageProvider() {
  const routes = createBrowserRouter(
    [
      { index: true, element: <LoginPage /> },
      {
        path: "/files",
        element: <ProtectedRoute />,
        children: [{ index: true, element: <FileListingPage /> }],
      },
    ],
    {},
  );

  return (
    <Suspense fallback={<Loading />}>
      <RouterProvider router={routes} />;
    </Suspense>
  );
}

function Loading() {
  return <>Loading</>;
}

export default PageProvider;
