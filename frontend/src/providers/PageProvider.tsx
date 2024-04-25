import ProtectedRoute from "@/components/ProtectedRoute";
import LoginPage from "@/pages/Login";
import FileListingPage from "@/pages/files/FileListing";
import { RouterProvider, createBrowserRouter } from "react-router-dom";

function PageProvider() {
  const routes = createBrowserRouter([
    { index: true, element: <LoginPage /> },
    {
      path: "/files",
      element: <ProtectedRoute />,
      children: [{ index: true, element: <FileListingPage /> }],
    },
  ]);

  return <RouterProvider router={routes} />;
}

export default PageProvider;
