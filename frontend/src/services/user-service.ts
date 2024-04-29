import { QUERY_KEYS } from "@/constant/query-key";
import api from "@/lib/api";
import { Resources } from "@/types/auth";

export async function fetchUser() {
  const res = await api.get<Resources>("/auth/user");
  return res.data;
}

export const getUserQuery = () => ({
  queryKey: [QUERY_KEYS.auth_user],
  queryFn: fetchUser,
});
