import { lazyImport } from "@/utils/lazyImport";
import { RouteObject } from "react-router-dom";

const { Login } = lazyImport(() => import("@/features/auth"), "Login");

export const publicRoutes: RouteObject[] = [
  {
    path: "/",
    element: <Login />,
  },
];
