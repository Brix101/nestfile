import { Serializable } from "@/lib/utils";
import { UserResource } from "@/types/user";

// eslint-disable-next-line @typescript-eslint/ban-types
export type IsSerializable<T> = T extends Function ? false : true;

export type ServerGetTokenOptions = { template?: string };
export type ServerGetToken = (
  options?: ServerGetTokenOptions,
) => Promise<string | null>;

export type InitialState = Serializable<{
  userId: number | undefined;
  user: UserResource | undefined | null;
}>;

export interface Resources {
  user?: UserResource | null;
}
