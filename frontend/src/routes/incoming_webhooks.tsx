import { WebhooksIndex } from "../components/webhook";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/incoming_webhooks")({
  component: () => <WebhooksIndex></WebhooksIndex>,
});
