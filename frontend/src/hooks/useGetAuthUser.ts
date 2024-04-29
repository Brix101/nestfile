import { UseQueryOptions, useQuery } from "@tanstack/react-query";
import { AxiosError } from "axios";

import { getUserQuery } from "@/services/user-service";
import { Resources } from "@/types/auth";

function useGetAuthUser(options?: UseQueryOptions<Resources, AxiosError>) {
  return useQuery({
    ...getUserQuery(),
    notifyOnChangeProps: "all",
    ...options,
  });
}
export default useGetAuthUser;
