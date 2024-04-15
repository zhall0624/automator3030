import { createFileRoute } from '@tanstack/react-router'
import { WorkflowManagement } from '../components/workflow'

export const Route = createFileRoute('/workflows/$id')({
  component: () => <WorkflowManagement></WorkflowManagement>
})

