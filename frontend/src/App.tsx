import { QueryClientProvider, QueryClient } from "@tanstack/react-query";
import Webhooks from "./components/Webhooks";
import "@fontsource/inter";

const queryClient = new QueryClient();
function App() {
  return (
    <>
      <QueryClientProvider client={queryClient}>
        <Webhooks></Webhooks>
      </QueryClientProvider>
    </>
  );
}

export default App;
