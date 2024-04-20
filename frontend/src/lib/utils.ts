import { IsSerializable } from "@/types/auth";
import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

/**
 * Excludes any non-serializable prop from an object
 */
export type Serializable<T> = {
  [K in keyof T as IsSerializable<T[K]> extends true ? K : never]: T[K];
};
