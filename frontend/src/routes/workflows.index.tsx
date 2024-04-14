import { createFileRoute } from '@tanstack/react-router'
import { Workflows } from "../components/workflows"

export const Route = createFileRoute('/workflows/')({
  component: () => <Workflows></Workflows>,
})

