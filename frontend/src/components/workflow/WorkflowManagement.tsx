import { Typography } from '@mui/material'
import { useQuery } from '@tanstack/react-query'
import { Workflow } from '../../types'
import { Route } from '../../routes/workflows.$id'

export default function WorkflowManagement() {
  const { id } = Route.useParams()

  const fetchWorkflow = (): Promise<Workflow> =>
    fetch(`http://localhost:8080/workflows/${id}`).then((response) =>
      response.json(),
    );

  const { data, isError, isLoading } = useQuery<Workflow>({
    queryKey: ["workflows", id],
    queryFn: fetchWorkflow,
  });

  if (isError) {
    return (<Typography>There was an error</Typography>)
  }

  if (isLoading) {
    return (<Typography>Loading Workflow</Typography>)
  }

  if (!data) {
    return (<Typography>Why is data undefined?</Typography>)
  }

  return (
    <>
      <Typography>Workflow: {data.name}</Typography>
      <Typography></Typography>
    </>
  )
}
