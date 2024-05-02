import { z } from "zod";
import { loginSchema } from "../schema";

export interface UserResource {
  id: number;
  username: string;
  createdAt: string;
  updatedAt: string;
}

export type LoginDTO = z.infer<typeof loginSchema>;
