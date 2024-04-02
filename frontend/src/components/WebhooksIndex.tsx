import WebhookForm from "./WebhookForm";
import { Stack, Typography } from "@mui/joy";
import Webhooks from "./Webhooks";

function WebhooksIndex() {
  return (
    <>
      <Typography level="h1">Incoming Webhooks</Typography>
      <Stack>
        <Webhooks></Webhooks>
        <WebhookForm></WebhookForm>
      </Stack>
    </>
  );
}

export default WebhooksIndex;
