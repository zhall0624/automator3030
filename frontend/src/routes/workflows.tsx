import { createFileRoute, Link as RouterLink } from '@tanstack/react-router'
import { Container, Table, TableHead, TableCell, Typography, TableRow, Link, Input, Button, TableBody } from "@mui/material"
import { useQuery, useQueryClient, useMutation } from '@tanstack/react-query'
import { useForm } from 'react-hook-form'
import { Workflow } from "../types"

export const Route = createFileRoute('/workflows')({
  component: () => <Workflows></Workflows>,
})

function Workflows() {
  const { register, handleSubmit } = useForm<Workflow>();

  const queryClient = useQueryClient();

  const { mutateAsync } = useMutation({
    mutationFn: (data: Workflow) => {
      return fetch("http://localhost:8080/workflows", {
        method: "post",
        body: JSON.stringify(data),
        headers: {
          "Content-Type": "application/json",
        },
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["workflows"] });
    },
  });

  const onSubmit: SubmitHandler<Workflow> = (data) => mutateAsync(data);

  const fetchWorkflows = (): Promise<Workflow[]> =>
    fetch("http://localhost:8080/workflows").then((response) =>
      response.json(),
    );

  const { data, isError, isLoading } = useQuery<Workflow[]>({
    queryKey: ["workflows"],
    queryFn: fetchWorkflows,
  });

  if (isError) {
    return <Typography>Error Loading Workflows</Typography>;
  }

  if (isLoading) {
    return <Typography>Workflows Loading</Typography>;
  }

  return (
    < Container >
      <Table>
        <TableHead>
          <TableRow>
            <TableCell>Workflow Name</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {data?.map((row) => (
            <TableRow key={row.id} >
              <TableCell><Link component={RouterLink} to="/blog/$id" params={{ id: row.id }}>{row.name}</Link></TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
      <form onSubmit={handleSubmit(onSubmit)}>
        <Input {...register("name")}></Input>
        <Button type="submit">Create Workflow</Button>
      </form>
    </Container >
  )
}
