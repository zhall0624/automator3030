import WebhookForm from "./WebhookForm";
import { Stack, Sheet, Typography, Box, Link } from "@mui/joy";
import { useQuery } from "@tanstack/react-query";

type Webhook = {
  name: string;
};
function Webhooks() {
  const fetchWebhooks = (): Promise<Webhook[]> =>
    fetch("localhost:3000/incoming_webhooks").then((response) =>
      response.json(),
    );
  const { data } = useQuery({
    queryKey: ["webhooks"],
    queryFn: fetchWebhooks,
  });
  return (
    <>
      <Typography level="h1">Incoming Webhooks</Typography>
      <Stack>
        {!data ? (
          <></>
        ) : (
          data.map((webhook) => (
            <Box>
              <Sheet>Name: {webhook.name}</Sheet>
              <Sheet>
                <Link href="localhost:8080"> URL </Link>
              </Sheet>
            </Box>
          ))
        )}

        <WebhookForm></WebhookForm>
      </Stack>
    </>
  );
}

export default Webhooks;
