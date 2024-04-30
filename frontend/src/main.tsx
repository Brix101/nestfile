import React from "react";
import ReactDOM from "react-dom/client";

import App from "@/App.tsx";
import AuthProvider from "@/providers/AuthProvider";
import ReactQueryProvider from "@/providers/ReactQueryProvider";
import { ThemeProvider } from "@/providers/ThemeProvider";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <ReactQueryProvider>
      <AuthProvider>
        <ThemeProvider defaultTheme="light">
          <App />
        </ThemeProvider>
      </AuthProvider>
    </ReactQueryProvider>
  </React.StrictMode>,
);
