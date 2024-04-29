import { useQueryClient } from "@tanstack/react-query";

import { getUserQuery } from "@/services/user-service";
import { InitialState, Resources } from "@/types/auth";
import AuthContextProvider from "./AuthContextProvider";

function AuthProvider({ children }: React.PropsWithChildren) {
  const queryClient = useQueryClient();

  const query = getUserQuery();
  const data = queryClient.getQueryData<Resources>(query.queryKey);
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
