import React from "react";
import { QueryClientProvider } from "@tanstack/react-query";

import { queryClient } from "@/lib/react-query";
import AuthProvider from "./AuthProvider";
import { ThemeProvider } from "./ThemeProvider";
import { TailwindIndicator } from "@/components/TailwindIndicator";
import { Toaster } from "@/components/ui/sonner";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { Loader } from "@/components/Loader";
import { BrowserRouter as Router } from "react-router-dom";

interface AppProviderProps extends React.PropsWithChildren {}

export function AppProvider({ children }: AppProviderProps) {
  return (
    <React.Suspense
      fallback={
        <div className="fixed top-0 left-0 z-50 h-screen w-full flex justify-center items-center">
          <Loader />
        </div>
      }
    >
      <QueryClientProvider client={queryClient}>
        <AuthProvider>
          <ThemeProvider defaultTheme="light">
            <Router>{children}</Router>
          </ThemeProvider>
        </AuthProvider>
        <ReactQueryDevtools initialIsOpen={false} />
        <TailwindIndicator />
        <Toaster />
      </QueryClientProvider>
    </React.Suspense>
  );
}
