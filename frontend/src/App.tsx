import "@/assets/css/index.css";
import "@/assets/css/custom.css";

import { TailwindIndicator } from "@/components/TailwindIndicator";
import { Toaster } from "@/components/ui/sonner";
import PageProvider from "@/providers/PageProvider";

function App() {
  return (
    <>
      <PageProvider />
      <TailwindIndicator />
      <Toaster />
    </>
  );
}

export default App;
