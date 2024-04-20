import LoginPage from "@/pages/Login";
import { RouterProvider, createBrowserRouter } from "react-router-dom";

function BrowserProvider() {
  const routes = createBrowserRouter([{ index: true, element: <LoginPage /> }]);

  return <RouterProvider router={routes} />;
}

export default BrowserProvider;
