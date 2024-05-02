import { QUERY_KEYS } from "@/constant/query-key";
import api from "@/lib/api";
import { LoginInput, Resources } from "@/types/auth";
import { UserResource } from "@/types/user";

export async function fetchUser() {
  const res = await api.get<Resources>("/auth/me");
  return res.data;
}

export async function loginUser(data: LoginInput) {
  return api.post<UserResource>("/auth/login", JSON.stringify(data));
}

export async function logoutUser() {
  return api.post<UserResource>("/auth/logout");
}

export const getUserQuery = () => ({
  queryKey: [QUERY_KEYS.auth_user],
  queryFn: fetchUser,
});
