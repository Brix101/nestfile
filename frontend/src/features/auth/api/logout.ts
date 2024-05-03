import api from "@/lib/api";
import { UserResource } from "../types";
import { MutationConfig } from "@/lib/react-query";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { QUERY_KEYS } from "@/constant/query-key";

export async function logoutUser() {
  return api.post<UserResource>("/auth/logout");
}

export function useLogout(options?: MutationConfig<typeof logoutUser>) {
  const queryClient = useQueryClient();

  return useMutation({
    onSuccess: (response) => {
      queryClient.setQueryData([QUERY_KEYS.AUTH_USER], response.data);
    },
    ...options,
    mutationFn: logoutUser,
  });
}
