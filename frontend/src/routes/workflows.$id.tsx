import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/workflows/$id')({
  component: () => <div>Hello /workflows/$id!</div>
})