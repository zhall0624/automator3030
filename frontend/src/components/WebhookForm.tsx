import { SubmitHandler, useForm } from "react-hook-form";
import { useMutation } from "@tanstack/react-query";
import { Input, Button, Typography } from "@mui/joy";
import Webhook from "../types/webhook";

function WebhookForm() {
  const { register, handleSubmit } = useForm<Webhook>();
  const { mutateAsync } = useMutation({
    mutationFn: (data: Webhook) => {
      console.log("hi");
      return fetch("http://localhost:8080/incoming-webhooks", {
        method: "post",
        mode: "no-cors",
        body: JSON.stringify(data),
      });
    },
  });
  const onSubmit: SubmitHandler<Webhook> = (data) => mutateAsync(data);
  return (
    <>
      <Typography level="h2">New Incoming Webhook</Typography>
      <form onSubmit={handleSubmit(onSubmit)}>
        <Input placeholder="Webhook Name" {...register("name")} />
        <Button type="submit">Create Webhook</Button>
      </form>
    </>
  );
}

export default WebhookForm;
