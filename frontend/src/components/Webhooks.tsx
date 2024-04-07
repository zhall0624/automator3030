import { Box, Link, Typography, Container } from "@mui/material";
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
    return <Typography>Error Loading Webhooks</Typography>;
  }

  if (isLoading) {
    return <Typography>Loading Webhooks</Typography>;
  }
  return (
    <>
      {data?.map((webhook) => (
        <Box>
          <Typography>Name: {webhook.name}</Typography>
          <Container>
            <Link href="localhost:8080"> URL </Link>
          </Container>
        </Box>
      ))}
    </>
  );
}

export default Webhooks;
