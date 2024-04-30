import { z } from "zod";

import { loginSchema } from "@/lib/validations/auth";
import { UserResource } from "@/types/user";

/**
 * Excludes any non-serializable prop from an object
 */
export type Serializable<T> = {
  [K in keyof T as IsSerializable<T[K]> extends true ? K : never]: T[K];
};

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

export type LoginInput = z.infer<typeof loginSchema>;
