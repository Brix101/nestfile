import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";

import { QUERY_KEYS } from "@/constant/query-key";
import api from "@/lib/api";
import { authSchema } from "@/lib/validations/auth";
import { UserResource } from "@/types/user";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { Icons } from "../icons";
import { PasswordInput } from "../password-input";
import { Button } from "../ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "../ui/form";
import { Input } from "../ui/input";

type Inputs = z.infer<typeof authSchema>;

export function UserAuthForm() {
  const queryClient = useQueryClient();

  const { mutate, isPending } = useMutation({
    mutationFn: (data: Inputs) => {
      return api.post<UserResource>("/auth/login", JSON.stringify(data));
    },
    onSuccess: (response) => {
      queryClient.setQueryData([QUERY_KEYS.auth_user], response.data);
    },
    onError: (error) => {
      console.log(error);
    },
  });

  const form = useForm<Inputs>({
    resolver: zodResolver(authSchema),
    defaultValues: {
      username: "",
      password: "",
    },
  });

  async function onSubmit(data: Inputs) {
    mutate(data);
  }

  return (
    <Form {...form}>
      <form className="grid gap-4" onSubmit={form.handleSubmit(onSubmit)}>
        <FormField
          control={form.control}
          name="username"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Username</FormLabel>
              <FormControl>
                <Input type="text" placeholder="username" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="password"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Password</FormLabel>
              <FormControl>
                <PasswordInput placeholder="**********" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button type="submit" className="mt-2" disabled={isPending}>
          {isPending && (
            <Icons.spinner
              className="mr-2 size-4 animate-spin"
              aria-hidden="true"
            />
          )}
          Login
          <span className="sr-only">Login</span>
        </Button>
      </form>
    </Form>
  );
}
