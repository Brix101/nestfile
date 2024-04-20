import { createContextAndHook } from "@/lib/createContextAndHook";
import { UserResource } from "@/types/user";

const [UserContext, useUserContext] = createContextAndHook<
  UserResource | null | undefined
>("UserContext");

export { UserContext, useUserContext };
