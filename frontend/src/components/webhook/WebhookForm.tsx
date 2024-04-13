import { SubmitHandler, useForm } from "react-hook-form";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { Input, Button, Typography } from "@mui/material";
import { Webhook } from "../../types";

function WebhookForm() {
  const { register, handleSubmit } = useForm<Webhook>();
  const queryClient = useQueryClient();
  const { mutateAsync } = useMutation({
    mutationFn: (data: Webhook) => {
      return fetch("http://localhost:8080/incoming-webhooks", {
        method: "post",
        body: JSON.stringify(data),
        headers: {
          "Content-Type": "application/json",
        },
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["webhooks"] });
    },
  });
  const onSubmit: SubmitHandler<Webhook> = (data) => mutateAsync(data);
  return (
    <>
      <Typography variant="h2">New Incoming Webhook</Typography>
      <form onSubmit={handleSubmit(onSubmit)}>
        <Input placeholder="Webhook Name" {...register("name")} />
        <Button type="submit">Create Webhook</Button>
      </form>
    </>
  );
}

export default WebhookForm;
