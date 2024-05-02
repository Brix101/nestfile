import { queryClient } from "@/lib/react-query";
import { QueryClientProvider } from "@tanstack/react-query";
import React from "react";
import AuthProvider from "./AuthProvider";
import { ThemeProvider } from "./ThemeProvider";
import { TailwindIndicator } from "@/components/TailwindIndicator";
import { Toaster } from "@/components/ui/sonner";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";

interface AppProviderProps extends React.PropsWithChildren {}

export function AppProvider({ children }: AppProviderProps) {
  return (
    <QueryClientProvider client={queryClient}>
      <AuthProvider>
        <ThemeProvider defaultTheme="light">{children}</ThemeProvider>
      </AuthProvider>
      <ReactQueryDevtools initialIsOpen={false} />
      <TailwindIndicator />
      <Toaster />
    </QueryClientProvider>
  );
}
