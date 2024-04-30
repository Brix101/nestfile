import {
  UseMutationOptions,
  useMutation,
  useQueryClient,
} from "@tanstack/react-query";
import { AxiosError, AxiosResponse } from "axios";

import { QUERY_KEYS } from "@/constant/query-key";
import { logoutUser } from "@/services/user-service";
import { UserResource } from "@/types/user";

function useUserLogout(
  options?: UseMutationOptions<AxiosResponse<UserResource, AxiosError>>,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: logoutUser,
    onSuccess: (response) => {
      queryClient.setQueryData([QUERY_KEYS.auth_user], response.data);
    },
    ...options,
  });
}

export default useUserLogout;
