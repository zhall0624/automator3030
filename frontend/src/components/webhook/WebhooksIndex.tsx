import WebhookForm from "./WebhookForm"
import { Stack, Typography } from "@mui/material";
import Webhooks from "./Webhooks"

export default function WebhooksIndex() {
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
