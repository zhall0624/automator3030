import { QueryClientProvider, QueryClient } from "@tanstack/react-query";
import WebhooksIndex from "./components/WebhooksIndex";
import "@fontsource/inter";

const queryClient = new QueryClient();
function App() {
  return (
    <>
      <QueryClientProvider client={queryClient}>
        <WebhooksIndex></WebhooksIndex>
      </QueryClientProvider>
    </>
  );
}

export default App;
