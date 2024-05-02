import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";

import PasswordInput from "@/components/PasswordInput";
import { Icons } from "@/components/icons";
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { loginSchema } from "@/lib/validations/auth";
import { LoginInput } from "@/types/auth";
import { toast } from "sonner";
import { useUserLogin } from "../api/login";
import { LoginDTO } from "../types";

export function UserAuthForm() {
  const form = useForm<LoginDTO>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      username: "",
      password: "",
    },
  });

  const { mutate, isPending } = useUserLogin({
    onError: (error) => {
      const res = error.response;
      if (res?.status === 401) {
        form.setError("username", res.data);
      } else {
        const message = res?.data?.message || error.message;
        toast(res?.statusText, {
          description: message,
        });
      }
    },
  });

  async function onSubmit(data: LoginInput) {
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
                <Input type="text" placeholder="Brix101" {...field} />
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
