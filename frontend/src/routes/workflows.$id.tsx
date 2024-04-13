import { createFileRoute } from '@tanstack/react-router'
import { WorkflowManagement } from '../components/workflows'

export const Route = createFileRoute('/workflows/$id')({
  component: () => <WorkflowManagement></WorkflowManagement>
})

