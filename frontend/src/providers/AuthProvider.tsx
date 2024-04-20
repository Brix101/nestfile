import { useQuery } from "@tanstack/react-query";
import React from "react";

import { UserContext } from "@/context/user";
import api from "@/lib/api";
import { deriveState } from "@/lib/deriveState";
import { InitialState, Resources } from "@/types/auth";
import { UserResource } from "@/types/user";

export type AuthContextProviderState = Resources;

interface AuthProviderProps extends React.PropsWithChildren {
  initialState?: InitialState;
}

function AuthProvider({ children, initialState }: AuthProviderProps) {
  const { data, isLoading } = useQuery({
    queryKey: ["auth-user"],
    queryFn: async () => {
      const response = await api.get<{ user?: UserResource | null }>(
        "/auth/user",
      );
      return response.data;
    },
  });

  const [state, setState] = React.useState<AuthContextProviderState>({
    user: data?.user,
  });

  React.useEffect(() => {
    if (data) {
      setState({ user: data.user });
    }
    return () => {
      setState({});
    };
  }, [data]);

  const { user } = deriveState(isLoading, state, initialState);
  const userCtx = React.useMemo(() => ({ value: user }), [user]);

  return (
    <UserContext.Provider value={userCtx}>{children}</UserContext.Provider>
  );
}

export function useAssertWrappedByAuthProvider(
  displayNameOrFn: string | (() => void),
): void {
  const ctx = React.useContext(UserContext);

  if (!ctx) {
    if (typeof displayNameOrFn === "function") {
      displayNameOrFn();
      return;
    }

    throw new Error(
      `${displayNameOrFn} can only be used within the <AuthProvider /> component.`,
    );
  }
}

export default AuthProvider;
