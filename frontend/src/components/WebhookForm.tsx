import { useForm, SubmitHandler } from "react-hook-form"

type FormValues = {
  name: string
}

function WebhookForm() {
  const { register, handleSubmit } = useForm<FormValues>()
  const onSubmit: SubmitHandler<FormValues> = (data) =>
    fetch('localhost:3000/incoming_webhooks', { method: 'post', body: JSON.stringify(data) })
  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <input {...register("name")} />
      <button type="submit">Create Webhook</button>
    </form>
  )
}

export default WebhookForm
