import { SubmitHandler, useForm } from "react-hook-form";
import { useMutation } from "@tanstack/react-query";
import { Input, Button, Typography } from "@mui/joy";

type FormValues = {
  name: string;
};

function WebhookForm() {
  const { register, handleSubmit } = useForm<FormValues>();
  const { mutate } = useMutation({
    mutationFn: (data: FormValues) => {
      return fetch("localhost:3000/incoming_webhooks", {
        method: "post",
        body: JSON.stringify(data),
      });
    },
  });
  const onSubmit: SubmitHandler<FormValues> = (data) => mutate(data);
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
