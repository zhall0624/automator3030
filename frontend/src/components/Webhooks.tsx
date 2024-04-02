import { Sheet, Box, Link } from "@mui/joy";
import { useQuery } from "@tanstack/react-query";
import Webhook from "../types/webhook";

function Webhooks() {
  const fetchWebhooks = (): Promise<Webhook[]> =>
    fetch("http://localhost:8080/incoming-webhooks").then((response) =>
      response.json(),
    );
  const { data, isError, isLoading } = useQuery<Webhook[]>({
    queryKey: ["webhooks"],
    queryFn: fetchWebhooks,
  });

  if (isError) {
    return <Sheet>Error Loading Webhooks</Sheet>;
  }

  if (isLoading) {
    return <Sheet>Loading Webhooks</Sheet>;
  }
  return (
    <>
      {data?.map((webhook) => (
        <Box>
          <Sheet>Name: {webhook.name}</Sheet>
          <Sheet>
            <Link href="localhost:8080"> URL </Link>
          </Sheet>
        </Box>
      ))}
    </>
  );
}

export default Webhooks;
