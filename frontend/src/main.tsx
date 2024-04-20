import React from "react";
import ReactDOM from "react-dom/client";

import App from "@/App.tsx";
import ReactQueryProvider from "@/providers/ReactQueryProvider";
import AuthProvider from "@/providers/AuthProvider";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <ReactQueryProvider>
      <AuthProvider>
        <App />
      </AuthProvider>
    </ReactQueryProvider>
  </React.StrictMode>,
);
