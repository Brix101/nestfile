import React from "react";

import { UserContext } from "@/context/user";
import useGetAuthUser from "@/hooks/useGetAuthUser";
import { deriveState } from "@/lib/deriveState";
import { InitialState, Resources } from "@/types/auth";

export type AuthContextProviderState = Resources;

interface AuthProviderProps extends React.PropsWithChildren {
  initialState?: InitialState;
}

function AuthProvider({ children, initialState }: AuthProviderProps) {
  const { data, isLoading } = useGetAuthUser();

  const [state, setState] = React.useState<AuthContextProviderState>({
    ...data,
  });

  React.useEffect(() => {
    if (data) {
      setState(data);
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
