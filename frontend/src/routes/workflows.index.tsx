import { createFileRoute } from '@tanstack/react-router'
import { Workflows } from "../components/workflow"

export const Route = createFileRoute('/workflows/')({
  component: () => <Workflows></Workflows>,
})

