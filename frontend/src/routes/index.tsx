import { useRoutes } from "react-router-dom";

import { useUser } from "@/hooks/useUser";
import { protectedRoutes } from "./protected";
import { publicRoutes } from "./public";

export function AppRoutes() {
  const user = useUser();

  const routes = user.isSignedIn ? protectedRoutes : publicRoutes;

  const elements = useRoutes([...routes]);

  return <>{elements}</>;
}
