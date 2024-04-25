import api from "@/lib/api";
import { Resources } from "@/types/auth";

export async function getUserQuery() {
  const res = await api.get<Resources>("/auth/user");
  return res.data;
}
