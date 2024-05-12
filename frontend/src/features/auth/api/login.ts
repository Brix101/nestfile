import api from "@/lib/api";
import { UserResource } from "../types";
import { MutationConfig } from "@/lib/react-query";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { QUERY_KEYS } from "@/constant/query-key";
import * as z from "zod";

export const loginSchema = z.object({
  username: z.string(),
  password: z
    .string()
    .min(8, {
      message: "Password must be at least 8 characters long",
    })
    .max(100, {
      message: "Password must be at most 100 characters long",
    }),
});

export type LoginDTO = z.infer<typeof loginSchema>;

export async function loginUser(data: LoginDTO) {
  return api.post<UserResource>("/auth/login", JSON.stringify(data));
}

export function useUserLogin(options?: MutationConfig<typeof loginUser>) {
  const queryClient = useQueryClient();

  return useMutation({
    onSuccess: (response) => {
      queryClient.setQueryData([QUERY_KEYS.AUTH_USER], response.data);
    },
    ...options,
    mutationFn: loginUser,
  });
}
