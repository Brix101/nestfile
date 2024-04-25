import api from "@/lib/api";
import { Resources } from "@/types/auth";

export async function getUserQuery() {
  const response = await api.get<Resources>("/auth/user");
  return response.data;
}
