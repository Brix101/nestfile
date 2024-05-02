import { InitialState, Resources } from "@/types/auth";
import { UserResource } from "@/types/user";

export const deriveState = (
  authLoaded: boolean,
  state: Resources,
  initialState: InitialState | undefined,
) => {
  if (!authLoaded && initialState) {
    return deriveFromSsrInitialState(initialState);
  }
  return deriveFromClientSideState(state);
};

const deriveFromSsrInitialState = (initialState: InitialState) => {
  const userId = initialState.userId;
  const user = initialState.user as UserResource;

  return {
    userId,
    user,
  };
};

const deriveFromClientSideState = (state: Resources) => {
  const userId: number | null | undefined = state.user
    ? state.user.id
    : state.user;
  const user = state.user;

  return {
    userId,
    user,
  };
};
