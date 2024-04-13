import { WebhookForm } from "../components/webhook";
import { Stack, Typography } from "@mui/material";
import { Webhooks } from "../components/webhook";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/incoming_webhooks")({
  component: WebhooksIndex,
});

function WebhooksIndex() {
  return (
    <>
      <Typography variant="h1">Incoming Webhooks</Typography>
      <Stack>
        <Webhooks></Webhooks>
        <WebhookForm></WebhookForm>
      </Stack>
    </>
  );
}
