import api from "@/lib/api";
import { LoginDTO, UserResource } from "../types";
import { MutationConfig } from "@/lib/react-query";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { QUERY_KEYS } from "@/constant/query-key";

export async function loginUser(data: LoginDTO) {
  return api.post<UserResource>("/auth/login", JSON.stringify(data));
}

export function useUserLogin(options?: MutationConfig<typeof loginUser>) {
  const queryClient = useQueryClient();

  return useMutation({
    onSuccess: (response) => {
      queryClient.setQueryData([QUERY_KEYS.auth_user], response.data);
    },
    ...options,
    mutationFn: loginUser,
  });
}
