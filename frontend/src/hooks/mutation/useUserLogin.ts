import {
  UseMutationOptions,
  useMutation,
  useQueryClient,
} from "@tanstack/react-query";
import { AxiosError, AxiosResponse } from "axios";

import { QUERY_KEYS } from "@/constant/query-key";
import { loginUser } from "@/services/user-service";
import { LoginInput } from "@/types/auth";
import { UserResource } from "@/types/user";

function useUserLogin(
  options?: UseMutationOptions<
    AxiosResponse<UserResource, AxiosError>,
    AxiosError<ServerError>,
    LoginInput
  >,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: loginUser,
    onSuccess: (response) => {
      queryClient.setQueryData([QUERY_KEYS.auth_user], response.data);
    },
    ...options,
  });
}

export default useUserLogin;
