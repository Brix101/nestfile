import {
  Navigate,
  RouterProvider,
  createBrowserRouter,
} from "react-router-dom";
import { Suspense, lazy } from "react";

import ProtectedRoute from "@/components/ProtectedRoute";
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
        element: <ProtectedRoute />,
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
  return <>Loading...</>;
}

export default PageProvider;
