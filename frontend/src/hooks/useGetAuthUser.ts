import { UseQueryOptions, useQuery } from "@tanstack/react-query";
import { AxiosError } from "axios";

import { QUERY_KEYS } from "@/constant/query-key";
import { getUserQuery } from "@/services/user-service";
import { Resources } from "@/types/auth";

function useGetAuthUser(options?: UseQueryOptions<Resources, AxiosError>) {
  return useQuery({
    queryKey: [QUERY_KEYS.AUTH_USER],
    queryFn: getUserQuery,
    notifyOnChangeProps: "all",
    ...options,
  });
}
export default useGetAuthUser;
