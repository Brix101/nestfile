import { lazyImport } from "@/utils/lazyImport";

const { Login } = lazyImport(() => import("@/features/auth"), "Login");

export const publicRoutes = [
  {
    path: "/",
    element: <Login />,
  },
];
