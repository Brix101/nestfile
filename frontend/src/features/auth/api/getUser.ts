import { useQuery } from "@tanstack/react-query";

import { QUERY_KEYS } from "@/constant/query-key";
import api from "@/lib/api";
import { QueryConfig } from "@/lib/react-query";
import { Resources } from "../types";

export async function fetchUser() {
  const res = await api.get<Resources>("/auth/me");
  return res.data;
}

export const getUserQuery = () => ({
  queryKey: [QUERY_KEYS.AUTH_USER],
  queryFn: fetchUser,
});

type UseGetUserOptions = QueryConfig<typeof fetchUser>;

export function useGetUser(options?: UseGetUserOptions) {
  return useQuery({
    ...options,
    ...getUserQuery(),
  });
}
