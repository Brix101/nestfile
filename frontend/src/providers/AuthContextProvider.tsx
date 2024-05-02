import React from "react";

import { UserContext } from "@/context/user";
import { InitialState, Resources } from "@/types/auth";
import { useQueryAuthUser } from "@/hooks/query";
import { deriveState } from "@/utils/deriveState";

export type AuthContextProviderState = Resources;

interface AuthProviderProps extends React.PropsWithChildren {
  initialState?: InitialState;
}

function AuthContextProvider({ children, initialState }: AuthProviderProps) {
  const { data, isLoading } = useQueryAuthUser();

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

  const derivedState = deriveState(!isLoading, state, initialState);
  const { user, userId } = derivedState;

  const userCtx = React.useMemo(() => ({ value: user }), [userId, user]);

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

export default AuthContextProvider;
