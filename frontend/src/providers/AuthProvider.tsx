import { useQueryClient } from "@tanstack/react-query";

import { InitialState, getUserQuery } from "@/features/auth";
import AuthContextProvider from "./AuthContextProvider";

function AuthProvider({ children }: React.PropsWithChildren) {
  const queryClient = useQueryClient();
  const query = getUserQuery();

  type QueryType = Awaited<ReturnType<typeof query.queryFn>>;

  const data = queryClient.getQueryData<QueryType>(query.queryKey);

  const initialState: InitialState = {
    userId: data?.user?.id,
    user: data?.user,
  };

  return (
    <AuthContextProvider initialState={initialState}>
      {children}
    </AuthContextProvider>
  );
}

export default AuthProvider;
